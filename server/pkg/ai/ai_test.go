package ai_test

import (
	"encoding/json"
	"testing"

	"github.com/differential-games/differential-space/pkg/ai/strategy"

	"github.com/differential-games/differential-space/pkg/ai"
	"github.com/differential-games/differential-space/pkg/game"
)

var hardAI = ai.Hard(1.0)

func TestPickMoves(t *testing.T) {
	// The AIs should never pick incorrect moves.
	g, err := game.New(game.DefaultOptions)
	if err != nil {
		t.Fatal(err)
	}

	// Simulate 100 turns of moves. This should be sufficient to reasonably guarantee the AI
	// never generates invalid moves.
	moves := make([]strategy.Move, 1000)
	for i := 0; i < 100; i++ {
		g.NextTurn()
		playerId := g.PlayerTurn

		picked := hardAI.PickMoves(playerId, g.Planets, moves)
		err := ai.Turn(g, hardAI, moves)
		if err != nil {
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
	got := hardAI.PickMoves(2, p, moves)

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
		Strategies: []strategy.VectorBuilder{
			{
				Builder: strategy.Builder(strategy.MovePriority(-1, -1, 0)),
				Weight:  1.0,
			},
			{
				Builder: strategy.ReinforceFront,
				Weight:  1.0,
			},
		},
	}
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
		Strategies: []strategy.VectorBuilder{
			{
				Builder: strategy.Builder(strategy.MovePriority(-1, -1, 0)),
				Weight:  1.0,
			},
			{
				Builder: strategy.ReinforceFront,
				Weight:  1.0,
			},
		},
	}
	got := player.PickMoves(2, p, moves)

	if len(got) != 8 {
		t.Errorf("got %d moves, want 8", len(got))
	}
}
