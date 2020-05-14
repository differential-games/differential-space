package strategy

import (
	"github.com/differential-games/differential-space/pkg/game"
)

func Strength(move Move) float64 {
	return float64(move.FromStrength) / float64(game.MaxFleetSize)
}

func AttackAtStrength(strength int) Strategy {
	return func(move Move) float64 {
		if move.FromStrength >= strength {
			return 1.0
		}
		return 0.0
	}
}

func Probability(move Move) float64 {
	return game.BattleWinChance(move.FromStrength, move.ToStrength, move.Distance)
}

func AttackAtProbability(min float64) Strategy {
	return func(move Move) float64 {
		p := game.BattleWinChance(move.FromStrength, move.ToStrength, move.Distance)
		if p >= min {
			return 1.0
		}
		return 0.0
	}
}
