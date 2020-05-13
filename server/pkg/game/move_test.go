package game

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"testing"
)

func TestWinProbability(t *testing.T) {
	tcs := []struct {
		name string
		d    float64
		want float64
	}{
		{
			name: "100% at 0 distance",
			d:    0.0,
			want: 1.0,
		},
		{
			name: "50% at 5 distance",
			d:    5.0,
			want: 0.5,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			p := WinProbability(tc.d)

			diff := cmp.Diff(tc.want, p, cmpopts.EquateApprox(0, 0.001))
			if diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestGame_ValidAttack(t *testing.T) {
	g := Game{
		PlayerTurn: 1,
		Planets: []Planet{
			{Owner: 1, Ready: true, Strength: 1},
			{Owner: 2},
		},
		Players: GeneratePlayers(PlayerOptions{NumPlayers: 2}),
	}

	err := g.Move(Move{
		From: 0,
		To:   1,
	})

	if err != nil {
		t.Error(err)
	}

	if g.Planets[1].Strength != 0 {
		t.Errorf("got Strength = %d, want 0", g.Planets[1].Strength)
	}
}

func TestGame_AttackSelf(t *testing.T) {
	g := Game{
		PlayerTurn: 1,
		Planets: []Planet{
			{Owner: 1, Ready: true, Strength: 1},
			{Owner: 2},
		},
		Players: GeneratePlayers(PlayerOptions{NumPlayers: 2}),
	}

	err := g.Move(Move{
		From: 0,
		To:   0,
	})

	if err == nil {
		t.Error("got Move() = nil, want err")
	}
}

func TestGame_AttackError(t *testing.T) {
	tcs := []struct {
		name     string
		turn     int
		attacker Planet
		defender int
	}{
		{
			name: "player cannot attack out of turn",
			turn: 2,
			attacker: Planet{
				Owner:    1,
				Ready:    true,
				Strength: 1,
			},
			defender: 2,
		},
		{
			name: "planet cannot attack if not ready",
			turn: 1,
			attacker: Planet{
				Owner:    1,
				Ready:    false,
				Strength: 1,
			},
			defender: 2,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			g := Game{
				PlayerTurn: tc.turn,
				Planets: []Planet{
					tc.attacker,
					{Owner: tc.defender},
				},
				Players: []Player{
					{},
					{},
					{},
				},
			}
			err := g.Move(Move{
				From: 0,
				To:   1,
			})

			if err == nil {
				t.Error("got Move() = nil, want err")
			}
		})
	}
}

func TestMove_Colonize(t *testing.T) {
	g := Game{
		PlayerTurn: 1,
		Planets: []Planet{
			{
				Owner: 1,
				Strength: 1,
				Ready: true,
			},
			{
				Owner: 0,
			},
		},
	}

	err := g.Move(Move{
		From: 0,
		To:   1,
	})

	if err != nil {
		t.Fatal(err)
	}

	from := g.Planets[0]
	if diff := cmp.Diff(Planet{Owner: 1, Strength: 0, Ready: false}, from); diff != "" {
		t.Error(diff)
	}
	to := g.Planets[1]
	if diff := cmp.Diff(Planet{Owner: 1, Strength: 0, Ready: false}, to); diff != "" {
		t.Error(diff)
	}
}
