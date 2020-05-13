package ai

import "github.com/differential-games/differential-space/pkg/game"

type FilterResult int

const (
	NoResult FilterResult = iota
	Keep
	Discard
)

type MoveFilterBuilder func([]Move) MoveFilter

type MoveFilter func(Move) FilterResult

func AttackAtMaxStrength([]Move) MoveFilter {
	return func(move Move) FilterResult {
		if move.FromStrength == 8 {
			return Keep
		}
		return NoResult
	}
}

type FromTo struct {
	From, To int
}

// diff is the penalty for larger numbers of attackers to overcome the
func CoordinatedAttackAtDiff(diff1, diff2 float64) MoveFilterBuilder {
	// Assumes incoming moves are sorted closest-first.

	return func(moves []Move) MoveFilter {
		// possibleAttacks it a map from a possible defender to the list of potential attackers
		// and the damage the attack is expected to do.
		defenderToAttackerToDamage := make(map[int]map[int]float64)
		// possibleAttacks it a map from a possible defender to the list of potential attackers
		// and the damage the attack is expected to do.
		attackerToDefenderToDamage := make(map[int]map[int]float64)

		// defenderToDamage is the potential damage done to each targer.
		defenderToDamage := make(map[int]float64)

		wantFromTo := make(map[FromTo]bool)

		for _, m := range moves {
			if defenderToDamage[m.To] > float64(m.ToStrength + 1) + diff1 + diff2 * float64(len(defenderToAttackerToDamage[m.To])) {
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

			if defenderToDamage[m.To] > float64(m.ToStrength + 1) + diff1 + diff2 * float64(len(defenderToAttackerToDamage[m.To])) {
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

		return func(move Move) FilterResult {
			if wantFromTo[FromTo{
				From: move.From,
				To:   move.To,
			}] {
				return Keep
			}
			return NoResult
		}
	}
}

func AttackAtProbability(min float64) MoveFilterBuilder {
	return func(moves []Move) MoveFilter {
		return func(move Move) FilterResult {
			p := game.BattleWinChance(move.FromStrength, move.ToStrength, move.Distance)
			if p >= min {

				return Keep
			}
			return Discard
		}
	}
}
