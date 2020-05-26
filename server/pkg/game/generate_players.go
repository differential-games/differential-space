package game

import (
	"math"
)

func GeneratePlayers(options PlayerOptions) []Player {
	// Generate players.
	// The first player, index 0, is the environment.
	players := make([]Player, options.NumPlayers+1)
	for i := range players {
		// The environment player does not have technology.
		if i != 0 {
			players[i] = Player{}
		}
	}
	return players
}

// GenerateStarts generates staring Planets for non-Environment Players.
//
// Assumes:
// 1) len(planets) is greater than numPlayers.
// 2) no planets are outside maxRadius.
func GenerateStarts(planets []Planet, maxRadius, numPlayers int) []int {
	result := make([]int, numPlayers)
	// Keep track of already-used starts so two Players don't begin
	// on the same Planet.
	starts := make(map[int]bool)

	angleDiff := 2 * math.Pi / float64(numPlayers)

	for i := 0; i < numPlayers; i++ {
		angle := angleDiff * float64(i)
		x := float64(maxRadius) * math.Sin(angle)
		y := float64(maxRadius) * math.Cos(angle)

		idx := closestTo(x, y, planets, starts)
		starts[idx] = true
		result[i] = idx
	}
	return result
}
