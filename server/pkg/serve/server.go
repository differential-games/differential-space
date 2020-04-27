package serve

import (
	"github.com/differential-games/differential-space/pkg/game"
	"sync"
)

type Server struct {
	mux sync.Mutex
	*game.Game
}
