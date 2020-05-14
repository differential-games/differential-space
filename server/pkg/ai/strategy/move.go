package strategy

import (
	"github.com/differential-games/differential-space/pkg/game"
)

// Move is a superset of game.Move which includes all necessary metadata
// for the AI To choose optimal moves.
type Move struct {
	game.Move
	// Distance is the distance between the two Planets.
	Distance float64
	// FromStrength is the strength of the Planet being moved from.
	FromStrength int
	// ToStrength is the strength of the Planet being moved to.
	ToStrength int
}
