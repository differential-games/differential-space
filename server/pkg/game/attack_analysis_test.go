package game

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestBattleWinChance(t *testing.T) {
	tcs := []struct {
		name           string
		attack, defend int
		p              float64
		want           float64
	}{
		{
			name:   "1v0 @ 90% = 90%",
			attack: 1,
			defend: 0,
			p:      0.9,
			want:   0.9,
		},
		{
			name:   "2v0 @ 90% = 99%",
			attack: 2,
			defend: 0,
			p:      0.9,
			want:   0.99,
		},
		{
			name:   "3v0 @ 90% = 99.9%",
			attack: 3,
			defend: 0,
			p:      0.9,
			want:   0.999,
		},
		{
			name:   "3v2 @ 50% = 12.5%",
			attack: 3,
			defend: 2,
			p:      0.5,
			want:   0.125,
		},
		{
			name:   "3v1 @ 50% = 50%",
			attack: 3,
			defend: 1,
			p:      0.5,
			want:   0.5,
		},
		{
			name:   "6v3 @ 80% = 90.1%",
			attack: 6,
			defend: 3,
			p:      0.8,
			want:   0.901,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got := BattleWinChance(tc.attack, tc.defend, tc.p)

			if diff := cmp.Diff(tc.want, got, cmpopts.EquateApprox(0.0, 0.001)); diff != "" {
				t.Error(diff)
			}
		})
	}
}
