package strategy_test

import (
	"testing"

	"github.com/differential-games/differential-space/pkg/ai"
	"github.com/differential-games/differential-space/pkg/ai/strategy"
	"github.com/differential-games/differential-space/pkg/game"
)

func TestAttackAtStrength(t *testing.T) {
	tcs := []TestCase{
		{
			name: "ignore below threshold",
			planets: []game.Planet{
				game.NewPlanet(0, 0, 1, true, 3),
				game.NewPlanet(0, 1, 2, false, 0),
			},
			want: []FromTo{},
		},
		{
			name: "use at threshold",
			planets: []game.Planet{
				game.NewPlanet(0, 0, 1, true, 4),
				game.NewPlanet(0, 1, 2, false, 0),
			},
			want: []FromTo{
				{From: 0, To: 1},
			},
		},
		{
			name: "ignore colonize",
			planets: []game.Planet{
				game.NewPlanet(0, 0, 1, true, 4),
				game.NewPlanet(0, 1, 0, false, 0),
			},
		},
		{
			name: "use above threshold",
			planets: []game.Planet{
				game.NewPlanet(0, 0, 1, true, 5),
				game.NewPlanet(0, 1, 2, false, 0),
			},
			want: []FromTo{
				{From: 0, To: 1},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, tc.Run(func() ai.Interface {
			return &ai.AI{
				Difficulty: 1.0,
				Strategies: []strategy.Strategy{
					strategy.NewMovePriority(-1, -1, -1),
					strategy.NewAttackAtStrength(4),
				},
			}
		}))
	}
}

func TestAttackIfWinLikely(t *testing.T) {
	tcs := []TestCase{
		{
			name: "ignore unlikely",
			planets: []game.Planet{
				game.NewPlanet(0, 0, 1, true, 3),
				game.NewPlanet(0, 1, 2, false, 3),
			},
			want: []FromTo{},
		},
		{
			name: "attack likely",
			planets: []game.Planet{
				game.NewPlanet(0, 0, 1, true, 3),
				game.NewPlanet(0, 1, 2, false, 2),
			},
			want: []FromTo{
				{From: 0, To: 1},
			},
		},
		{
			name: "ignore colonize",
			planets: []game.Planet{
				game.NewPlanet(0, 0, 1, true, 4),
				game.NewPlanet(0, 1, 0, false, 0),
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, tc.Run(func() ai.Interface {
			return &ai.AI{
				Difficulty: 1.0,
				Strategies: []strategy.Strategy{
					strategy.NewMovePriority(-1, -1, -1),
					strategy.AttackIfWinLikely{},
				},
			}
		}))
	}
}
