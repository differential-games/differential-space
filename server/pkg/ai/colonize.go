package ai

import (
	"math/rand"
	"sort"
)

type Priority int

const (
	None Priority = iota
	Left
	Right
)

// PickColonization chooses a set of good colonization moves from the set of possible moves.
func PickColonization(moves []Move, analyzers []MoveAnalyzer) ([]Move, []Move) {
	var colonize []Move
	var nonColonize []Move
	for _, m := range moves {
		if m.MoveType == Colonize {
			colonize = append(colonize, m)
		} else {
			nonColonize = append(nonColonize, m)
		}
	}

	var priorities []MovePriority
	for _, analyzer := range analyzers {
		priorities = append(priorities, analyzer(moves))
	}

	// Sort moves from greatest to least distance.
	// We prefer spreading out as much as possible when colonizing.
	sort.Slice(colonize, func(i, j int) bool {
		for _, priority := range priorities {
			p := priority(colonize[i], colonize[j])
			switch p {
			case Left: return true
			case Right: return false
			}
		}

		// Choose randomly.
		return rand.Float64() < 0.5
	})

	// Track optimal colony moves.
	// Colonization always succeeds, so use the Planets least likely to
	// be able to do anything else.
	// Generally this will result in almost the largest number of colonization
	// moves possible.
	from := make(map[int]bool)
	to := make(map[int]bool)
	var colonizes []Move
	for _, m := range colonize {
		if from[m.From] || to[m.To] {
			// We're already launching from the source planet or colonizing the target.
			continue
		}
		from[m.From] = true
		from[m.To] = true
		colonizes = append(colonizes, m)
	}

	// Keep moves that aren't colonize Moves separate.
	var otherMoves []Move
	for _, m := range nonColonize {
		// Ignore all other moves from/to colonizing.
		if from[m.From] || to[m.To] {
			continue
		}
		otherMoves = append(otherMoves, m)
	}

	return colonizes, otherMoves
}
