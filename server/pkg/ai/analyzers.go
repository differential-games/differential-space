package ai

import "github.com/differential-games/differential-space/pkg/game"

type MoveAnalyzer func(colonizes []Move) func(left, right Move) Priority

type MovePriority func(left, right Move) Priority

func PreferCloser(colonizes []Move) func(left, right Move) Priority {
	return func(left, right Move) Priority {
		if left.Distance < right.Distance {
			return Left
		}
		return Right
	}
}

func PreferFurther(colonizes []Move) func(left, right Move) Priority {
	return func(left, right Move) Priority {
		if left.Distance < right.Distance {
			return Right
		}
		return Left
	}
}

func PreferFewerNeighbors(colonizes []Move) func(left, right Move) Priority {
	colonizers := make(map[int]int)
	for _, m := range colonizes {
		colonizers[m.To]++
	}

	return func(left, right Move) Priority {
		if colonizers[left.To] < colonizers[right.To] {
			return Left
		}
		if colonizers[right.To] < colonizers[left.To] {
			return Right
		}
		return None
	}
}

func PreferMoreNeighbors(colonizes []Move) func(left, right Move) Priority {
	colonizers := make(map[int]int)
	for _, m := range colonizes {
		colonizers[m.To]++
	}

	return func(left, right Move) Priority {
		if colonizers[left.To] > colonizers[right.To] {
			return Left
		}
		if colonizers[right.To] > colonizers[left.To] {
			return Right
		}
		return None
	}
}

func PreferMaxStrength(moves []Move) func(left, right Move) Priority {
	return func(left, right Move) Priority {
		if left.FromStrength != right.FromStrength {
			if left.FromStrength == 8 {
				return Left
			}
			if right.FromStrength == 8 {
				return Right
			}
		}
		return None
	}
}

func PreferProbableWin(moves []Move) func(left, right Move) Priority {
	return func(left, right Move) Priority {
		leftScore := game.BattleWinChance(left.FromStrength, left.To, left.Distance)
		rightScore := game.BattleWinChance(right.FromStrength, right.To, right.Distance)
		if leftScore < rightScore {
			return Right
		}
		return Left
	}
}
