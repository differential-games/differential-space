package ai

import (
	"math/rand"
	"sort"
)

// PickAttacks picks good attack moves and returns:
// 1) the set of picked attack moves, and
// 2) the set of reinforcement moves that are still possible after those attacks.
//
// moves is a list of Attack and Reinforcement Moves. It must not contain any colonization Moves.
//
// This algorithm ignores coordinated attacks from multiple Planets, i.e. if one Planet alone could
// not reasonably win against a target, it will choose not to even if four combined could win easily.
func PickAttacks(moves []Move, analyzers []MoveAnalyzer, filterBuilders []MoveFilterBuilder) ([]Move, []Move) {
	var attack, reinforce []Move
	for _, m := range moves {
		if m.MoveType == Attack {
			attack = append(attack, m)
		} else {
			reinforce = append(reinforce, m)
		}
	}

	var priorities []MovePriority
	for _, analyzer := range analyzers {
		priorities = append(priorities, analyzer(attack))
	}

	// Sort moves from least to greatest distance.
	// We get the best returns in attacking closer enemies.
	rand.Shuffle(len(attack), func(i, j int) {
		attack[i], attack[j] = attack[j], attack[i]
	})
	sort.Slice(attack, func(i, j int) bool {
		for _, priority := range priorities {
			p := priority(attack[i], attack[j])
			switch p {
			case Left: return true
			case Right: return false
			}
		}

		// Choose randomly.
		return rand.Float64() < 0.5
	})

	var filters []MoveFilter
	for _, builder := range filterBuilders {
		filters = append(filters, builder(attack))
	}

	var attacks []Move
	from := make(map[int]bool)
	// Keep track of the planets we could launch an attack from.
	// Reinforcing from it would leave the planet vulnerable.
	couldAttackFrom := make(map[int]bool)

	attackLoop:
	for _, m := range attack {
		// Mark that we could attack from this Planet, even though we may choose not to.
		couldAttackFrom[m.From] = true
		if from[m.From] {
			// We're already attacking from here.
			continue
		}

		filterLoop:
		for _, filter := range filters {
			switch filter(m) {
			case Keep:
				break filterLoop
			case Discard:
				continue attackLoop
			}
		}

		from[m.From] = true
		attacks = append(attacks, m)
	}

	var reinforcements []Move
	for _, m := range reinforce {
		if couldAttackFrom[m.From] {
			// We could attack from this planet, so don't consider reinforcement moves from it since
			// that will leave it vulnerable to attack.
			continue
		}
		reinforcements = append(reinforcements, m)
	}
	return attacks, reinforcements
}
