package game

import (
	"math"
	"math/rand"
)

func Fight(attacker, defender *Planet) bool {
	p := WinProbability(math.Sqrt(DistSq(attacker.X, attacker.Y, defender.X, defender.Y)))
	for ; attacker.Strength > 0 && defender.Strength >= 0; attacker.Strength-- {
		// A battle takes place in rounds, up to to the strength of the attacker.
		if rand.Float64() < p {
			// Each round, the attacker has p probability of doing a damage to the enemy.
			defender.Strength--
		}
	}
	// The attacker won if they reduced the defender strength to 0, and then won one more attack.
	if defender.Strength < 0 {
		// The attacker won; make sure to set the defender strength to 0.
		defender.Strength = 0
		attacker.Strength++
		return true
	}
	return false
}
