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
func Hard(nPlanets int, difficulty float64) Interface {
	return &AI{
		Difficulty: difficulty,
		Strategies: []strategy.Vector{
			NewVector(1.0, strategy.NewMovePriority(0.0, -1.0, -1.0)),
			NewVector(1.0, &strategy.PreferFewerNeighbors{}),
			NewVector(1.0, strategy.PreferFurther{}),
			NewVector(1.0, strategy.PreferCloser{}),
			NewVector(1.0, strategy.NewAttackAtStrength(game.MaxFleetSize)),
			NewVector(1.0, &strategy.CoordinatedAttack{}),
			NewVector(1.0, strategy.AttackIfWinLikely{}),
			NewVector(1.0, strategy.NewReinforceFront(1000)),
		},
	}
}

func NewVector(weight float64, s strategy.Strategy) strategy.Vector {
	return strategy.Vector{
		Strategy: s,
		Weight:   weight,
	}
}
