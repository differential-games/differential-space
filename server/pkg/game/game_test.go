package game

import "testing"

func TestNew(t *testing.T) {
	_, err := New(DefaultOptions)

	if err != nil {
		t.Error(err)
	}
}

func TestNew_GenerateStartsFailure(t *testing.T) {
	_, err := New(func() Options {
		options := DefaultOptions
		options.PlanetOptions.NumPlanets = 5
		options.PlayerOptions.NumPlayers = 10
		return options
	}())

	if err == nil {
		t.Error("got New() = nil, want err")
	}
}

func TestNew_GeneratePlanetFailure(t *testing.T) {
	_, err := New(func() Options {
		options := DefaultOptions
		options.PlanetOptions.Radius = 0
		return options
	}())

	if err == nil {
		t.Error("got New() = nil, want err")
	}
}

func TestGame_NextTurnMovesPlanets(t *testing.T) {
	g, err := New(DefaultOptions)
	if err != nil {
		t.Fatal(err)
	}

	before := g.Planets[0]
	g.NextTurn()
	after := g.Planets[0]

	if before.Angle == after.Angle {
		t.Error("planet did not change Angle")
	}
	if before.X == after.X && before.Y == after.Y {
		t.Error("planet did not move")
	}
}

func TestGame_NextTurnReadiesPlayerPlanets(t *testing.T) {
	g, err := New(DefaultOptions)
	if err != nil {
		t.Fatal(err)
	}

	g.Planets[0].Owner = 1
	g.Planets[0].Ready = false

	for i := 0; i < len(g.Players)+1; i++ {
		g.NextTurn()
	}

	if !g.Planets[0].Ready {
		t.Error("got Ready = false, want true")
	}
}

func TestGame_NextTurnBuildsShips(t *testing.T) {
	g, err := New(DefaultOptions)
	if err != nil {
		t.Fatal(err)
	}

	g.Planets[0].Owner = 1
	g.Planets[0].Ready = true
	g.Planets[0].Strength = 0

	g.NextTurn()

	if !g.Planets[0].Ready {
		t.Error("got Ready = false, want true")
	}
	if g.Planets[0].Strength != 1 {
		t.Errorf("got Strength = %d, want 1", g.Planets[0].Strength)
	}
}

func TestGame_NextTurnDoesNotReadyEnvironment(t *testing.T) {
	g, err := New(DefaultOptions)
	if err != nil {
		t.Fatal(err)
	}
	g.PlayerTurn = 0
	g.Planets[0].Owner = 0
	g.Planets[0].Ready = false

	for i := 0; i < len(g.Players)+1; i++ {
		g.NextTurn()
	}

	if g.Planets[0].Ready {
		t.Error("got Ready = true, want false")
	}
}
