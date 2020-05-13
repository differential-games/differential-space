package ai

import (
	"github.com/differential-games/differential-space/pkg/game"
	"math/rand"
)

type Random struct {}

func (r Random) PickMoves(player int, planets []game.Planet) []game.Move {
	possible := GenerateMoves(player, planets)

	movesByPlanet := make(map[int][]Move)
	for _, m := range possible {
		movesByPlanet[m.From] = append(movesByPlanet[m.From], m)
	}

	movedTo := make(map[int]bool)
	var moves []game.Move
	for planet, fromMoves := range movesByPlanet {
		if movedTo[planet] {
			continue
		}

		idx := rand.Int() % (len(fromMoves) + 1)
		if idx >= len(fromMoves) {
			continue
		}

		moves = append(moves, fromMoves[idx].Move)
		movedTo[fromMoves[idx].To] = true
	}
	return moves
}
