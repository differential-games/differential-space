package stats

import (
	"fmt"
	"math"
	"strings"
)

type Counter struct {
	// Input accepts new elements.
	// Elements must be from 0 to 99.
	Input chan int

	Data [100]int
}

func (c *Counter) Process() {
	for i := range c.Input {
		c.Data[i]++
	}
}

func (c *Counter) Total() int {
	total := 0
	for _, ct := range c.Data {
		total += ct
	}
	return total
}

func (c *Counter) String() string {
	total := float64(c.Total())

	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("Count: %d\n", c.Total()))
	sb.WriteString(fmt.Sprintf("Entropy: %.2f\n", c.Entropy()))
	for i, ct := range c.Data {
		if ct > 0 {
			sb.WriteString(fmt.Sprintf("%d: %.2f%%\n", i, (100*float64(ct))/total))
		}
	}
	return sb.String()
}

func (c *Counter) Bins() int {
	bins := 0
	for _, ct := range c.Data {
		if ct > 0 {
			bins++
		}
	}
	return bins
}

func (c *Counter) Entropy() float64 {
	total := float64(c.Total())
	bins := float64(c.Bins())

	entropy := 0.0
	for _, ct := range c.Data {
		if ct > 0 {
			p := float64(ct) * bins / total
			entropy += -p * math.Log2(p)
		}
	}
	return entropy / bins
}
