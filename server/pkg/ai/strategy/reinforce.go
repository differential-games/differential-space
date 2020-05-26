package strategy

const nPlanets = 60

func ReinforceFront(moves []Move) Strategy {
	reinforces := make([]Move, len(moves))
	nReinforces := 0
	distances := make([]int, nPlanets)
	// The more potential moves, the higher priority.
	for _, m := range moves {
		switch m.MoveType {
		case Reinforce:
			reinforces[nReinforces] = m
			nReinforces++
		default:
			distances[m.From] = 1
		}
	}

	reinforces = reinforces[:nReinforces]
	// The further away from potential moves, the lower priority.
	for nReinforces > 0 {
		var toDelete []int
		for i, r := range reinforces {
			if distances[r.From] != 0 {
				// We've already got the distance from the starting planet.
				toDelete = append(toDelete, i)
				continue
			}
			if distances[r.To] != 0 {
				// We know the distance from the destination planet.
				toDelete = append(toDelete, i)
				distances[r.From] = distances[r.To] + 1
			}
		}
		if len(toDelete) == 0 {
			break
		}
		for _, d := range toDelete {
			reinforces[d] = reinforces[nReinforces-1]
			nReinforces--
		}
		reinforces = reinforces[:nReinforces]
	}

	return func(move Move) float64 {
		if move.MoveType != Reinforce {
			// Ignore non-reinforce moves.
			return 0.0
		}
		if distances[move.From] <= distances[move.To] {
			// Don't move away from the front line.
			return 0.0
		}
		return float64(move.FromStrength) - 1.5*float64(distances[move.From]) + 0.0
	}
}
