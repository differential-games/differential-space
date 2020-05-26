package game

import (
	"fmt"
	"math"
)

func Analyze(attack, defend int, p float64) string {
	if attack < defend+1 {
		// Can't win, just return expected damage.
		// Expect is the expected damage done to the defender.
		expect := int(math.Min(float64(attack)*p, float64(defend)))
		return fmt.Sprintf("Attack: expected damage: %d", expect)
	}

	winChance := BattleWinChance(attack, defend, p)
	return fmt.Sprintf("Attack: win chance: %.2f%%", winChance*100)
}

func WinLikely(attack, defend int, p float64) bool {
	return float64(attack)*p > float64(defend)
}

func factorial(n int) float64 {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return float64(result)
}

func nCr(n, k int) float64 {
	return factorial(n) / (factorial(k) * factorial(n-k))
}

func BattleWinChance(attack, defend int, p float64) float64 {
	winChance := float64(0)
	for i := defend + 1; i <= attack; i++ {
		delta := math.Pow(p, float64(i)) * math.Pow(1-p, float64(attack-i)) * nCr(attack, i)
		winChance += delta
	}
	return winChance
}
