package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/differential-games/differential-space/pkg/ai"
	"github.com/differential-games/differential-space/pkg/ai/competition"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	players := make([]func() ai.Interface, 3)
	for i := 0; i < 3; i++ {
		players[i] = func() ai.Interface {
			return ai.Hard(float64(i+1) / 5.0)
		}
	}

	league := competition.NewLeague(players)

	//league.Elos = []float64{-300, -100, 0, 100, 300}
	for i := 0; i < 100; i++ {
		league.Tournament()
	}

	fmt.Printf("%+v\n", league.Elos)
}
