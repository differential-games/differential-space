package ai

import (
	"github.com/differential-games/differential-space/pkg/ai/strategy"
	"github.com/differential-games/differential-space/pkg/game"
)

func Turn(g *game.Game, player Interface, moves []strategy.Move) error {
	g.NextTurn()
	playerId := g.PlayerTurn

	picked := player.PickMoves(playerId, g.Planets, moves)

	for _, m := range picked {
		from := g.Planets[m.From]
		if !from.Ready || (from.Strength == 0) {
			continue
		}
		to := g.Planets[m.To]
		if (to.Strength == 8) && (to.Owner == playerId) {
			// Don't bother reinforcing an already-full planet.
			continue
		}
		err := g.Move(m.Move)
		if err != nil {
			return err
		}
	}
	return nil
}

func RunGame(g *game.Game, turn func(g *game.Game) (bool, error)) error {
	done := false
	var err error
	for ; !done; done, err = turn(g) {
		if err != nil {
			return err
		}
	}

	return nil
}
