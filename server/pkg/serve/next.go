package serve

import (
	"fmt"
	"github.com/differential-games/differential-space/pkg/ai"
	"net/http"
	"strconv"
)

func (s *Server) HandleNextTurn(w http.ResponseWriter, r *http.Request) {
	s.mux.Lock()
	defer s.mux.Unlock()

	switch r.Method {
	case http.MethodPost:
		startPlayer := s.game.PlayerTurn
		nextPlayer, noPlanets := s.game.NextTurn()
		for ; noPlanets && s.game.PlayerTurn != startPlayer; {
			nextPlayer, noPlanets = s.game.NextTurn()
		}

		_, err := w.Write([]byte(strconv.FormatInt(int64(nextPlayer), 10)))
		if err != nil {
			fmt.Println(err)
		}

		player := s.game.PlayerTurn

		isHuman := false
		for _, p := range s.options.HumanPlayers {
			if player == p {
				isHuman = true
				break
			}
		}

		if isHuman {
			// This is a Human player.
			return
		}

		moves := ai.PickMoves(player, s.game.Planets, s.options.Difficulty)
		for _, m := range moves {
			err := s.game.Move(m)
			if err != nil {
				http.Error(w, "AI tried invalid move", http.StatusInternalServerError)
				return
			}
		}

	default:
		http.Error(w, "not allowed", http.StatusMethodNotAllowed)
	}
}
