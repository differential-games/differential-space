package game

type Options struct {
	Planets PlanetOptions
	Players PlayerOptions
	// Difficulty is how optimal the computer plays.
	// 1.0 is hard.
	// 0.0 the opponent does nothing.
	Difficulty float64
}

type PlanetOptions struct {
	// NumPlanets is the number of Planets in the game.
	NumPlanets int
	// Radius is the max radius from the center of the galaxy.
	Radius int
	// MinRadius is the min radius from the center of the galaxy.
	MinRadius int
}

type PlayerOptions struct {
	// NumPlayers is the number of Players in the game, excluding
	// the Environment.
	NumPlayers int
}

var DefaultOptions = Options{
	Difficulty: 0.5,
	Planets: PlanetOptions{
		NumPlanets: 40,
		Radius:     10,
		MinRadius:  5,
	},
	Players: PlayerOptions{
		NumPlayers: 4,
	},
}
