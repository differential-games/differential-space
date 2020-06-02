package ai_test

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/differential-games/differential-space/pkg/ai/strategy"

	"github.com/differential-games/differential-space/pkg/ai"
	"github.com/differential-games/differential-space/pkg/game"
)

func Test_MaxScore(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	t.Skip()

	rand.Seed(time.Now().UTC().UnixNano())
	// Ensures that with default options, the first player to get over 2/3rds
	// of all planets is the winner at least 99% of the time.

	exceptions := 0
	// Simulate 1000 AI-controlled games and track the max score of the second-place player.
	for i := 0; i < 1000; i++ {
		got := maxSecondPlaceGame(t)
		if got > game.DefaultOptions.NumPlanets*2/3 {
			t.Log(got)
			exceptions++
		}
	}
	if exceptions > 10 {
		// We do actually get cases where players lose after passing the
		// 2/3rds threshold, but they're less than 1% of games played at a high level.
		t.Errorf("got exceptions = %d, want <= 5", exceptions)
	}
}

// maxSecondPlaceGame returns the score of the second-place player in a game.
func maxSecondPlaceGame(t *testing.T) int {
	g, err := game.New(game.DefaultOptions)
	if err != nil {
		t.Fatal(err)
	}

	maxScores := make([]int, game.DefaultOptions.NumPlayers+1)

	err = ai.RunGame(g, func(g *game.Game) (bool, error) {
		moves := make([]strategy.Move, 1000)
		err := ai.Turn(g, ai.Hard(1.0), moves)
		if err != nil {
			return false, err
		}
		maxScores = ai.Maxes(maxScores, ai.Scores(game.DefaultOptions.NumPlayers, g.Planets))
		return ai.MaxScore(maxScores) >= game.DefaultOptions.NumPlanets, nil
	})

	if err != nil {
		t.Fatal(err)
	}

	sort.Ints(maxScores)

	return maxScores[3]
}
