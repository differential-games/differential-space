package strategy

import (
	"github.com/differential-games/differential-space/pkg/game"
)

// PreferFurther prefers moves further away from the starting point.
func PreferFurther(t MoveType) Strategy {
	return func(move Move) float64 {
		if t != move.MoveType {
			return 0.0
		}
		return move.Distance * game.InvMaxMoveDistance
	}
}

// PreferCloser prefers moves closer to the starting point.
func PreferCloser(t MoveType) Strategy {
	return func(move Move) float64 {
		if t != move.MoveType {
			return 0.0
		}
		return 1.0 - move.Distance*game.InvMaxMoveDistance
	}
}
