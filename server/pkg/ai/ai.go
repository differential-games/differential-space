package ai

import (
	"github.com/differential-games/differential-space/pkg/ai/strategy"
	"github.com/differential-games/differential-space/pkg/game"
)

type Interface interface {
	Initialize(game.Game)

	PickMoves(player int, planets []game.Planet, moves []strategy.Move) []strategy.Move
}

type AI struct {
	Difficulty float64

	Strategies []strategy.Strategy
}

var _ Interface = &AI{}

func (ai *AI) Initialize(g game.Game) {
	for _, s := range ai.Strategies {
		s.Initialize(g)
	}
}

func filterWorst(moves []strategy.Move, difficulty float64) []strategy.Move {
	if difficulty >= 1.0 {
		return moves
	}

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

	nResult := 0
	for _, m := range moves {
		switch m.MoveType {
		case strategy.Colonize:
			if nColonize > 0 {
				moves[nResult] = m
				nResult++
				nColonize--
			}
		case strategy.Attack:
			if nAttack > 0 {
				moves[nResult] = m
				nResult++
				nAttack--
			}
		case strategy.Reinforce:
			if nReinforce > 0 {
				moves[nResult] = m
				nResult++
				nReinforce--
			}
		}
	}

	return moves[:nResult]
}

// PickMoves executes the AI algorithm for a Player
func (ai *AI) PickMoves(player int, planets []game.Planet, moves []strategy.Move) []strategy.Move {
	moves = strategy.GenerateMoves(player, planets, moves)
	moves = Pick(moves, ai.Strategies)
	return filterWorst(moves, ai.Difficulty)
}
