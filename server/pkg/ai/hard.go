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
		Strategies: []strategy.Strategy{
			strategy.NewMovePriority(0.0, -1.0, -1.0),
			&strategy.PreferFewerNeighbors{},
			strategy.PreferFurther{},
			strategy.PreferCloser{},
			strategy.NewAttackAtStrength(game.MaxFleetSize),
			&strategy.CoordinatedAttack{},
			strategy.AttackIfWinLikely{},
			strategy.NewReinforceFront(1000),
		},
	}
}
