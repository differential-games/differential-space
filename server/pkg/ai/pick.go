package ai

import (
	"sort"

	"github.com/differential-games/differential-space/pkg/ai/strategy"
)

// Pick picks good moves and returns the set of picked attack moves based on the passed strategies.
func Pick(moves []strategy.Move, strats []strategy.Vector) []strategy.Move {
	for _, svb := range strats {
		svb.Strategy.Analyze(moves)
	}

	nMoves := 0
	for i, m := range moves {
		score := 0.0
		for _, strat := range strats {
			score += strat.Weight * strat.Strategy.Score(m)
		}
		moves[i].Score = score
		if score >= 0.0 {
			nMoves++
		}
	}

	sort.Slice(moves, func(i, j int) bool {
		mi := moves[i]
		mj := moves[j]
		if (mi.Score < 0.0) != (mj.Score < 0.0) {
			return mj.Score < 0.0
		}
		if (mi.MoveType == strategy.Reinforce) != (mj.MoveType == strategy.Reinforce) {
			return mi.MoveType != strategy.Reinforce
		}
		return mi.Score > mj.Score
	})

	return moves[:nMoves]
}
