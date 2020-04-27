package main

import (
	"fmt"
	"github.com/differential-games/differential-space/pkg/game"
	"github.com/differential-games/differential-space/pkg/serve"
	"math/rand"
	"net/http"
	"os"
	"time"
)

// Get a planet's position.
// curl -s localhost:8080/planets/1

// Set a planet's position.
// curl -X PUT -d '{"x": 1, "y": 2}' localhost:8080/planets/1

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	g, err := game.New(game.DefaultOptions)
	if err != nil {
		panic(err)
	}

	s := serve.Server{
		Game: g,
	}
	http.HandleFunc("/attack", s.HandleAttack)
	http.HandleFunc("/attack/predict", s.HandleAttackPredict)
	http.HandleFunc("/game", s.HandleGame)
	http.HandleFunc("/planets", s.HandlePlanets)
	http.HandleFunc("/planets/", s.HandlePlanets)
	http.HandleFunc("/players", s.HandlePlayers)
	http.HandleFunc("/next", s.HandleNextTurn)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
