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

func PrintWinRate(n int, a, b Interface) EloDiff {
	score := make(chan int)

	rounds := make(chan int, n)
	go func() {
		for i := 0; i < n; i++ {
			rounds <- i
		}
		close(rounds)
	}()

	planets := 60

	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			moves := make([]strategy.Move, 1000)
			for round := range rounds {
				err := doGame(planets, round, a, b, moves, score)
				if err != nil {
					panic(err)
				}
			}
			wg.Done()
		}()
	}

	sum := 0
	scoreWg := sync.WaitGroup{}
	scoreWg.Add(1)
	go func() {
		for s := range score {
			sum += s
		}
		scoreWg.Done()
	}()

	wg.Wait()
	close(score)
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

func doGame(planets, round int, a, b Interface, moves []strategy.Move, score chan int) error {
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
		return err
	}

	winner := 0
	err = RunGame(g, func(g *game.Game) (bool, error) {
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
		scores := Scores(2, g.Planets)
		win := MaxScore(scores) >= planets*2/3
		if win {
			for p, score := range scores {
				if score >= planets*2/3 {
					winner = p
				}
			}
		}
		return win, nil
	})
	if err != nil {
		return err
	}

	if (winner+round)%2 == 1 {
		score <- 1
	}
	return nil
}
