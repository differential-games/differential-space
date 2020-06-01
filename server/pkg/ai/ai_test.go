package ai_test

import (
	"encoding/json"
	"testing"

	"github.com/differential-games/differential-space/pkg/ai/strategy"

	"github.com/differential-games/differential-space/pkg/ai"
	"github.com/differential-games/differential-space/pkg/game"
)

func TestPickMoves(t *testing.T) {
	// The AIs should never pick incorrect moves.
	g, err := game.New(game.DefaultOptions)
	if err != nil {
		t.Fatal(err)
	}

	// Simulate 100 turns of moves. This should be sufficient to reasonably guarantee the AI
	// never generates invalid moves.
	moves := make([]strategy.Move, 1000)
	player := ai.Hard(game.DefaultOptions.NumPlanets, 1.0)
	player.Initialize(*g)
	for i := 0; i < 100; i++ {
		g.NextTurn()
		playerId := g.PlayerTurn

		err := ai.Turn(g, player, moves)
		if err != nil {
			picked := player.PickMoves(playerId, g.Planets, moves)
			t.Error("turn", i)
			jsn, _ := json.MarshalIndent(g.Planets, "", "  ")
			t.Error(string(jsn))
			jsn, _ = json.MarshalIndent(picked, "", "  ")
			t.Error(string(jsn))
			t.Fatal(err)
		}
	}
}

func TestAttack(t *testing.T) {
	p := []game.Planet{
		// Target Planet
		{
			X:     0,
			Owner: 1,
		},
		// Attacker.
		{
			X:        4,
			Owner:    2,
			Strength: 1,
			Ready:    true,
		},
	}

	moves := make([]strategy.Move, 5)
	player := ai.Hard(2, 1.0)
	player.Initialize(game.Game{Planets: p})
	got := player.PickMoves(2, p, moves)

	if len(got) != 1 {
		t.Errorf("got %d moves, want 1", len(got))
	}
}

func TestReinforce_1(t *testing.T) {
	p := []game.Planet{
		// Target Planet
		{
			X:     0,
			Owner: 1,
		},
		// Attacker.
		{
			X:        4,
			Owner:    2,
			Strength: 1,
			Ready:    true,
		},
		// Reinforcer
		{
			X:        8,
			Owner:    2,
			Strength: 1,
			Ready:    true,
		},
	}

	moves := make([]strategy.Move, 100)
	player := ai.AI{
		Difficulty: 1.0,
		Strategies: []strategy.Vector{
			{
				Strategy: strategy.NewMovePriority(-1, -1, 0),
				Weight:   1.0,
			},
			{
				Strategy: strategy.NewReinforceFront(100),
				Weight:   1.0,
			},
		},
	}
	player.Initialize(game.Game{Planets: p})
	got := player.PickMoves(2, p, moves)

	if len(got) != 1 {
		t.Errorf("got %d moves, want 1", len(got))
		jsn, _ := json.MarshalIndent(got, "", "  ")
		t.Log(string(jsn))
	}
}

func TestReinforce_8(t *testing.T) {
	p := []game.Planet{
		// Target Planet
		{
			X:     0,
			Owner: 1,
		},
		// Attacker.
		{
			X:        4,
			Owner:    2,
			Strength: 1,
			Ready:    true,
		},
		// Reinforcers
		{
			X:        8,
			Owner:    2,
			Strength: 1,
			Ready:    true,
		},
		{
			X:        12,
			Owner:    2,
			Strength: 2,
			Ready:    true,
		},
		{
			X:        16,
			Owner:    2,
			Strength: 3,
			Ready:    true,
		},
		{
			X:        20,
			Owner:    2,
			Strength: 4,
			Ready:    true,
		},
		{
			X:        24,
			Owner:    2,
			Strength: 5,
			Ready:    true,
		},
		{
			X:        28,
			Owner:    2,
			Strength: 6,
			Ready:    true,
		},
		{
			X:        32,
			Owner:    2,
			Strength: 7,
			Ready:    true,
		},
		{
			X:        36,
			Owner:    2,
			Strength: 8,
			Ready:    true,
		},
	}

	moves := make([]strategy.Move, 100)
	player := ai.AI{
		Difficulty: 1.0,
		Strategies: []strategy.Vector{
			{
				Strategy: strategy.NewMovePriority(-1, -1, 0),
				Weight:   1.0,
			},
			{
				Strategy: strategy.NewReinforceFront(100),
				Weight:   1.0,
			},
		},
	}
	player.Initialize(game.Game{Planets: p})
	got := player.PickMoves(2, p, moves)

	if len(got) != 8 {
		t.Errorf("got %d moves, want 8", len(got))
	}
}
