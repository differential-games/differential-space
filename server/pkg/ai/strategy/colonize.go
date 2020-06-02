package strategy

import (
	"github.com/differential-games/differential-space/pkg/game"
)

func containsColonize(moves []Move) bool {
	for i := range moves {
		if moves[i].MoveType == Colonize {
			return true
		}
	}
	return false
}

type PreferFewerNeighbors struct {
	colonizers       []int
	max              int
	containsColonize bool
}

func (s *PreferFewerNeighbors) Initialize(g game.Game) {
	s.colonizers = make([]int, len(g.Planets))
}

func (s *PreferFewerNeighbors) Analyze(moves []Move) {
	s.containsColonize = containsColonize(moves)
	if !s.containsColonize {
		return
	}

	for i := range s.colonizers {
		s.colonizers[i] = 0
	}

	s.max = 1
	for _, m := range moves {
		if m.MoveType != Colonize {
			continue
		}
		s.colonizers[m.To]++
		if s.colonizers[m.To] > s.max {
			s.max = s.colonizers[m.To]
		}
	}
}

func (s *PreferFewerNeighbors) Score(move *Move) float64 {
	if move.MoveType != Colonize {
		return 0.0
	}
	return float64(s.max - s.colonizers[move.To] + 1)
}
