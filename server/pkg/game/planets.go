package game

import "math"

type Planet struct {
	X float64
	Y float64

	Radius    float64
	InvRadius float64

	Angle float64

	// Owner is the id of the Player who owns this Planet.
	// 0 indicates no owner.
	Owner int

	// Ready indicates if the Planet has taken an action this turn.
	Ready bool

	// Strength is the number of ships at the Planet.
	Strength int
}

func distToSq(x, y float64, p *Planet) float64 {
	return (x-p.X)*(x-p.X) + (y-p.Y)*(y-p.Y)
}

func DistSq(p1, p2 *Planet) float64 {
	return (p1.X-p2.X)*(p1.X-p2.X) + (p1.Y-p2.Y)*(p1.Y-p2.Y)
}

func closestTo(x, y float64, planets []Planet, ignore map[int]bool) int {
	closest := 0
	minDistanceSq := math.MaxFloat64
	for i, p := range planets {
		if ignore[i] {
			continue
		}
		distanceSq := distToSq(x, y, &p)
		if distanceSq < minDistanceSq {
			closest = i
			minDistanceSq = distanceSq
		}
	}
	return closest
}
