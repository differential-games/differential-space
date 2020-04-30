package ai

import (
	"github.com/differential-games/differential-space/pkg/game"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func Scores(planets []game.Planet) []int {
	result := make([]int, 5)
	for _, p := range planets {
		if p.Owner != 0 {
			result[p.Owner]++
		}
	}
	return result
}

func Maxes(left, right []int) []int {
	result := make([]int, 5)
	for i := range left {
		if left[i] > right[i] {
			result[i] = left[i]
		} else {
			result[i] = right[i]
		}
	}
	return result
}

func MaxScore(scores []int) int {
	result := scores[0]
	for _, s := range scores {
		if s > result {
			result = s
		}
	}
	return result
}

func MaxSecondPlaceGame(t *testing.T) int {
	g, err := game.New(game.DefaultOptions)
	if err != nil {
		t.Fatal(err)
	}

	maxScores := make([]int, 5)
	for ; MaxScore(maxScores) < game.DefaultOptions.NumPlanets; {
		g.NextTurn()
		player := g.PlayerTurn
		// There's a lot more back-and-forth between sub-optimal players.
		// We're optimizing the win condition for reasonably good play.
		moves := PickMoves(player, g.Planets, 1.0)

		for _, m := range moves {
			err = g.Move(m)
			if err != nil {
				t.Fatal(err)
			}
		}
		maxScores = Maxes(maxScores, Scores(g.Planets))
	}
	sort.Ints(maxScores)
	return maxScores[3]
}

func Test_MaxScore(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	rand.Seed(time.Now().UTC().UnixNano())
	// Ensures that with default options, the first player to get over 2/3rds
	// of all planets is the winner at least 99% of the time.

	exceptions := 0
	// Simulate 1000 AI-controlled games and track the max score of the second-place player.
	for i := 0; i < 1000; i++ {
		got := MaxSecondPlaceGame(t)
		if got > game.DefaultOptions.NumPlanets * 2 / 3 {
			t.Log(got)
			exceptions++
		}
	}
	if exceptions >= 5 {
		// We do actually get cases where players lose after passing the
		// 2/3rds threshold, but they're less than 1% of games played at a high level.
		t.Errorf("got exceptions = %d, want < 5", exceptions)
	}
}
