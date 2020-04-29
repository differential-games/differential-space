package serve

import (
	"fmt"
	"github.com/differential-games/differential-space/pkg/game"
	"net/http"
)

func (s *Server) HandleMove(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "not allowed", http.StatusMethodNotAllowed)
		return
	}

	move := game.Move{}
	putJson(w, r.Body, &move)
	isHuman := false
	for _, p := range s.options.HumanPlayers {
		if p == s.game.Planets[move.From].Owner {
			isHuman = true
		}
	}
	if !isHuman {
		http.Error(w, "attempted action for AI-controlled player", http.StatusBadRequest)
		return
	}

	err := s.game.Move(move)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

type AttackProbability struct {
	P float64
}

func (s *Server) HandleMovePredict(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "not allowed", http.StatusMethodNotAllowed)
		return
	}

	move := game.Move{}
	putJson(w, r.Body, &move)

	if s.game.PlayerTurn == 0 {
		http.Error(w, "This game has not started yet.", http.StatusBadRequest)
		return
	}
	if s.game.Planets[move.From].Owner == 0 {
		http.Error(w, "This planet is unowned.", http.StatusBadRequest)
		return
	}
	isHuman := false
	for _, p := range s.options.HumanPlayers {
		if p == s.game.Planets[move.From].Owner {
			isHuman = true
		}
	}
	if !isHuman {
		http.Error(w, "You can't act for an AI-controlled player.", http.StatusBadRequest)
		return
	}

	message, valid := s.game.AssessAttack(move)
	if !valid {
		http.Error(w, message, http.StatusBadRequest)
		return
	}

	_, err := w.Write([]byte(message))
	if err != nil {
		fmt.Println(err)
	}
}
