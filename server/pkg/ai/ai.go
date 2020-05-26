package ai

import (
	"github.com/differential-games/differential-space/pkg/ai/strategy"
	"github.com/differential-games/differential-space/pkg/game"
)

type Interface interface {
	PickMoves(player int, planets []game.Planet, moves []strategy.Move) []strategy.Move
}

type AI struct {
	Difficulty float64

	Strategies []strategy.VectorBuilder
}

var _ Interface = &AI{}

func filterWorst(moves []strategy.Move, difficulty float64) []strategy.Move {
	nColonize := 0
	nAttack := 0
	nReinforce := 0
	for _, m := range moves {
		switch m.MoveType {
		case strategy.Colonize:
			nColonize++
		case strategy.Attack:
			nAttack++
		case strategy.Reinforce:
			nReinforce++
		}
	}

	nColonize = nColonize * int(difficulty*100) / 100
	nAttack = nAttack * int(difficulty*100) / 100
	nReinforce = nReinforce * int(difficulty*100) / 100
	result := make([]strategy.Move, nColonize+nAttack+nReinforce)

	nResult := 0
	for _, m := range moves {
		switch m.MoveType {
		case strategy.Colonize:
			if nColonize > 0 {
				result[nResult] = m
				nResult++
				nColonize--
			}
		case strategy.Attack:
			if nAttack > 0 {
				result[nResult] = m
				nResult++
				nAttack--
			}
		case strategy.Reinforce:
			if nReinforce > 0 {
				result[nResult] = m
				nResult++
				nReinforce--
			}
		}
	}

	return result
}

// PickMoves executes the AI algorithm for a Player
func (ai *AI) PickMoves(player int, planets []game.Planet, moves []strategy.Move) []strategy.Move {
	moves = strategy.GenerateMoves(player, planets, moves)
	moves = Pick(moves, ai.Strategies)
	return filterWorst(moves, ai.Difficulty)
}
