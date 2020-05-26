package game

import (
	"math"
	"math/rand"

	"github.com/pkg/errors"
)

const maxPlanets = 1000

func GeneratePlanets(options PlanetOptions) ([]Planet, error) {
	if options.NumPlanets > maxPlanets {
		return nil, errors.Errorf("numPlanets '%d' higher than maximum '%d'", options.NumPlanets, maxPlanets)
	}

	planets := make([]Planet, options.NumPlanets)

	d, err := NewDistribution(options.MinRadius, options.Radius)
	if err != nil {
		return nil, err
	}

	for i := 0; i < options.NumPlanets; i++ {
		tries := 0
		var newPlanet *Planet
		for ; tries < 100; tries++ {
			newPlanet = generatePlanet(options.MinRadius, *d)
			if !hasCollision(newPlanet, planets[:i]) {
				break
			}
		}
		if tries == 100 {
			return nil, errors.Errorf(
				"unable to find space for planet %d; increase galaxy radius", i)
		}
		planets[i] = *newPlanet
	}

	return planets, nil
}

type Distribution struct {
	// Starts is a monotonically-increasing list of probabilities in the range (0, 1).
	Starts []float64
}

func (d Distribution) RandInt() int {
	r := rand.Float64()
	for i, p := range d.Starts {
		if r < p {
			return i
		}
	}
	return len(d.Starts)
}

func NewDistribution(min, max int) (*Distribution, error) {
	if min > max {
		return nil, errors.Errorf("min radius %d greater than max radius %d", min, max)
	}

	var sum int
	for i := min; i <= max; i++ {
		sum += i
	}

	d := &Distribution{
		Starts: make([]float64, max-min),
	}

	csum := 0.0
	for i := min; i < max; i++ {
		csum += float64(i) / float64(sum)
		d.Starts[i-min] = csum
	}
	return d, nil
}

func hasCollision(newPlanet *Planet, planets []Planet) bool {
	for _, p := range planets {
		if DistSq(newPlanet, &p) < 1 {
			return true
		}
	}
	return false
}

func generatePlanet(minRadius int, d Distribution) *Planet {
	radius := float64(minRadius) + float64(d.RandInt())

	angle := rand.Float64() * 2 * math.Pi

	return &Planet{
		X: radius * math.Sin(angle),
		Y: radius * math.Cos(angle),

		Radius:    radius,
		InvRadius: 1.0 / radius,
		Angle:     angle,
	}
}
