package ai

import (
	"github.com/differential-games/differential-space/pkg/ai/strategy"
	"github.com/differential-games/differential-space/pkg/game"
	"math/rand"
	"sort"
)

type Interface interface {
	PickMoves(player int, planets []game.Planet) []game.Move
}

type AI struct {
	Difficulty float64

	Colonize []strategy.VectorBuilder
	Attack []strategy.VectorBuilder
}

func deduplicate(moves []strategy.Move) ([]strategy.Move, map[int]bool) {
	var result []strategy.Move
	used := make(map[int]bool)

	for _, m := range moves {
		if used[m.From] {
			continue
		}
		result = append(result, m)
		used[m.From] = true
	}
	return result, used
}

func filterUsed(used map[int]bool, moves []strategy.Move) []strategy.Move {
	var result []strategy.Move
	for _, m := range moves {
		if !used[m.From] {
			result = append(result, m)
		}
	}
	return result
}

func filterRandom(moves []strategy.Move, p float64) []strategy.Move {
	if p >= 1.0 {
		return moves
	}
	var result []strategy.Move
	for _, m := range moves {
		if rand.Float64() < p {
			result = append(result, m)
		}
	}
	return result
}

// PickMoves executes the AI algorithm for a Player
func (ai *AI) PickMoves(player int, planets []game.Planet) []game.Move {
	colonizes, attacks, reinforces := GenerateMoves(player, planets)

	colonizes = Pick(colonizes, ai.Colonize)
	colonizes = filterRandom(colonizes, ai.Difficulty)
	colonizes, usedColonizes := deduplicate(colonizes)

	attacks = filterUsed(usedColonizes, attacks)
	attacks = Pick(attacks, ai.Attack)
	attacks = filterRandom(attacks, ai.Difficulty)
	attacks, usedAttacks := deduplicate(attacks)

	reinforces = filterUsed(usedColonizes, reinforces)
	reinforces = filterUsed(usedAttacks, reinforces)
	reinforces = filterRandom(reinforces, ai.Difficulty)
	reinforces = PickReinforcements(colonizes, attacks, reinforces)

	var result []game.Move
	for _, c := range colonizes {
		result = append(result, c.Move)
	}
	for _, a := range attacks {
		result = append(result, a.Move)
	}
	for _, r := range reinforces {
		result = append(result, r.Move)
	}
	return result
}

// GenerateMoves generates all possible Moves for a Player.
func GenerateMoves(player int, planets []game.Planet) ([]strategy.Move, []strategy.Move, []strategy.Move) {
	var colonizes []strategy.Move
	var attacks []strategy.Move
	var reinforces []strategy.Move
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
			m := strategy.Move{
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
				colonizes = append(colonizes, m)
			case player:
				reinforces = append(reinforces, m)
			default:
				attacks = append(attacks, m)
			}
		}
	}

	return colonizes, attacks, reinforces
}

// PickReinforcements chooses a reasonable reinforcement pattern based on which Planets are being
// moved from this turn.
//
// Specifically, it prioritizes:
// 1) Attackers
// 2) Colonizer
// 3) Reinforcers
func PickReinforcements(colonizes, attacks, reinforcements []strategy.Move) []strategy.Move {
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
	var result []strategy.Move
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
