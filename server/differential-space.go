package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/differential-games/differential-space/pkg/serve"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var config = flag.String("config",
	"",
	"path to the Options file")

func readOptions(file string) serve.Options {
	options := serve.DefaultOptions
	if file != "" {
		bytes, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = json.Unmarshal(bytes, &options)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	jsn, _ := json.MarshalIndent(options, "", "  ")
	fmt.Println(string(jsn))
	return options
}

func main() {
	flag.Parse()

	rand.Seed(time.Now().UTC().UnixNano())

	options := readOptions(*config)

	s, err := serve.New(options)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/move", s.HandleMove)
	http.HandleFunc("/move/predict", s.HandleMovePredict)
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
