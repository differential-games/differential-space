package main

import (
	"flag"
	"fmt"
	"github.com/differential-games/differential-space/pkg/stats"
	"log"
	"math/rand"
	"os"
	"runtime/pprof"
	"time"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
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

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	digest := stats.NewTDigest(500)

	for i := 0; i < 1e8; i++ {
		val := r.Float64()
		digest.Add(val)
	}

	fmt.Println(digest.String())

	err := 0.0
	for i := 0; i <= 10; i++ {
		p := float64(i) / 10
		q := digest.Quantile(p)
		//fmt.Println(i, q)
		err += (p - q) * (p - q)
	}
	fmt.Printf("SSE: %f\n", err)
}
