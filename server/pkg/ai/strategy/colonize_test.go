package strategy_test

import (
	"testing"

	"github.com/differential-games/differential-space/pkg/ai"
	"github.com/differential-games/differential-space/pkg/ai/strategy"
	"github.com/differential-games/differential-space/pkg/game"
)

type FromTo struct {
	From, To int
}

type TestCase struct {
	name    string
	planets []game.Planet
	want    []FromTo
}

func (tc TestCase) Run(playerFn func() ai.Interface) func(t *testing.T) {
	return func(t *testing.T) {
		player := playerFn()
		player.Initialize(game.Game{Planets: tc.planets})

		got := player.PickMoves(1, tc.planets, make([]strategy.Move, 100))

		if len(got) != len(tc.want) {
			t.Fatalf("got %d moves, want %d", len(got), len(tc.want))
		}
		for i := range got {
			if (got[i].From != tc.want[i].From) || (got[i].To != tc.want[i].To) {
				t.Errorf("%d: got %+v, want %+v", i, got[i], tc.want[i])
			}
		}
	}
}

func TestPreferFewerNeighbors(t *testing.T) {
	tcs := []TestCase{
		{
			name: "one before two neighbors",
			planets: []game.Planet{
				game.NewPlanet(-4, 0, 0, false, 0),
				game.NewPlanet(0, 0, 1, true, 1),
				game.NewPlanet(4, 0, 0, false, 0),
				game.NewPlanet(8, 0, 1, true, 1),
			},
			want: []FromTo{
				{From: 1, To: 0},
				{From: 1, To: 2},
				{From: 3, To: 2},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, tc.Run(func() ai.Interface {
			return &ai.AI{
				Difficulty: 1.0,
				Strategies: []strategy.Strategy{
					strategy.NewMovePriority(0, -1, -1),
					&strategy.PreferFewerNeighbors{},
				},
			}
		}))
	}
}
