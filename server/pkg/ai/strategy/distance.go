package strategy

import (
	"github.com/differential-games/differential-space/pkg/game"
)

// PreferFurther prefers moves further away from the starting point.
func PreferFurther(move Move) float64 {
	return move.Distance / game.MaxMoveDistance
}

// PreferCloser prefers moves closer to the starting point.
func PreferCloser(move Move) float64 {
	return 1.0 - move.Distance / game.MaxMoveDistance
}
