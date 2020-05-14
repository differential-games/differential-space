package strategy

import (
	"github.com/differential-games/differential-space/pkg/game"
)

const coordinatePenalty = 0.25

type FromTo struct {
	From, To int
}

func CoordinatedAttack(moves []Move) Strategy {
	// possibleAttacks it a map from a possible defender to the list of potential attackers
	// and the damage the attack is expected to do.
	defenderToAttackerToDamage := make(map[int]map[int]float64)
	// possibleAttacks it a map from a possible defender to the list of potential attackers
	// and the damage the attack is expected to do.
	attackerToDefenderToDamage := make(map[int]map[int]float64)

	// defenderToDamage is the potential damage done to each targer.
	defenderToDamage := make(map[int]float64)

	// track the moves we want.
	wantFromTo := make(map[FromTo]bool)

	for _, m := range moves {
		attackScore := float64(m.ToStrength + 1) + coordinatePenalty * (1.0 + float64(len(defenderToAttackerToDamage[m.To])))
		if defenderToDamage[m.To] > attackScore {
			// We've already got a coordinated attack to this defender.
			continue
		}

		expectedDamage := game.WinProbability(m.Distance) * float64(m.FromStrength)

		defenderToDamage[m.To] += expectedDamage

		defend := defenderToAttackerToDamage[m.To]
		if defend == nil {
			defend = make(map[int]float64)
		}
		defend[m.From] = expectedDamage
		defenderToAttackerToDamage[m.To] = defend

		attack := attackerToDefenderToDamage[m.To]
		if attack == nil {
			attack = make(map[int]float64)
		}
		attack[m.From] = expectedDamage
		attackerToDefenderToDamage[m.To] = attack

		attackScore = float64(m.ToStrength + 1) + coordinatePenalty * (1.0 + float64(len(defenderToAttackerToDamage[m.To])))
		if defenderToDamage[m.To] > attackScore {
			// Launch the attack!
			// Send all attackers
			for attacker := range defenderToAttackerToDamage[m.To] {
				wantFromTo[FromTo{
					From: attacker,
					To: m.To,
				}] = true

				// Subtract damage from other defenders.
				for defender := range attackerToDefenderToDamage[attacker] {
					if defender == m.To {
						continue
					}
					defenderToDamage[defender] -= attackerToDefenderToDamage[attacker][defender]
				}
			}
		}
	}

	return func(move Move) float64 {
		if wantFromTo[FromTo{
			From: move.From,
			To:   move.To,
		}] {
			return 1.0
		}
		return 0.0
	}
}
