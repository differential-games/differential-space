package strategy_test

import (
	"testing"

	"github.com/differential-games/differential-space/pkg/ai"
	"github.com/differential-games/differential-space/pkg/ai/strategy"
	"github.com/differential-games/differential-space/pkg/game"
)

func TestPreferCloser(t *testing.T) {
	tcs := []TestCase{
		{
			name: "order by distance",
			planets: []game.Planet{
				game.NewPlanet(0, 0, 1, true, 1),
				game.NewPlanet(4, 0, 2, false, 0),
				game.NewPlanet(1, 0, 2, false, 0),
				game.NewPlanet(3, 0, 2, false, 0),
				game.NewPlanet(2, 0, 2, false, 0),
			},
			want: []FromTo{
				{From: 0, To: 2},
				{From: 0, To: 4},
				{From: 0, To: 3},
				{From: 0, To: 1},
			},
		},
		{
			name: "ignore colonize",
			planets: []game.Planet{
				game.NewPlanet(0, 0, 1, true, 1),
				game.NewPlanet(4, 0, 0, false, 0),
				game.NewPlanet(1, 0, 0, false, 0),
				game.NewPlanet(3, 0, 0, false, 0),
				game.NewPlanet(2, 0, 0, false, 0),
			},
			want: []FromTo{},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, tc.Run(func() ai.Interface {
			return &ai.AI{
				Difficulty: 1.0,
				Strategies: []strategy.Vector{
					ai.NewVector(1.0, strategy.NewMovePriority(-0.01, 0, -1)),
					ai.NewVector(1.0, &strategy.PreferCloser{}),
				},
			}
		}))
	}
}

func TestPreferFurther(t *testing.T) {
	tcs := []TestCase{
		{
			name: "order by distance",
			planets: []game.Planet{
				game.NewPlanet(0, 0, 1, true, 1),
				game.NewPlanet(4, 0, 0, false, 0),
				game.NewPlanet(1, 0, 0, false, 0),
				game.NewPlanet(3, 0, 0, false, 0),
				game.NewPlanet(2, 0, 0, false, 0),
			},
			want: []FromTo{
				{From: 0, To: 1},
				{From: 0, To: 3},
				{From: 0, To: 4},
				{From: 0, To: 2},
			},
		},
		{
			name: "ignore attack",
			planets: []game.Planet{
				game.NewPlanet(0, 0, 1, true, 1),
				game.NewPlanet(4, 0, 2, false, 0),
				game.NewPlanet(1, 0, 2, false, 0),
				game.NewPlanet(3, 0, 2, false, 0),
				game.NewPlanet(2, 0, 2, false, 0),
			},
			want: []FromTo{},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, tc.Run(func() ai.Interface {
			return &ai.AI{
				Difficulty: 1.0,
				Strategies: []strategy.Vector{
					ai.NewVector(1.0, strategy.NewMovePriority(0, -0.01, -1)),
					ai.NewVector(1.0, &strategy.PreferFurther{}),
				},
			}
		}))
	}
}
