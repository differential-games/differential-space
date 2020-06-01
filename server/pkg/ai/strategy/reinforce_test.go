package strategy_test

import (
	"github.com/differential-games/differential-space/pkg/ai"
	"github.com/differential-games/differential-space/pkg/ai/strategy"
	"github.com/differential-games/differential-space/pkg/game"
	"testing"
)

func TestReinforceFront(t *testing.T) {
	tcs := []TestCase{
		{
			name: "sequence of max",
			planets: []game.Planet{
				game.NewPlanet(0, 0, 2, false, 0),
				game.NewPlanet(0, 4, 1, true, 1),
				game.NewPlanet(0, 8, 1, true, 8),
				game.NewPlanet(0, 12, 1, true, 8),
				game.NewPlanet(0, 16, 1, true, 8),
				game.NewPlanet(0, 20, 1, true, 8),
			},
			want: []FromTo{
				{From: 2, To: 1},
				{From: 3, To: 2},
				{From: 4, To: 3},
				{From: 5, To: 4},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, tc.Run(func() ai.Interface {
			return &ai.AI{
				Difficulty: 1.0,
				Strategies: []strategy.Vector{
					ai.NewVector(1.0, strategy.NewMovePriority(-1, -1, -1)),
					ai.NewVector(1.0, strategy.NewReinforceFront(100)),
				},
			}
		}))
	}
}
