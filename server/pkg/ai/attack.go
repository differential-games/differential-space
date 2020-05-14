package ai

import (
	"github.com/differential-games/differential-space/pkg/ai/strategy"
	"sort"
)

// Pick picks good moves and returns the set of picked attack moves based on the passed strategies.
func Pick(moves []strategy.Move, svbs []strategy.VectorBuilder) []strategy.Move {
	var strats []strategy.Vector
	for _, svb := range svbs {
		strats = append(strats, strategy.Vector{
			Strategy: svb.Builder(moves),
			Weight:   svb.Weight,
		})
	}

	var result []strategy.Move
	scores := make(map[strategy.Move]float64, len(moves))
	for _, attack := range moves {
		score := 0.0
		for _, strat := range strats {
			score += strat.Weight * strat.Strategy(attack)
		}
		if score >= 1.0 {
			scores[attack] = score
			result = append(result, attack)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return scores[result[i]] > scores[result[j]]
	})

	return result
}
