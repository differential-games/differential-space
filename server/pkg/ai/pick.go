package ai

import (
	"sort"

	"github.com/differential-games/differential-space/pkg/ai/strategy"
)

// Pick picks good moves and returns the set of picked attack moves based on the passed strategies.
func Pick(moves []strategy.Move, strats []strategy.Strategy) []strategy.Move {
	for _, strat := range strats {
		strat.Analyze(moves)
	}

	nMoves := 0
	for i := range moves {
		score := 0.0
		for _, strat := range strats {
			score += strat.Score(&moves[i])
		}
		if score >= 0.0 {
			moves[nMoves] = moves[i]
			moves[nMoves].Score = score
			nMoves++
		}
	}

	sort.Slice(moves[:nMoves], func(i, j int) bool {
		mi := moves[i]
		mj := moves[j]
		if (mi.MoveType == strategy.Reinforce) != (mj.MoveType == strategy.Reinforce) {
			return mi.MoveType != strategy.Reinforce
		}
		return mi.Score > mj.Score
	})

	return moves[:nMoves]
}
