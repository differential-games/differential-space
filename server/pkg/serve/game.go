package serve

import (
	"net/http"

	"github.com/differential-games/differential-space/pkg/game"
)

func (s *Server) HandleGame(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		s.createGame(w)
	default:
		http.Error(w, "not allowed", http.StatusMethodNotAllowed)
	}
}

func (s *Server) createGame(w http.ResponseWriter) {
	g, err := game.New(s.options.Options)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s.game = g
}
