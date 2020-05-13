package aitest

import "github.com/differential-games/differential-space/pkg/game"

func Scores(planets []game.Planet) []int {
	// TODO: Don't hardcode this length.
	result := make([]int, 17)
	for _, p := range planets {
		if p.Owner != 0 {
			result[p.Owner]++
		}
	}
	return result
}

func Maxes(left, right []int) []int {
	// TODO: Don't hardcode this length.
	result := make([]int, 17)
	for i := range left {
		if left[i] > right[i] {
			result[i] = left[i]
		} else {
			result[i] = right[i]
		}
	}
	return result
}

func MaxScore(scores []int) int {
	result := scores[0]
	for _, s := range scores {
		if s > result {
			result = s
		}
	}
	return result
}

