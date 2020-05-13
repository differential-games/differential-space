package game

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestFight_CertainOutcome(t *testing.T) {
	// For all of these tests, the distance is zero so the outcome is certain.
	if p := WinProbability(0.0); p != 1.0 {
		// This test will fail or be flaky otherwise.
		t.Fatalf("want WinProbability(0.0) = 1.0, got %f", p)
	}

	tcs := []struct {
		attacker     int
		defender     int
		wantWin      bool
		wantAttacker int
		wantDefender int
	}{
		{
			attacker:     1,
			defender:     0,
			wantWin:      true,
			wantAttacker: 1,
			wantDefender: 0,
		},
		{
			attacker:     1,
			defender:     1,
			wantWin:      false,
			wantAttacker: 0,
			wantDefender: 0,
		},
		{
			attacker:     8,
			defender:     0,
			wantWin:      true,
			wantAttacker: 8,
			wantDefender: 0,
		},
		{
			attacker:     8,
			defender:     4,
			wantWin:      true,
			wantAttacker: 4,
			wantDefender: 0,
		},
		{
			attacker:     4,
			defender:     8,
			wantWin:      false,
			wantAttacker: 0,
			wantDefender: 4,
		},
		{
			attacker:     8,
			defender:     8,
			wantWin:      false,
			wantAttacker: 0,
			wantDefender: 0,
		},
	}

	for _, tc := range tcs {
		name := fmt.Sprintf("%d vs %d", tc.attacker, tc.defender)
		t.Run(name, func(t *testing.T) {
			attacker := Planet{
				Strength: tc.attacker,
			}
			defender := Planet{
				Strength: tc.defender,
			}
			gotWin := Fight(&attacker, &defender)

			if tc.wantWin != gotWin {
				t.Errorf("got fight() = %t, want %t", gotWin, tc.wantWin)
			}
			if attacker.Strength != tc.wantAttacker {
				t.Errorf("got attacker.Strength = %d, want %d", attacker.Strength, tc.wantAttacker)
			}
			if defender.Strength != tc.wantDefender {
				t.Errorf("got defender.Strength = %d, want %d", defender.Strength, tc.wantDefender)
			}
		})
	}
}

func TestFight_UncertainOutcome(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	rand.Seed(time.Now().UTC().UnixNano())

	tcs := []struct {
		attacker          int
		defender          int
		dist              float64
		wantWinPercentage float64
	}{
		{
			attacker:          1,
			defender:          0,
			dist:              1.0,
			wantWinPercentage: 0.9,
		},
		{
			attacker:          2,
			defender:          0,
			dist:              1.0,
			wantWinPercentage: 0.99,
		},
		{
			attacker:          1,
			defender:          0,
			dist:              5.0,
			wantWinPercentage: 0.5,
		},
		{
			attacker:          2,
			defender:          0,
			dist:              5.0,
			wantWinPercentage: 0.75,
		},
		{
			attacker:          8,
			defender:          4,
			dist:              5.0,
			wantWinPercentage: 0.36328125,
		},
		{
			attacker:          8,
			defender:          4,
			dist:              2.5,
			wantWinPercentage: 0.8861846923828125,
		},
	}

	// At n=100,000, the tests are sensitive to differences of less than 1%.
	n := 100000

	for _, tc := range tcs {
		name := fmt.Sprintf("%d vs %d @ %.2f units", tc.attacker, tc.defender, tc.dist)
		t.Run(name, func(t *testing.T) {
			gotWins := 0
			gotLosses := 0
			// Simulate n battles and tabulate wins and losses.
			for i := 0; i < n; i++ {
				attacker := Planet{
					Strength: tc.attacker,
				}
				defender := Planet{
					Strength: tc.defender,
					X:        tc.dist,
				}

				win := Fight(&attacker, &defender)
				if win {
					gotWins++
				} else {
					gotLosses++
				}
			}

			// Calculate the difference between observed and expected win percentages.
			wantWins := float64(n) * tc.wantWinPercentage
			diff := math.Abs(float64(gotWins)-wantWins) / float64(n)

			// This error tolerance is sensitive to deviances greater than 1% with a very low flake rate.
			if diff > 0.01 {
				t.Errorf("got %d wins, want %.2f, p=%.2f", gotWins, wantWins, tc.wantWinPercentage*100)
			}
		})
	}
}
