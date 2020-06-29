package strategy

import (
	"math"

	"github.com/differential-games/differential-space/pkg/game"
)

const NPlanets = 60

type ReinforceFront struct {
	maxMoves   int
	reinforces []Move
	distances  []float64
}

func NewReinforceFront(maxMoves int) *ReinforceFront {
	return &ReinforceFront{maxMoves: maxMoves}
}

func (s *ReinforceFront) Initialize(g game.Game) {
	s.reinforces = make([]Move, s.maxMoves)
	s.distances = make([]float64, len(g.Planets))
}

func (s *ReinforceFront) Analyze(moves []Move) {
	for i := range s.distances {
		s.distances[i] = 0
	}

	nReinforces := 0
	// The more potential moves, the higher priority.
	for _, m := range moves {
		if m.MoveType == Reinforce {
			s.reinforces[nReinforces] = m
			nReinforces++
		} else {
			s.distances[m.From] = 1
		}
	}

	// The further away from potential moves, the lower priority.
	nDelete := 0
	reinforces := s.reinforces[:nReinforces]
	for {
		// Keep track of the number deleted each iteration.
		// We aren't guaranteed there is a path to the front from every planet.
		// For example, if a few are isolated from the rest of the galaxy for a few turns.
		nDelete = 0
		for i, r := range reinforces {
			if s.distances[r.From] != 0 {
				// We've already got the distance from the starting planet.
				reinforces[i] = reinforces[nReinforces-1]
				nReinforces--
				nDelete++
			} else if s.distances[r.To] != 0 {
				// We know the distance from the destination planet.
				reinforces[i] = reinforces[nReinforces-1]
				nReinforces--
				nDelete++
				s.distances[r.From] = s.distances[r.To] + 1
			}
		}
		if nDelete == 0 || nReinforces == 0 {
			// We've got the jump distance for each reinforce move to the front.
			// Or any remaining ones have no path to the front.
			break
		}
		reinforces = reinforces[:nReinforces]
	}
}

func (s *ReinforceFront) Score(move *Move) float64 {
	if move.MoveType != Reinforce {
		// Ignore non-reinforce moves.
		return 0.0
	}
	fromDist := s.distances[move.From]
	if fromDist < 2 || fromDist <= s.distances[move.To] {
		// Don't move away from the front line.
		return 0.0
	}
	maxD := math.MaxFloat64
	switch move.FromStrength {
	case 1:
		maxD = 1
	case 2:
		maxD = 2
	case 3:
		maxD = 2
	case 4:
		maxD = 3
	case 5:
		maxD = 3
	case 6:
		maxD = 4
	case 7:
		maxD = 5
	}
	if fromDist > maxD {
		return 0.0
	}
	return 100 - 10*s.distances[move.From] + float64(move.ToStrength)
}
