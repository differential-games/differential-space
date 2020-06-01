package strategy

import (
	"github.com/differential-games/differential-space/pkg/game"
)

const coordinatePenalty = 0.25

type CoordinatedAttack struct {
	wantFromTo [nPlanets][nPlanets]bool
}

func (s *CoordinatedAttack) Initialize(game.Game) {}

func (s *CoordinatedAttack) reset() {
	for _, w := range s.wantFromTo {
		for i := range w {
			w[i] = false
		}
	}
}

func (s *CoordinatedAttack) Analyze(moves []Move) {
	s.reset()

	// Track the moves we want.
	wantFrom := make([]bool, nPlanets)

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
				s.wantFromTo[from][m.To] = true

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
}

func (s *CoordinatedAttack) Score(move Move) float64 {
	if s.wantFromTo[move.From][move.To] {
		return 1.0
	}
	return 0.0
}
