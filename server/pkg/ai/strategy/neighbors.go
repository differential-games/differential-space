package strategy

func PreferFewerNeighbors(moves []Move) Strategy {
	colonizers := make(map[int]int)
	max := 1
	for _, m := range moves {
		colonizers[m.To]++
		if colonizers[m.To] > max {
			max = colonizers[m.To]
		}
	}

	return func(move Move) float64 {
		return float64(max) - float64(colonizers[move.To] - 1)
	}
}
