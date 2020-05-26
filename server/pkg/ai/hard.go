package ai

import (
	"github.com/differential-games/differential-space/pkg/ai/strategy"
	"github.com/differential-games/differential-space/pkg/game"
)

// 0.0 is a baseline that takes random moves.
// 0.0 =    0
// 0.1 =  480
// 0.2 =  846
// 0.3 = 1042
// 0.4 = 1179
// 0.5 = 1277
// 0.6 = 1374
// 0.7 = 1443
// 0.8 = 1499
// 0.9 = 1541
// 1.0 = 1575

// Hard returns an AI using the most-optimized strategies.
func Hard(difficulty float64) Interface {
	return &AI{
		Difficulty: difficulty,
		Strategies: []strategy.VectorBuilder{
			{
				Builder: strategy.Builder(strategy.MovePriority(
					2.0, 0.0, 0.0,
				)),
			},
			{
				Builder: strategy.PreferFewerNeighbors,
				Weight:  1.0,
			},
			{
				Builder: strategy.Builder(strategy.HasStrength(2)),
				Weight:  1.0,
			},
			{
				Builder: strategy.Builder(strategy.PreferFurther(strategy.Colonize)),
				Weight:  1.0,
			},
			{
				Builder: strategy.Builder(strategy.PreferCloser(strategy.Attack)),
				Weight:  1.0,
			},
			{
				Builder: strategy.Builder(strategy.AttackAtStrength(game.MaxFleetSize)),
				Weight:  1.0,
			},
			{
				Builder: strategy.Builder(strategy.AttackIfWinLikely()),
				Weight:  1.0,
			},
			{
				Builder: strategy.ReinforceFront,
				Weight:  1.0,
			},
		},
	}
}
