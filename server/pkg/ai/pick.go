package ai

import (
	"sort"

	"github.com/differential-games/differential-space/pkg/ai/strategy"
)

// Pick picks good moves and returns the set of picked attack moves based on the passed strategies.
func Pick(moves []strategy.Move, svbs []strategy.VectorBuilder) []strategy.Move {
	strats := make([]strategy.Vector, len(svbs))
	for i, svb := range svbs {
		strats[i] = strategy.Vector{
			Strategy: svb.Builder(moves),
			Weight:   svb.Weight,
		}
	}

	nMoves := 0
	for i, m := range moves {
		score := 0.0
		for _, strat := range strats {
			score += strat.Weight * strat.Strategy(m)
		}
		if score >= 0.0 {
			moves[i].Score = score
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
