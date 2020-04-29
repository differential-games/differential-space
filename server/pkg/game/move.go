package game

import (
	"github.com/pkg/errors"
	"math"
	"math/rand"
)

const (
	ColonizeRadius = 5.0
	MaxColonizeRadius = 10.0
)

type Move struct {
	// From is the id of the planet being moved from.
	From int
	// To is the id of the planet being moved to.
	To   int
}

// Move performs an move from one planet to another.
//
// Returns an error if the Move is invalid.
func (g *Game) Move(move Move) error {
	p, err := g.AttackProbability(move)
	if err != nil {
		return err
	}

	// Colonization and reinforcement automatically succeeds.
	win := false
	if g.Planets[move.To].Owner == 0 {
		win = true
		g.Planets[move.From].Strength--
	} else if g.Planets[move.To].Owner == g.Planets[move.From].Owner {
		win = true
	}

	// The planet's forces have taken an action, so they cannot move again.
	g.Planets[move.From].Ready = false

	if !win {
		for {
			if g.Planets[move.From].Strength == 0 || g.Planets[move.To].Strength < 0 {
				win = g.Planets[move.To].Strength == -1
				break;
			}
			g.Planets[move.From].Strength--
			if rand.Float64() < p {
				g.Planets[move.To].Strength--
			}
		}
	}

	if win {
		// The planet has been taken over by the mover.
		g.Planets[move.To].Owner = g.Planets[move.From].Owner
		// The just-conquered Planet is not ready.
		g.Planets[move.To].Ready = false

		if g.Planets[move.To].Strength < 0 {
			g.Planets[move.To].Strength = 0
		}
		sumStrength := g.Planets[move.To].Strength + g.Planets[move.From].Strength
		// Up to 8 ships move.
		if sumStrength >= 8 {
			g.Planets[move.To].Strength = 8
			g.Planets[move.From].Strength = sumStrength - 8
		} else {
			g.Planets[move.To].Strength = sumStrength
			g.Planets[move.From].Strength = 0
		}
	}

	return nil
}

func (g *Game) AssessAttack(move Move) (string, bool) {
	p, err := g.AttackProbability(move)
	if err != nil {
		return err.Error(), false
	}

	if g.Planets[move.To].Owner == 0 {
		return "Colonize", true
	}

	if g.Planets[move.From].Owner == g.Planets[move.To].Owner {
		return "Reinforce", true
	}
	return Analyze(g.Planets[move.From].Strength, g.Planets[move.To].Strength, p), true
}

func (g *Game) AttackProbability(move Move) (float64, error) {
	if move.From == move.To {
		return 0.0, errors.New("Cannot move from and to the same planet.")
	}

	from := g.Planets[move.From]
	if from.Owner != g.PlayerTurn {
		return 0.0, errors.Errorf("It is not Player %d's turn.", from.Owner)
	}
	if !from.Ready {
		return 0.0, errors.Errorf("This planet has already moved this turn.")
	}
	if from.Strength == 0 {
		return 0.0, errors.Errorf("This planet has no ships to move.")
	}

	to := g.Planets[move.To]
	d := Dist(from, to)
	if d > 5 {
		return 0.0, errors.Errorf("Target planet is too far.")
	}
	if to.Owner == 0 || to.Owner == from.Owner {
		return 1.0, nil
	}

	return WinProbability(d), nil
}

// WinProbability returns the probability of an attacker inflicting damage in a round of battle.
//
// from is the id of the Planet launching the attack.
// to is the id of the Planet being attacked.
//
// attackerTech is the tech level of the attacker.
// defenderTech is the tech level of the defender.
func WinProbability(d float64) float64 {
	// Decrease by 10% per unit.
	return math.Max(0, 1.0 - d / 10)
}
