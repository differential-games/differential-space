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
		Strategies: []strategy.Vector{
			ai.NewVector(1.0, strategy.NewMovePriority(0.0, -1.0, -1.0)),
			ai.NewVector(1.0, &strategy.PreferFewerNeighbors{}),
			ai.NewVector(1.0, strategy.PreferFurther{}),
			ai.NewVector(1.0, strategy.PreferCloser{}),
			ai.NewVector(1.0, &strategy.CoordinatedAttack{}),
			ai.NewVector(1.0, strategy.AttackIfWinLikely{}),
			ai.NewVector(1.0, strategy.NewReinforceFront(500)),
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
		Strategies: []strategy.Vector{
			ai.NewVector(1.0, strategy.NewMovePriority(0.0, -1.0, -1.0)),
			ai.NewVector(1.0, &strategy.PreferFewerNeighbors{}),
			ai.NewVector(1.0, strategy.PreferFurther{}),
			ai.NewVector(1.0, strategy.PreferCloser{}),
			ai.NewVector(1.0, &strategy.CoordinatedAttack{}),
			ai.NewVector(1.0, strategy.AttackIfWinLikely{}),
			ai.NewVector(1.0, strategy.NewReinforceFront(500)),
		},
	}
}
