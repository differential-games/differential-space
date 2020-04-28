package serve

import "github.com/differential-games/differential-space/pkg/game"

type Options struct {
	// Difficulty is how optimally the computer plays.
	// 1.0 is hard.
	// 0.5 is easy.
	// 0.0 the opponent does nothing.
	Difficulty float64 `json:"difficulty"`

	// HumanPlayers is the human-controlled Players.
	// The server does not allow humans to take actions for non-human Players.
	// 0 is the environment and can't take actions.
	HumanPlayers []int `json:"humanPlayers"`

	// Options are the options for creating the game galaxy.
	game.Options
}

// DefaultOptions are the default server options.
var DefaultOptions = Options{
	Difficulty:   0.5,
	HumanPlayers: []int{1},
	Options:      game.DefaultOptions,
}
