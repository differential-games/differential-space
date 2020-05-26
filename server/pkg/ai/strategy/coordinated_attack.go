package strategy

import (
	"github.com/differential-games/differential-space/pkg/game"
)

const coordinatePenalty = 0.25

func CoordinatedAttack(moves []Move) Strategy {
	// Track the moves we want.
	wantFrom := make([]bool, nPlanets)
	var wantFromTo [nPlanets][nPlanets]bool

	// attacks from x to y
	var attacks [nPlanets][nPlanets]bool
	// defends x from y
	var defends [nPlanets][nPlanets]bool

	var totalDamage [nPlanets]float64
	var damages [nPlanets][nPlanets]float64

	for _, m := range moves {
		if m.MoveType != Attack {
			continue
		}

		if wantFrom[m.From] || (totalDamage[m.To] > float64(m.ToStrength+1)) {
			// 1. We've already got an attack from this attacker.
			// 2. We've already got a coordinated attack to this defender.
			continue
		}

		damage := game.WinProbability(m.Distance) * float64(m.FromStrength)
		damages[m.From][m.To] = damage

		totalDamage[m.To] += damage - coordinatePenalty
		attacks[m.From][m.To] = true
		defends[m.To][m.From] = true

		if totalDamage[m.To] > float64(m.ToStrength+1) {
			// Launch the attack!
			for from, isFrom := range defends[m.To] {
				if !isFrom {
					continue
				}

				wantFrom[from] = true
				wantFromTo[from][m.To] = true

				for to, isTo := range attacks[m.From] {
					if !isTo || to == m.To {
						continue
					}
					// Problem: These aren't sorted.
					totalDamage[m.To] -= damages[from][to] + coordinatePenalty
				}
			}

			continue
		}
	}

	return func(move Move) float64 {
		if wantFromTo[move.From][move.To] {
			return 1.0
		}
		return 0.0
	}
}
