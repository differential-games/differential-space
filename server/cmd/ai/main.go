package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"runtime/pprof"
	"time"

	"github.com/differential-games/differential-space/pkg/ai"
	"github.com/differential-games/differential-space/pkg/ai/strategy"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}

		err = pprof.StartCPUProfile(f)
		if err != nil {
			log.Fatal(err)
		}
		defer pprof.StopCPUProfile()
	}

	ai.PrintWinRate(100000, a, b)
}

func a() ai.Interface {
	return &ai.AI{
		Difficulty: 1.0,
		Strategies: []strategy.Strategy{
			strategy.NewMovePriority(0.0, -1.0, -1.0),
			&strategy.PreferFewerNeighbors{},
			strategy.PreferFurther{},
			strategy.PreferCloser{},
			&strategy.CoordinatedAttack{},
			strategy.AttackIfWinLikely{},
			strategy.NewReinforceFront(strategy.NPlanets * 10),
		},
	}
}

func base() ai.Interface {
	return &ai.AI{
		Difficulty: 1.0,
		Strategies: nil,
	}
}

func b() ai.Interface {
	return &ai.AI{
		Difficulty: 1.0,
		Strategies: []strategy.Strategy{
			strategy.NewMovePriority(0.0, -1.0, -1.0),
			&strategy.PreferFewerNeighbors{},
			strategy.PreferFurther{},
			strategy.PreferCloser{},
			&strategy.CoordinatedAttack{},
			strategy.AttackIfWinLikely{},
			strategy.NewReinforceFront(strategy.NPlanets * 10),
		},
	}
}
