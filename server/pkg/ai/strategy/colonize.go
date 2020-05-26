package strategy

func containsColonize(moves []Move) bool {
	for i := range moves {
		if moves[i].MoveType == Colonize {
			return true
		}
	}
	return false
}

func HasStrength(s int) Strategy {
	return func(move Move) float64 {
		if move.MoveType != Colonize {
			return 0.0
		}
		if move.FromStrength < s {
			return -1.0
		}
		return 1.0
	}
}

func PreferFewerNeighbors(moves []Move) Strategy {
	if !containsColonize(moves) {
		return func(move Move) float64 {
			return 0.0
		}
	}

	colonizers := make([]int, 60)
	max := 1
	for _, m := range moves {
		colonizers[m.To]++
		if colonizers[m.To] > max {
			max = colonizers[m.To]
		}
	}

	return func(move Move) float64 {
		if move.MoveType != Colonize {
			return 0.0
		}
		return float64(max - colonizers[move.To] + 1)
	}
}
