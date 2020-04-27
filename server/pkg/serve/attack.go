package serve

import (
	"github.com/differential-games/differential-space/pkg/game"
	"net/http"
)

func (s *Server) HandleAttack(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "not allowed", http.StatusMethodNotAllowed)
		return
	}

	attack := game.Move{}
	putJson(w, r.Body, &attack)

	err := s.Move(attack)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

type AttackProbability struct {
	P float64
}

func (s *Server) HandleAttackPredict(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "not allowed", http.StatusMethodNotAllowed)
		return
	}

	attack := game.Move{}
	putJson(w, r.Body, &attack)

	p, err := s.AttackProbability(attack)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	result := AttackProbability{P: p}
	writeJson(w, result)
}
