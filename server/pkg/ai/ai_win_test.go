package ai_test

import (
	"github.com/differential-games/differential-space/pkg/ai"
	"github.com/differential-games/differential-space/pkg/ai/aitest"
	"github.com/differential-games/differential-space/pkg/game"
	"math/rand"
	"sort"
	"sync"
	"testing"
	"time"
)

func MaxSecondPlaceGame(t *testing.T) int {
	g, err := game.New(game.DefaultOptions)
	if err != nil {
		t.Fatal(err)
	}

	maxScores := make([]int, 17)
	player := &ai.AI{
		Difficulty: 1.0,
		Colonize: []ai.MoveAnalyzer{
			ai.PreferFewerNeighbors,
			ai.PreferFurther,
		},
		Attack: []ai.MoveAnalyzer{
			ai.PreferCloser,
		},
	}
	aitest.RunGame(g, func(g *game.Game) bool {
		aitest.Turn(t, g, player)
		maxScores = aitest.Maxes(maxScores, aitest.Scores(g.Planets))
		return aitest.MaxScore(maxScores) >= game.DefaultOptions.NumPlanets
	})

	sort.Ints(maxScores)
	return maxScores[3]
}

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
		got := MaxSecondPlaceGame(t)
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

type Counter struct {
	mtx sync.Mutex
	count int
}

func (c *Counter) Inc() {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.count++
}

func (c *Counter) Get() int {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	return c.count
}

// 0.25 = +861
// 0.5 = +1553
// 0.75 = +1891
// 1.0 = +2172

func Test_WinRate(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	rand.Seed(time.Now().UTC().UnixNano())

	// Odd-numbered players
	a := &ai.AI{
		Difficulty:     1.0,
		Colonize: []ai.MoveAnalyzer{
			ai.PreferFewerNeighbors,
			ai.PreferFurther,
		},
		Attack: []ai.MoveAnalyzer{
			ai.PreferCloser,
		},
		AttackFilter: []ai.MoveFilterBuilder{
			ai.CoordinatedAttackAtDiff(0.25, 0.25),
			ai.AttackAtMaxStrength,
			ai.AttackAtProbability(0.5),
		},
	}

	// Even-numbered players
	b := &ai.AI{
		Difficulty:     0.5,
		Colonize: []ai.MoveAnalyzer{
			//ai.PreferFewerNeighbors,
			ai.PreferFurther,
		},
		Attack: []ai.MoveAnalyzer{
			ai.PreferCloser,
		},
		AttackFilter: []ai.MoveFilterBuilder{
			//ai.CoordinatedAttackAtDiff(0.25, 0.25),
			ai.AttackAtMaxStrength,
			ai.AttackAtProbability(0.5),
		},
	}

	winA := Counter{}
	planets := 60

	var wg sync.WaitGroup
	n := 10000
	for i := 0; i < n; i++ {
		wg.Add(1)

		go func(round int) {
			defer wg.Done()

			g, err := game.New(game.Options{
				PlanetOptions: game.PlanetOptions{
					NumPlanets: planets,
					Radius:     12,
					MinRadius:  4,
				},
				PlayerOptions: game.PlayerOptions{
					NumPlayers: 2,
				},
				RotationSpeed: 2.0,
			})
			if err != nil {
				t.Fatal(err)
			}

			winner := 0
			aitest.RunGame(g, func(g *game.Game) bool {
				if (g.PlayerTurn+round+1)%2 == 1 {
					aitest.Turn(t, g, a)
				} else {
					aitest.Turn(t, g, b)
				}
				scores := aitest.Scores(g.Planets)
				win := aitest.MaxScore(scores) >= planets*2/3
				if win {
					for p, score := range scores {
						if score >= planets*2/3 {
							winner = p
						}
					}
				}
				return win
			})

			if (winner+round)%2 == 1 {
				winA.Inc()
			}
		}(i)
	}

	wg.Wait()
	t.Fatalf("%.2f%%", float64(winA.count)/(float64(n) / 100.0))
}
