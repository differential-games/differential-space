package serve

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/differential-games/differential-space/pkg/ai/strategy"

	"github.com/differential-games/differential-space/pkg/ai"
)

func (s *Server) HandleNextTurn(w http.ResponseWriter, r *http.Request) {
	s.mux.Lock()
	defer s.mux.Unlock()

	switch r.Method {
	case http.MethodPost:
		startPlayer := s.game.PlayerTurn
		nextPlayer, noPlanets := s.game.NextTurn()
		for noPlanets && s.game.PlayerTurn != startPlayer {
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

		aiPlayer := ai.Hard(s.options.Difficulty)
		aiPlayer.Initialize(*s.game)

		moves := make([]strategy.Move, 1000)
		moves = aiPlayer.PickMoves(player, s.game.Planets, moves)

		for _, m := range moves {
			if !s.game.Planets[m.From].Ready {
				continue
			}
			to := s.game.Planets[m.To]
			if (to.Strength == 8) && (to.Owner == player) {
				// Don't bother reinforcing an already-full planet.
				continue
			}

			err := s.game.Move(m.Move)
			if err != nil {
				http.Error(w, "AI tried invalid move", http.StatusInternalServerError)
				return
			}
		}

	default:
		http.Error(w, "not allowed", http.StatusMethodNotAllowed)
	}
}
