package serve

import (
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

	p, err := s.game.AttackProbability(move)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	result := AttackProbability{P: p}
	writeJson(w, result)
}
