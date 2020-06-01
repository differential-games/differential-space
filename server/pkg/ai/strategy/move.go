package strategy

import (
	"github.com/differential-games/differential-space/pkg/game"
)

type MoveType int

const (
	Colonize MoveType = iota
	Attack
	Reinforce
)

// Move is a superset of game.Move which includes all necessary metadata
// for the AI To choose optimal moves.
type Move struct {
	game.Move
	// Distance is the distance between the two Planets.
	Distance float64
	// FromStrength is the strength of the Planet being moved from.
	FromStrength int
	// ToStrength is the strength of the Planet being moved to.
	ToStrength int

	MoveType

	// Score is how good the move is. Higher is better.
	Score float64
}

type MovePriority struct {
	colonize, attack, reinforce float64
}

func NewMovePriority(colonize, attack, reinforce float64) MovePriority {
	return MovePriority{
		colonize:  colonize,
		attack:    attack,
		reinforce: reinforce,
	}
}

func (s MovePriority) Initialize(game.Game) {}

func (s MovePriority) Analyze([]Move) {}

func (s MovePriority) Score(move Move) float64 {
	switch move.MoveType {
	case Colonize:
		return s.colonize
	case Attack:
		return s.attack
	default:
		return s.reinforce
	}
}
