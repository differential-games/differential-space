package ai

import (
	"fmt"
	"math"
	"sync"

	"github.com/differential-games/differential-space/pkg/ai/strategy"

	"github.com/differential-games/differential-space/pkg/game"
)

type EloDiff struct {
	// Mean is the measured Elo difference between two AIs.
	Mean float64

	// D5 is the (negative) p=5% difference between Mean and the minimum of a 95% confidence threshold.
	D5 float64

	// D95 is the (positive) p=95% difference between Mean and maximum of a 95% confidence threshold.
	D95 float64
}

func (e EloDiff) String() string {
	return fmt.Sprintf("Elo Diff: (%.2f, %.2f, %.2f)\n", e.Mean+e.D5, e.Mean, e.Mean+e.D95)
}

func PrintWinRate(n int, a, b func() Interface) EloDiff {
	outcomes := make(chan Outcome)

	rounds := make(chan int, n)
	go func() {
		for i := 0; i < n; i++ {
			rounds <- i
		}
		close(rounds)
	}()

	planets := strategy.NPlanets

	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			moves := make([]strategy.Move, 1000)
			for round := range rounds {
				outcome, err := RunAIGame(planets, round, a(), b(), moves)
				if err != nil {
					panic(err)
				}
				outcomes <- outcome
			}
			wg.Done()
		}()
	}

	sum := 0
	scoreWg := sync.WaitGroup{}
	scoreWg.Add(1)
	go func() {
		for o := range outcomes {
			if (o.Round+o.Winner)%2 == 1 {
				// Player 1 won
				sum++
			}
		}
		scoreWg.Done()
	}()

	wg.Wait()
	close(outcomes)
	scoreWg.Wait()

	p := float64(sum) / (float64(n))
	sd := math.Sqrt(p * (1.0 - p) / float64(n))
	diff := -400 * math.Log10(1/p-1.0)
	diffMin := -400 * math.Log10(1/(p-2*sd)-1.0)
	diffMax := -400 * math.Log10(1/(p+2*sd)-1.0)

	elo := EloDiff{
		Mean: diff,
		D5:   diffMin - diff,
		D95:  diffMax - diff,
	}
	fmt.Printf("%.2f%%\n", p*100.0)
	fmt.Print(elo)
	return elo
}

type Outcome struct {
	Round  int
	Winner int
	Turns  int
}

func RunAIGame(planets, round int, a, b Interface, moves []strategy.Move) (Outcome, error) {
	g, err := game.New(game.Options{
		PlanetOptions: game.PlanetOptions{
			NumPlanets: planets,
			Radius:     10,
			MinRadius:  4,
		},
		PlayerOptions: game.PlayerOptions{
			NumPlayers: 2,
		},
		RotationSpeed: 2.0,
	})
	if err != nil {
		return Outcome{}, err
	}
	a.Initialize(*g)
	b.Initialize(*g)

	winner := 0
	nTurns := 0
	err = RunGame(g, func(g *game.Game) (bool, error) {
		nTurns++
		if (g.PlayerTurn+round+1)%2 == 1 {
			err := Turn(g, a, moves)
			if err != nil {
				return false, err
			}
		} else {
			err := Turn(g, b, moves)
			if err != nil {
				return false, err
			}
		}
		if g.Winner != 0 {
			winner = g.Winner
			return true, nil
		}
		return false, nil
	})
	if err != nil {
		return Outcome{}, err
	}

	return Outcome{
		Round:  round,
		Winner: winner,
		Turns:  nTurns,
	}, nil
}
