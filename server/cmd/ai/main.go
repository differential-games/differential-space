package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"runtime/pprof"
	"time"

	"github.com/differential-games/differential-space/pkg/game"

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

	ai.PrintWinRate(10000, a, b)
}

func vb(weight float64, builder strategy.StrategyBuilder) strategy.VectorBuilder {
	return strategy.VectorBuilder{
		Builder: builder,
		Weight:  weight,
	}
}

var a = &ai.AI{
	Difficulty: 1.0,
	Strategies: []strategy.VectorBuilder{
		vb(1.0, strategy.Builder(strategy.MovePriority(0.0, -1.0, -1.0))),
		vb(1.0, strategy.PreferFewerNeighbors),
		vb(1.0, strategy.Builder(strategy.PreferFurther(strategy.Colonize))),
		vb(1.0, strategy.Builder(strategy.PreferCloser(strategy.Attack))),
		vb(1.0, strategy.Builder(strategy.AttackAtStrength(game.MaxFleetSize))),
		vb(1.0, strategy.CoordinatedAttack),
		vb(1.0, strategy.ReinforceFront),
	},
}

var base = &ai.AI{
	Difficulty: 1.0,
	Strategies: nil,
}

var b = &ai.AI{
	Difficulty: 1.0,
	Strategies: []strategy.VectorBuilder{
		vb(1.0, strategy.Builder(strategy.MovePriority(0.0, -1.0, -1.0))),
		vb(1.0, strategy.PreferFewerNeighbors),
		vb(1.0, strategy.Builder(strategy.PreferFurther(strategy.Colonize))),
		vb(1.0, strategy.Builder(strategy.PreferCloser(strategy.Attack))),
		vb(1.0, strategy.Builder(strategy.AttackAtStrength(game.MaxFleetSize))),
		vb(1.0, strategy.CoordinatedAttack),
		vb(1.0, strategy.ReinforceFront),
	},
}
