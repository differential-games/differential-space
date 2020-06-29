package ai

import (
	"github.com/differential-games/differential-space/pkg/ai/strategy"
)

// 0.0 is a baseline that takes random moves.
// 0.00 =    0
// 0.15 =  100
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
		Strategies: []strategy.Strategy{
			strategy.MovePriority{Attack: -1.0, Reinforce: -1.0},
			&strategy.PreferFewerNeighbors{},
			strategy.PreferFurther{},
			strategy.PreferCloser{},
			&strategy.CoordinatedAttack{},
			strategy.AttackIfWinLikely{},
			strategy.NewReinforceFront(1000),
		},
	}
}
