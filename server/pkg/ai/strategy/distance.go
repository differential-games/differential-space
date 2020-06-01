package strategy

import (
	"github.com/differential-games/differential-space/pkg/game"
)

type PreferFurther struct{}

func (s PreferFurther) Initialize(game.Game) {}

func (s PreferFurther) Analyze(moves []Move) {}

func (s PreferFurther) Score(move Move) float64 {
	if move.MoveType != Colonize {
		return 0.0
	}
	return move.Distance * game.InvMaxMoveDistance
}

type PreferCloser struct{}

func (s PreferCloser) Initialize(game.Game) {}

func (s PreferCloser) Analyze(moves []Move) {}

func (s PreferCloser) Score(move Move) float64 {
	if move.MoveType != Attack {
		return 0.0
	}
	return 1.0 - move.Distance*game.InvMaxMoveDistance
}
