package game

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"testing"
)

func TestNewDistribution(t *testing.T) {
	tcs := []struct{
		name string
		min, max int
		want Distribution
	}{
		{
			name: "0, 1",
			min: 0,
			max: 1,
			want: Distribution{
				Starts: []float64{0.0},
			},
		},
		{
			name: "0, 2",
			min: 0,
			max: 2,
			want: Distribution{
				Starts: []float64{0.0, 0.333},
			},
		},
		{
			name: "0, 3",
			min: 0,
			max: 3,
			want: Distribution{
				Starts: []float64{0.0, 0.167, 0.5},
			},
		},
		{
			name: "1, 2",
			min: 1,
			max: 2,
			want: Distribution{
				Starts: []float64{0.333},
			},
		},
		{
			name: "1, 4",
			min: 1,
			max: 4,
			want: Distribution{
				Starts: []float64{0.1, 0.3, 0.6},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			d, err := NewDistribution(tc.min, tc.max)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(&tc.want, d, cmpopts.EquateApprox(0.0, 0.001)); diff != "" {
				t.Error(diff)
			}
		})
	}

}

func TestNewDistribution_Error(t *testing.T) {
	_, err := NewDistribution(2, 1)
	if err == nil {
		t.Error("got NewDistribution(2, 1) = nil, want err")
	}
}

func TestGeneratePlanets(t *testing.T) {
	planets, err := GeneratePlanets(DefaultOptions.Planets)

	if err != nil {
		t.Fatalf("got GeneratePlanets() = %s, want nil", err)
	}

	if len(planets) != DefaultOptions.Planets.NumPlanets {
		t.Errorf("got len(GeneratePlanets()) = %d, want %d", len(planets), DefaultOptions.Planets.NumPlanets)
	}
}

func TestGeneratePlanets_Failures(t *testing.T) {
	tcs := []struct{
		name string
		options PlanetOptions
	}{
		{
			name: "radius too small",
			options: PlanetOptions{
				NumPlanets: 100,
				Radius:     2,
				MinRadius:  1,
			},
		},
		{
			name: "min greater than max",
			options: PlanetOptions{
				NumPlanets: 100,
				Radius:     1,
				MinRadius:  2,
			},
		},
		{
			name: "min greater than max",
			options: PlanetOptions{
				NumPlanets: maxPlanets + 1,
				Radius:     100,
				MinRadius:  1,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			_, err := GeneratePlanets(tc.options)

			if err == nil {
				t.Fatal("got GeneratePlanets() = nil, want err")
			}
		})
	}
}
