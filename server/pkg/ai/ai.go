package ai

import (
	"github.com/differential-games/differential-space/pkg/game"
	"math/rand"
	"sort"
)

type Interface interface {
	PickMoves(player int, planets []game.Planet) []game.Move
}

type AI struct {
	Difficulty float64

	Colonize []MoveAnalyzer
	Attack []MoveAnalyzer
	AttackFilter []MoveFilterBuilder
}

// PickMoves executes the AI algorithm for a Player
//
// player is the id of the Player to act.
// planets is the full set of Planets in the Game.
// difficulty is how hard the AI opponent is to defeat.
//
// difficulty:
// 1.0 - very hard
// 0.5 - easy
// 0.0 - brain dead (the AI does nothing)
func (ai *AI) PickMoves(player int, planets []game.Planet) []game.Move {
	possible := GenerateMoves(player, planets)

	colonizes, others := PickColonization(possible, ai.Colonize)
	attacks, possibleReinforcements := PickAttacks(others, ai.Attack, ai.AttackFilter)
	reinforcements := PickReinforcements(colonizes, attacks, possibleReinforcements)

	chosen := colonizes
	chosen = append(chosen, attacks...)
	chosen = append(chosen, reinforcements...)

	var result []game.Move
	// Randomly filter out moves per the difficulty setting.
	// In general, make fewer moves the easier the difficulty is, chosen at random.
	// This results in less-optimal play but still interesting opponents.
	// At 1.0, this just includes all of the optimal moves.
	for _, m := range chosen {
		r := rand.Float64()
		switch m.MoveType {
		case Colonize:
			// Increase colonization rate on easier difficulties, or the AIs expand too slowly to be interesting opponents.
			if r*r < ai.Difficulty {
				result = append(result, m.Move)
			}
		default:
			// Limit attacking and reinforcing per the difficulty setting.
			// This makes easier AIs fight and reinforce themselves much less efficiently.
			if r < ai.Difficulty {
				result = append(result, m.Move)
			}
		}
	}
	return result
}

// GenerateMoves generates all possible Moves for a Player.
func GenerateMoves(player int, planets []game.Planet) []Move {
	var result []Move

	for i, from := range planets {
		if player != from.Owner {
			// Can't start move from someone else's Planet.
			continue
		}
		if !from.Ready || from.Strength == 0 {
			// Can't start move from a Planet that isn't Ready or has no troops.
			continue
		}
		for j, to := range planets {
			if i == j {
				// Can't move from a Planet to itself.
				continue
			}
			d := game.Dist(from, to)
			if d > 5.0 {
				// Moves over 5 units are invalid.
				continue
			}
			// Construct the potential Move and record the relevant metadata.
			m := Move{
				Move: game.Move{
					From: i,
					To:   j,
				},
				Distance:     d,
				FromStrength: from.Strength,
				ToStrength:   to.Strength,
			}
			switch to.Owner {
			case 0:
				m.MoveType = Colonize
			case player:
				m.MoveType = Reinforce
			default:
				m.MoveType = Attack
			}
			result = append(result, m)
		}
	}

	return result
}

// MoveType classifies a Move
type MoveType string

const (
	Colonize  MoveType = "Colonize"
	Reinforce MoveType = "Reinforce"
	Attack    MoveType = "Attack"
)

// PickReinforcements chooses a reasonable reinforcement pattern based on which Planets are being
// moved from this turn.
//
// Specifically, it prioritizes:
// 1) Attackers
// 2) Colonizer
// 3) Reinforcers
func PickReinforcements(colonizes, attacks, reinforcements []Move) []Move {
	attackFrom := make(map[int]bool)
	for _, m := range attacks {
		attackFrom[m.From] = true
	}
	colonizeFrom := make(map[int]bool)
	for _, m := range colonizes {
		colonizeFrom[m.From] = true
	}

	sort.Slice(reinforcements, func(i, j int) bool {
		// Prefer reinforcing attackers over other actions.
		if attackFrom[reinforcements[i].To] != attackFrom[reinforcements[j].To] {
			return attackFrom[reinforcements[i].To]
		}
		// Prefer reinforcing colonizers over reinforcements.
		if colonizeFrom[reinforcements[i].To] != colonizeFrom[reinforcements[j].To] {
			return colonizeFrom[reinforcements[i].To]
		}
		// Prefer reinforcing from as far away as possible in as few moves as possible.
		// This is a greedy algorithm for simplicity/speed of computation.
		return reinforcements[i].Distance > reinforcements[j].Distance
	})

	// Remember, we're considering them from furthest to nearest, but targets who attacked are prioritized.
	var result []Move
	// Keep track of how much we're already reinforcing the Planet. Don't try to send more than 8
	// reinforcements.
	gotReinforcements := make(map[int]int)
	reinforceFrom := make(map[int]bool)
	// Pick reinforcing attackers and colonizers.
	for _, m := range reinforcements {
		if reinforceFrom[m.From] {
			// We already got a good reinforcement from here.
			continue
		}
		if !attackFrom[m.To] && !colonizeFrom[m.To] {
			// The target Planet isn't being attacked or colonized from, so ignore it for now.
			continue
		}
		if gotReinforcements[m.To] >= 8 {
			// We can't send more than 8 reinforcements to the same Planet in a turn.
			continue
		}
		// We're attacking from the Planet this reinforcement is to, so do so.
		reinforceFrom[m.From] = true
		gotReinforcements[m.To] += m.FromStrength
		result = append(result, m)
	}

	// Lastly, reinforce planets that sent reinforcements.
	// Each iteration is building a tree where the root node is an attacker/colonizer.
	// We iterate multiple times so a well-guarded base of producing Planets can quickly transport
	// troops to the front lines. The further the reinforcement is from the front, the better it
	// is to wait before sending it, as otherwise we lose a turn of production.
	for minReinforcement := 2; minReinforcement <= 8; minReinforcement++ {
		for _, m := range reinforcements {
			if m.FromStrength < minReinforcement {
				// Only send reinforcements if the number of steps to the front is less
				// than the size of the reinforcements.
				continue
			}
			if reinforceFrom[m.From] {
				// We're already reinforcing from here.
				continue
			}
			if !reinforceFrom[m.To] {
				// The target of this move hasn't reinforced another Planet, so ignore it for now.
				continue
			}
			if gotReinforcements[m.To] >= 8 {
				// We can't send more than 8 reinforcements to the same Planet in a turn.
				continue
			}
			// We're reinforcing from the Planet this reinforcement is to, so do so.
			reinforceFrom[m.From] = true
			gotReinforcements[m.To] += m.FromStrength
			result = append(result, m)
		}
	}

	return result
}
