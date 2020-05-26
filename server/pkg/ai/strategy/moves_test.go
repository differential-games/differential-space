package strategy

import (
	"testing"

	"github.com/differential-games/differential-space/pkg/game"
)

func planet(x, y float64) game.Planet {
	return game.Planet{
		X:        x,
		Y:        y,
		Owner:    1,
		Strength: 1,
		Ready:    true,
	}
}

func TestGenerateMoves(t *testing.T) {
	tcs := []struct {
		name    string
		planets []game.Planet
		wantLen int
	}{
		{
			name: "1 planet",
			planets: []game.Planet{
				planet(0, 1),
			},
			wantLen: 0,
		},
		{
			name: "2 planets",
			planets: []game.Planet{
				planet(0, 1),
				planet(1, 0),
			},
			wantLen: 2,
		},
		{
			name: "3 planets",
			planets: []game.Planet{
				planet(0, 1),
				planet(1, 0),
				planet(0, -1),
			},
			wantLen: 6,
		},
		{
			name: "4 planets",
			planets: []game.Planet{
				planet(0, 1),
				planet(1, 0),
				planet(0, -1),
				planet(-1, 0),
			},
			wantLen: 12,
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			moves := make([]Move, 20)
			got := GenerateMoves(1, tc.planets, moves)

			if len(got) != tc.wantLen {
				t.Errorf("got %d moves, want %d", len(got), tc.wantLen)
			}
		})
	}
}
