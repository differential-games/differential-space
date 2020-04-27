package game

import "testing"

func TestGeneratePlayers(t *testing.T) {
	players := GeneratePlayers(DefaultOptions.Players)

	if len(players) != DefaultOptions.Players.NumPlayers+1 {
		t.Errorf("got len(GeneratePlayers()) = %d, want %d",
			len(players),
			DefaultOptions.Players.NumPlayers,
		)
	}
}

func TestGenerateStarts(t *testing.T) {
	planets, err := GeneratePlanets(DefaultOptions.Planets)
	if err != nil {
		t.Fatal(err)
	}

	starts := GenerateStarts(planets, DefaultOptions.Planets.Radius, DefaultOptions.Players.NumPlayers)

	if len(starts) != DefaultOptions.Players.NumPlayers {
		t.Errorf("got len(starts) = %d, want %d",
			len(starts),
			DefaultOptions.Players.NumPlayers,
		)
	}
}
