package competition

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/differential-games/differential-space/pkg/ai"
	"github.com/differential-games/differential-space/pkg/ai/strategy"
)

// League represents a pool of players in a league and their rankings.
type League struct {
	Players []func() ai.Interface
	Elos    []float64

	Moves []strategy.Move
	K     float64
}

func NewLeague(players []func() ai.Interface) League {
	return League{
		Players: players,
		Elos:    make([]float64, len(players)),
		Moves:   make([]strategy.Move, strategy.NPlanets*15),
		K:       1,
	}
}

func (l *League) Tournament() float64 {
	startElos := make([]float64, len(l.Elos))
	copy(startElos, l.Elos)

	for i := 0; i < 10; i++ {
		l.Round()
	}

	diff := 0.0
	for i, n := range l.Elos {
		newDiff := math.Abs(n - startElos[i])
		diff += newDiff * newDiff
	}
	fmt.Println(math.Sqrt(diff))
	fmt.Printf("%+v\n", l.Elos)
	return math.Sqrt(diff)
}

type AB struct {
	A, B int
}

func (l *League) Round() {
	var rounds []AB
	for i := range l.Players {
		for j := range l.Players {
			if i != j {
				rounds = append(rounds, AB{A: i, B: j})
			}
		}
	}
	rand.Shuffle(len(rounds), func(i, j int) {
		rounds[i], rounds[j] = rounds[j], rounds[i]
	})
	for _, r := range rounds {
		l.Fight(r.A, r.B)
	}
}

func (l *League) Fight(a, b int) {
	//fmt.Printf("%d vs %d\n", a, b)
	//_, err := ai.RunAIGame(strategy.NPlanets,
	//	0, l.Players[a](), l.Players[b](), l.Moves)
	//if err != nil {
	//	panic(err)
	//}

	ea := 1.0 / (1.0 + math.Pow(10.0, (l.Elos[b]-l.Elos[a])/400.0))

	aWins := false
	switch {
	case a-b == 2:
		aWins = rand.Float64() < 0.96
	case a-b == 1:
		aWins = rand.Float64() < 0.8
	case a-b == -1:
		aWins = rand.Float64() < 0.2
	case a-b == -2:
		aWins = rand.Float64() < 0.04
	}
	//if a > b && (rand.Float64() < 0.8) {
	//	aWins = true
	//} else if a < b && (rand.Float64() < 0.2) {
	//	aWins = true
	//}

	if aWins {
		//fmt.Println(a, "won")
		// a won
		l.Elos[a] += l.K * (1.0 - ea)
		l.Elos[b] -= l.K * (1.0 - ea)
	} else {
		//fmt.Println(b, "won")
		// b won
		l.Elos[a] -= l.K * ea
		l.Elos[b] += l.K * ea
	}
	//fmt.Printf("%+v\n", l.Elos)
}
