package strategy

import (
	"math"

	"github.com/differential-games/differential-space/pkg/game"
)

// GenerateMoves generates all possible Moves for a Player.
func GenerateMoves(player int, planets []game.Planet, moves []Move) []Move {
	nMoves := 0
	for i, from := range planets {
		if !from.Ready || player != from.Owner {
			// Can't start move from someone else's Planet.
			// Can't start move from a Planet that isn't Ready or has no troops.
			continue
		}
		for j, to := range planets {
			if i == j {
				// Can't move from a Planet to itself.
				continue
			}
			dsq := game.DistSq(from.X, from.Y, to.X, to.Y)
			if dsq > game.MaxMoveDistanceSq {
				// Moves over 5 units are invalid.
				continue
			}
			// Construct the potential Move and record the relevant metadata.
			mt := Attack
			switch {
			case from.Strength == 0:
				mt = Invalid
			case to.Owner == 0:
				mt = Colonize
			case to.Owner == player:
				mt = Reinforce
			}
			moves[nMoves] = Move{
				MoveType: mt,
				Move: game.Move{
					From: i,
					To:   j,
				},
				Distance:     math.Sqrt(dsq),
				FromStrength: from.Strength,
				ToStrength:   to.Strength,
			}
			nMoves++
		}
	}
	return moves[:nMoves]
}
