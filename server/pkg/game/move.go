package game

import (
	"math"

	"github.com/pkg/errors"
)

const (
	MaxFleetSize       = 8
	MaxMoveDistance    = 5.0
	MaxMoveDistanceSq  = MaxMoveDistance * MaxMoveDistance
	InvMaxMoveDistance = 1.0 / MaxMoveDistance
)

type Move struct {
	// From is the id of the planet being moved from.
	From int
	// To is the id of the planet being moved to.
	To int
}

func (g *Game) move(move Move) {
	from := &g.Planets[move.From]
	to := &g.Planets[move.To]

	// The mover cannot make another move.
	from.Ready = false

	if to.Owner != from.Owner && to.Owner != 0 {
		// The destination is colonized by another Player, so Fight over the Planet.
		win := Fight(from, to)
		if !win {
			// Battle lost, nothing else to do.
			return
		}
	}

	// Take over the Planet if it is not already owned by the mover.
	if to.Owner != from.Owner {
		g.Scores[to.Owner]--
		g.Scores[from.Owner]++
		if g.Scores[from.Owner] > 2*len(g.Planets)/3 {
			g.Winner = from.Owner
		}

		to.Owner = from.Owner
		from.Strength--
	}
	to.Ready = false

	// Move all possible ships.
	sumStrength := from.Strength + to.Strength
	if sumStrength <= MaxFleetSize {
		// All ships move.
		to.Strength = sumStrength
		from.Strength = 0
	} else {
		// Only move ships so that reinforced Planet has at most 8.
		to.Strength = MaxFleetSize
		from.Strength = sumStrength - to.Strength
	}
}

// Move performs an move from one planet to another.
//
// Returns an error if the Move is invalid.
func (g *Game) Move(move Move) error {
	err := g.validate(move)
	if err != nil {
		return err
	}

	g.move(move)
	return nil
}

func (g *Game) AssessAttack(move Move) (string, bool) {
	err := g.validate(move)
	if err != nil {
		return err.Error(), false
	}

	if g.Planets[move.To].Owner == 0 {
		return "Colonize", true
	}

	if g.Planets[move.From].Owner == g.Planets[move.To].Owner {
		return "Reinforce", true
	}
	from := g.Planets[move.From]
	to := g.Planets[move.To]
	return Analyze(from.Strength, to.Strength, WinProbability(
		math.Sqrt(DistSq(from.X, from.Y, to.X, to.Y)),
	)), true
}

func (g *Game) validate(move Move) error {
	if move.From == move.To {
		return errors.New("Cannot move from and to the same planet.")
	}

	from := g.Planets[move.From]
	if from.Owner != g.PlayerTurn {
		return errors.Errorf("It is not Player %d's turn.", from.Owner)
	}
	if !from.Ready {
		return errors.Errorf("This planet has already moved this turn.")
	}
	if from.Strength == 0 {
		return errors.Errorf("This planet has no ships to move.")
	}

	to := g.Planets[move.To]
	dsq := DistSq(from.X, from.Y, to.X, to.Y)
	if dsq > MaxMoveDistanceSq {
		return errors.Errorf("Target planet is too far.")
	}
	return nil
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
	return math.Max(0, 1.0-d*0.1)
}
