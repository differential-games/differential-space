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

type PlayerStart struct {
	Owner  int
	Starts []Start
}

type Start struct {
	Strength int
}

// GenerateStarts generates staring Planets for non-Environment Players.
//
// Assumes:
// 1) len(planets) is greater than numPlayers.
// 2) no planets are outside maxRadius.
func GenerateStarts(planets []Planet, maxRadius int, starts []PlayerStart) {
	// Keep track of already-used starts so two Players don't begin
	// on the same Planet.
	used := make(map[int]bool)

	angleDiff := 2 * math.Pi / float64(len(starts))

	for i, pStart := range starts {
		angle := angleDiff * float64(i)
		x := float64(maxRadius) * math.Sin(angle)
		y := float64(maxRadius) * math.Cos(angle)

		for _, start := range pStart.Starts {
			idx := closestTo(x, y, planets, used)
			planets[idx].Owner = pStart.Owner
			planets[idx].Strength = start.Strength

			used[idx] = true
		}
	}
}
