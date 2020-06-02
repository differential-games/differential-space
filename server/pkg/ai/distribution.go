package ai

import (
	"fmt"
	"math"
)

type Distribution struct {
	logBase float64
	counts  []int
}

func NewDistribution(logBase float64) Distribution {
	return Distribution{
		logBase: logBase,
		counts:  make([]int, 40),
	}
}

func (d *Distribution) Add(v int) {
	power := int(math.Log10(float64(v)) / d.logBase)
	d.counts[power]++
}

func (d *Distribution) Print() {
	fmt.Println("Distribution")
	total := 0
	for _, count := range d.counts {
		total += count
	}
	fmt.Println(total)

	for i, count := range d.counts {
		if count == 0 {
			continue
		}
		base := math.Pow(10, d.logBase)
		fmt.Printf("%.0f: %.3f%%\n", math.Pow(base, float64(i)), float64(100*count)/float64(total))
	}
}
