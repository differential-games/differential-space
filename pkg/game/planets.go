package game

import "math"

type Planet struct {
	X float64
	Y float64

	Radius float64
	Angle float64

	// Owner is the id of the Player who owns this Planet.
	// 0 indicates no owner.
	Owner int

	// Ready indicates if the Planet has taken an action this turn.
	Ready bool

	// Strength is the number of ships at the Planet.
	Strength int
}

func distTo(x, y float64, p Planet) float64 {
	return math.Sqrt(math.Pow(x - p.X, 2.0) + math.Pow(y - p.Y, 2.0))
}

func Dist(p1, p2 Planet) float64 {
	return distTo(p1.X, p1.Y, p2)
}

func closestTo(x, y float64, planets []Planet, ignore map[int]bool) int {
	closest := 0
	minDistance := math.MaxFloat64
	for i, p := range planets {
		if ignore[i] {
			continue
		}
		distance := distTo(x, y, p)
		if distance < minDistance {
			closest = i
			minDistance = distance
		}
	}
	return closest
}
