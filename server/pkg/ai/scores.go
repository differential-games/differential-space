package ai

import "github.com/differential-games/differential-space/pkg/game"

func Scores(nPlayers int, planets []game.Planet) []int {
	result := make([]int, nPlayers+1)
	for _, p := range planets {
		if p.Owner != 0 {
			result[p.Owner]++
		}
	}
	return result
}

func Maxes(left, right []int) []int {
	for i, l := range left {
		if right[i] > l {
			left[i] = l
		}
	}
	return left
}

func MaxScore(scores []int) int {
	result := 0
	for _, s := range scores {
		if s > result {
			result = s
		}
	}
	return result
}
