package strategy

import "github.com/differential-games/differential-space/pkg/game"

type Strategy interface {
	// Initialize performs any once-per-game initialization for the Strategy.
	Initialize(game game.Game)

	// Anaylze analyzed the set of available Moves for a given turn.
	Analyze(moves []Move)

	// Score returns how good the Strategy evaluates the Move to be.
	// Usually on the range [0,+1].
	//
	// Positive indicates signal to do the move.
	// 0.0 indicates no signal to do the move.
	// Negative indicates signal to not do the move.
	Score(move *Move) float64
}
