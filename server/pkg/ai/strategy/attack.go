package strategy

import (
	"github.com/differential-games/differential-space/pkg/game"
)

type AttackAtStrength struct {
	minStrength int
}

func NewAttackAtStrength(minStrength int) AttackAtStrength {
	return AttackAtStrength{minStrength: minStrength}
}

func (s AttackAtStrength) Initialize(game.Game) {}

func (s AttackAtStrength) Analyze([]Move) {}

func (s AttackAtStrength) Score(move Move) float64 {
	if move.MoveType != Attack {
		return 0.0
	}
	if move.FromStrength >= s.minStrength {
		return 1.0
	}
	return 0.0
}

type AttackIfWinLikely struct{}

func (s AttackIfWinLikely) Initialize(game.Game) {}

func (s AttackIfWinLikely) Analyze([]Move) {}

func (s AttackIfWinLikely) Score(move Move) float64 {
	if move.MoveType != Attack {
		return 0.0
	}
	if game.WinLikely(move.FromStrength, move.ToStrength, game.WinProbability(move.Distance)) {
		return 1.0
	}
	return 0.0
}
