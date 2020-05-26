package serve

import (
	"sync"

	"github.com/differential-games/differential-space/pkg/game"
)

type Server struct {
	mux sync.Mutex

	options Options

	game *game.Game
}

// New creates a new Server
func New(options Options) (*Server, error) {
	g, err := game.New(options.Options)
	if err != nil {
		return nil, err
	}

	s := &Server{
		options: options,
		game:    g,
	}
	return s, nil
}
