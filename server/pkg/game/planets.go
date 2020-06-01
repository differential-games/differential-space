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

func NewPlanet(x, y float64, owner int, ready bool, strength int) Planet {
	radius := math.Sqrt(x*x + y*y)
	invRadius := 1.0 / radius

	angle := math.Atan2(y, x)
	return Planet{
		X:         x,
		Y:         y,
		Radius:    radius,
		InvRadius: invRadius,
		Angle:     angle,
		Owner:     owner,
		Ready:     ready,
		Strength:  strength,
	}
}

func DistSq(p1x, p1y, p2x, p2y float64) float64 {
	return (p1x-p2x)*(p1x-p2x) + (p1y-p2y)*(p1y-p2y)
}

func closestTo(x, y float64, planets []Planet, ignore map[int]bool) int {
	closest := 0
	minDistanceSq := math.MaxFloat64
	for i, p := range planets {
		if ignore[i] {
			continue
		}
		distanceSq := DistSq(x, y, p.X, p.Y)
		if distanceSq < minDistanceSq {
			closest = i
			minDistanceSq = distanceSq
		}
	}
	return closest
}
