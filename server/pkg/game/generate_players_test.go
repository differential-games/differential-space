package game

import "testing"

func TestGeneratePlayers(t *testing.T) {
	players := GeneratePlayers(DefaultOptions.PlayerOptions)

	if len(players) != DefaultOptions.PlayerOptions.NumPlayers+1 {
		t.Errorf("got len(GeneratePlayers()) = %d, want %d",
			len(players),
			DefaultOptions.PlayerOptions.NumPlayers,
		)
	}
}

func TestGenerateStarts(t *testing.T) {
	planets, err := GeneratePlanets(DefaultOptions.PlanetOptions)
	if err != nil {
		t.Fatal(err)
	}

	starts := GenerateStarts(planets, DefaultOptions.PlanetOptions.Radius, DefaultOptions.PlayerOptions.NumPlayers)

	if len(starts) != DefaultOptions.PlayerOptions.NumPlayers {
		t.Errorf("got len(starts) = %d, want %d",
			len(starts),
			DefaultOptions.PlayerOptions.NumPlayers,
		)
	}
}
