package ai

import (
	"encoding/json"
	"github.com/differential-games/differential-space/pkg/game"
	"testing"
)

func TestPickMoves(t *testing.T) {
	// The AIs should never pick incorrect moves.
	g, err := game.New(game.DefaultOptions)
	if err != nil {
		t.Fatal(err)
	}

	// Simulate 100 turns of moves. This should be sufficient to reasonably guarantee the AI
	// never generates invalid moves.
	for i := 0; i < 100; i++ {
		g.NextTurn()
		player := g.PlayerTurn
		moves := PickMoves(player, g.Planets, 1.0)

		for _, m := range moves {
			err = g.Move(m)
			if err != nil {
				jsn, _ := json.MarshalIndent(g.Planets[m.From], "", "  ")
				t.Error(string(jsn))
				jsn, _ = json.MarshalIndent(g.Planets[m.To], "", "  ")
				t.Error(string(jsn))
				jsn, _ = json.MarshalIndent(m, "", "  ")
				t.Error(string(jsn))
				jsn, _ = json.MarshalIndent(moves, "", "  ")
				t.Error(string(jsn))
				t.Fatal(err)
			}
		}
	}
}

func TestAttack(t *testing.T) {
	p := []game.Planet{
		// Target Planet
		{
			X: 0,
			Owner: 1,
		},
		// Attacker.
		{
			X: 4,
			Owner: 2,
			Strength: 1,
			Ready: true,
		},
	}

	got := PickMoves(2, p, 1.0)

	if len(got) != 1 {
		t.Errorf("got %d moves, want 1", len(got))
	}
}

func TestReinforce_1(t *testing.T) {
	p := []game.Planet{
		// Target Planet
		{
			X: 0,
			Owner: 1,
		},
		// Attacker.
		{
			X: 4,
			Owner: 2,
			Strength: 1,
			Ready: true,
		},
		// Reinforcer
		{
			X: 8,
			Owner: 2,
			Strength: 1,
			Ready: true,
		},
	}

	got := PickMoves(2, p, 1.0)

	if len(got) != 2 {
		t.Errorf("got %d moves, want 2", len(got))
	}
}

func TestReinforce_8(t *testing.T) {
	p := []game.Planet{
		// Target Planet
		{
			X: 0,
			Owner: 1,
		},
		// Attacker.
		{
			X: 4,
			Owner: 2,
			Strength: 1,
			Ready: true,
		},
		// Reinforcers
		{
			X: 8,
			Owner: 2,
			Strength: 1,
			Ready: true,
		},
		{
			X: 12,
			Owner: 2,
			Strength: 2,
			Ready: true,
		},
		{
			X: 16,
			Owner: 2,
			Strength: 3,
			Ready: true,
		},
		{
			X: 20,
			Owner: 2,
			Strength: 4,
			Ready: true,
		},
		{
			X: 24,
			Owner: 2,
			Strength: 5,
			Ready: true,
		},
		{
			X: 28,
			Owner: 2,
			Strength: 6,
			Ready: true,
		},
		{
			X: 32,
			Owner: 2,
			Strength: 7,
			Ready: true,
		},
		{
			X: 36,
			Owner: 2,
			Strength: 8,
			Ready: true,
		},
	}

	got := PickMoves(2, p, 1.0)

	if len(got) != 9 {
		t.Errorf("got %d moves, want 8", len(got))
	}
}
