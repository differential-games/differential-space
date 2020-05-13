package aitest

import (
	"github.com/differential-games/differential-space/pkg/ai"
	"github.com/differential-games/differential-space/pkg/game"
	"testing"
)

func Turn(t *testing.T, g *game.Game, player ai.Interface) {
	g.NextTurn()
	playerId := g.PlayerTurn

	moves := player.PickMoves(playerId, g.Planets)

	for _, m := range moves {
		err := g.Move(m)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func RunGame(g *game.Game, turn func(g *game.Game) bool) {
	for ; !turn(g) ; {}
}
