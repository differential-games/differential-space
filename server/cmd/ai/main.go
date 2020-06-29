package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime/pprof"
	"sync"
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

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		strategy.PreferCloserCounter.Process()
		wg.Done()
	}()

	ai.PrintWinRate(10000, a, b)
	close(strategy.PreferCloserCounter.Input)

	wg.Wait()
	fmt.Println(strategy.PreferCloserCounter.String())
}

func colonize() ai.Interface {
	return &ai.AI{
		Difficulty: 1.0,
		Strategies: []strategy.Strategy{
			strategy.MovePriority{Attack: -1.0, Reinforce: -1.0},
			&strategy.PreferFewerNeighbors{},
			strategy.PreferFurther{},
		},
	}
}

func a() ai.Interface {
	return &ai.AI{
		Difficulty: 1.0,
		Strategies: []strategy.Strategy{
			strategy.MovePriority{Attack: -1.0, Reinforce: -1.0},
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
		Difficulty: 0.2,
		Strategies: []strategy.Strategy{
			strategy.MovePriority{Attack: -1.0, Reinforce: -1.0},
			&strategy.PreferFewerNeighbors{},
			strategy.PreferFurther{},
			strategy.PreferCloser{},
			&strategy.CoordinatedAttack{},
			strategy.AttackIfWinLikely{},
			strategy.NewReinforceFront(strategy.NPlanets * 10),
		},
	}
}
