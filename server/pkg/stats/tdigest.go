package stats

import (
	"fmt"
	"math"
	"strings"
)

type centroid struct {
	mean  float64
	count float64

	// maxCount is the cached maximum count.
	maxCount float64
	// nCentroids is the cached number of centroids the last time we calculated
	// maxCount.
	nCentroids int
}

func (c *centroid) String() string {
	return fmt.Sprintf("mean: %.4f, count: %d", c.mean, int(c.count))
}

// inc increments the centroid with val.
func (c *centroid) inc(val float64) {
	c.count++
	c.mean += (val - c.mean) / c.count
}

type TDigest struct {
	centroids   []*centroid
	compression float64
	count       float64

	nCentroids int

	// The cached estimates of the centroids containing the 5% and 95%
	// percentiles. Updated when new centroids are added.
	p5Centroid  int
	p95Centroid int

	// appendLower is whether to append to the lower of the two closest
	// centroids.
	appendLower bool
}

func (d *TDigest) String() string {
	sb := strings.Builder{}
	for _, c := range d.centroids {
		sb.WriteString(fmt.Sprintln(c.String()))
	}
	return sb.String()
}

func NewTDigest(compression float64) *TDigest {
	return &TDigest{
		compression: compression,
	}
}

// nearest returns the index such that the returned index and its immediate
// successor are indices of the two closest centroids.
//
// centroids is a list of sorted centroids of increasing mean.
// Must contain at least 2 elements.
func (d *TDigest) nearest(val float64) int {
	left, right := 0, d.nCentroids

	// While this makes the case of small nCentroids more inefficient, it
	// speeds up the case where nCentroids >= 10, which is the far more common
	// case as over 90% of elements are inserted in the middle 30% of centroids.
	if d.centroids[d.p5Centroid].mean < val {
		left = d.p5Centroid
		if val < d.centroids[d.p95Centroid].mean {
			right = d.p95Centroid + 1
		}
	} else {
		right = d.p5Centroid
	}

	diff := right - left
	// While the difference between left and right is greater than 32, use a
	// binary search.
	for ; diff > 32; diff = right - left {
		// Remember that middle is rounded down.
		// Middle for each iteration is guaranteed to be unique.
		middle := left + diff/2
		if d.centroids[middle].mean < val {
			// val is to the right of the middle considered centroid.
			if val < d.centroids[middle+1].mean {
				// middle is what we're looking for, so exit early.
				return middle
			}
			left = middle + 1
		} else {
			// val is to the left of the middle considered centroid.
			if d.centroids[middle-1].mean < val {
				// middle is to the right of what we're looking for, so exit early.
				return middle - 1
			}
			right = middle
		}
	}

	// Search upwards.
	for i, c := range d.centroids[left+1:] {
		if val < c.mean {
			return left + i
		}
	}
	return right - 1
}

// hasRoom returns true if the centroid at idx has room for more elements.
func (d *TDigest) hasRoom(idx int) bool {
	// With the naive implementation where we recalculate the limit every time,
	// this function is a huge bottleneck in the program, takes over 90% of the
	// runtime of TDigest.Add().
	//
	// Thus, we cache the value and only recalculate when we could possibly be
	// at the limit.
	centroid := d.centroids[idx]
	if centroid.count < centroid.maxCount {
		// We exit here instead of continuing 87% of the time.
		// The cached value says its okay, so we assume it hasn't decreased.
		return true
	}

	// Practically, the percentile of a given centroid doesn't change much. The
	// real variable that can increase capacity is the number of centroids. If
	// it hasn't increased, the weight limit is highly unlikely to have
	// increased.
	if centroid.nCentroids == d.nCentroids {
		// We exit here instead of continuing 96% of the time.
		return false
	}

	// We're at the cached value and the number of centroids has increased,
	// so actually check if the new weight limit has increased.
	centroid.maxCount = d.weightLimit(idx)
	centroid.nCentroids = d.nCentroids
	return centroid.count < centroid.maxCount
}

// weightLimit is the maximum acceptable count for the centroid at index idx.
func (d *TDigest) weightLimit(idx int) float64 {
	ptile := d.quantileOf(idx)
	return 4 * d.compression * ptile * (1 - ptile) * float64(d.nCentroids)
}

// quantileOf returns the approximate quantile of centroid idx.
func (d *TDigest) quantileOf(idx int) float64 {
	if idx > (d.nCentroids / 2) {
		// Since we're near the top, compute the quantile by beginning at the
		// top of the distribution, instead of the bottom. This keeps us from
		// having to iterate unnecessarily for large percentiles.
		var total float64
		for _, c := range d.centroids[idx+1:] {
			total += c.count
		}
		return 1.0 - (d.centroids[idx].count/2+total)/d.count
	}

	var total float64
	for _, c := range d.centroids[:idx] {
		total += c.count
	}
	return (d.centroids[idx].count/2 + total) / d.count
}

// addCentroid adds a new centroid at index idx with mean mean.
func (d *TDigest) addCentroid(idx int, mean float64) {
	d.nCentroids++
	d.centroids = append(d.centroids, nil)
	copy(d.centroids[idx+1:], d.centroids[idx:])
	d.centroids[idx] = &centroid{mean: mean, count: 1}

	if d.nCentroids >= 3 {
		// Cache the centroids that cover approximately the 5% to 95% case,
		// since most centroids are small edge cases near the boundary. This way
		// we can optimize for the 90% case, and cut down on iterations inside
		// the d.nearest() loop.
		//
		// We can peg this to specific index without computation as the
		// quantile index of the pth percentile converges to a constant fraction
		// of the total number of centroids as centroids increases. Here,
		// guessing is more performant than getting the exact answer.
		//
		// The improvement from this is marginal, but measurable. (~4ns)
		d.p5Centroid = d.nCentroids * 3 / 8
		d.p95Centroid = (d.nCentroids * 5 / 8) + 1
	}
}

// Add adds val to the TDigest.
func (d *TDigest) Add(val float64) {
	d.add(val)
	d.count++
}

// add adds a new value, val to the TDigest but does not increment the total
// count.
func (d *TDigest) add(val float64) {
	// Cover the trivial cases.
	switch d.nCentroids {
	case 0:
		// We haven't added any centroids.
		d.addCentroid(0, val)
		return
	case 1:
		// There is exactly one centroid.
		centroid := d.centroids[0]
		if centroid.count < d.compression {
			// It isn't full yet.
			centroid.inc(val)
			return
		}
		// We've got to add the second centroid.
		if val < centroid.mean {
			// val is less than the centroid, so it is now the lowest.
			d.addCentroid(0, val)
		} else {
			// val is greater than the centroid, so it is now the highest.
			d.addCentroid(1, val)
		}
		return
	}

	leftIdx := d.nearest(val)
	left := d.centroids[leftIdx]
	leftHasRoom := d.hasRoom(leftIdx)
	switch {
	case val < left.mean:
		if leftHasRoom {
			left.inc(val)
			return
		}
		// left has no room, so add a new centroid at index 0.
		d.addCentroid(0, val)
		return
	case leftIdx == len(d.centroids)-1:
		// val is a new maximum.
		if leftHasRoom {
			// Add val to the leftmost centroid.
			left := d.centroids[leftIdx]
			left.inc(val)
		} else {
			// Create a new centroid for the new maximum.
			d.addCentroid(len(d.centroids), val)
		}
		return
	}

	// val is between left and right.
	// This is the most common case.
	// Whichever centroid we add val to, it is guaranteed to not change the
	// ordering of left and right.
	rightHasRoom := d.hasRoom(leftIdx + 1)
	right := d.centroids[leftIdx+1]
	switch {
	case leftHasRoom && rightHasRoom:
		// It's most common for both to have room, so check this first.
		// Flip between the two.
		if d.appendLower {
			left.inc(val)
		} else {
			right.inc(val)
		}
		d.appendLower = !d.appendLower
	case leftHasRoom && !rightHasRoom:
		left.inc(val)
	case !leftHasRoom && rightHasRoom:
		right.inc(val)
	default:
		// Neither centroid has room, so create a new one between the two.
		d.addCentroid(leftIdx+1, val)
	}
}

func (d *TDigest) Quantile(q float64) float64 {
	n := len(d.centroids)
	switch n {
	case 0:
		return math.NaN()
	case 1:
		return d.centroids[0].mean
	}

	if q < 0 {
		q = 0
	} else if q > 1 {
		q = 1
	}

	// rescale into count units.
	q = d.count * q

	var qTotal float64
	var idx int
	for i, c := range d.centroids {
		if qTotal+c.count/2 >= q {
			idx = i
			break
		}
		qTotal += c.count
		idx = i
	}

	switch idx {
	case 0:
		c0 := d.centroids[0]
		c1 := d.centroids[1]
		slope := 2 * (c1.mean - c0.mean) / (c1.count + c0.count)
		deltaQ := q - c0.count/2
		return c0.mean + slope*deltaQ
	case n - 1:
		c0 := d.centroids[n-2]
		c1 := d.centroids[n-1]
		slope := 2 * (c1.mean - c0.mean) / (c1.count + c0.count)
		deltaQ := q - (qTotal - c1.count/2)
		return c1.mean + slope*deltaQ
	}

	c0 := d.centroids[idx-1]
	c1 := d.centroids[idx]
	slope := 2 * (c1.mean - c0.mean) / (c1.count + c0.count)
	deltaQ := q - (c1.count/2 + qTotal)
	return c1.mean + slope*deltaQ
}
