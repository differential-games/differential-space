package serve

import (
	"github.com/differential-games/differential-space/pkg/ai"
	"github.com/differential-games/differential-space/pkg/game"
	"net/http"
)

func (s *Server) HandleNextTurn(w http.ResponseWriter, r *http.Request) {
	s.mux.Lock()
	defer s.mux.Unlock()

	switch r.Method {
	case http.MethodPost:
		s.Game.NextTurn()
		player := s.PlayerTurn
		if player == 1 {
			// This is the Human player.
			// return
		}

		moves := ai.PickMoves(player, s.Planets, game.DefaultOptions.Difficulty)
		for _, m := range moves {
			err := s.Move(m)
			if err != nil {
				http.Error(w, "AI tried invalid move", http.StatusInternalServerError)
				return
			}
		}
	default:
		http.Error(w, "not allowed", http.StatusMethodNotAllowed)
	}
}
