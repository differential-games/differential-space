package game

type Options struct {
	PlanetOptions
	PlayerOptions

	RotationSpeed float64 `json:"rotationSpeed"`
}

type PlanetOptions struct {
	// NumPlanets is the number of Planets in the game.
	NumPlanets int `json:"numPlanets"`
	// Radius is the max radius from the center of the galaxy.
	Radius int `json:"radius"`
	// MinRadius is the min radius from the center of the galaxy.
	MinRadius int `json:"minRadius"`
}

type PlayerOptions struct {
	// NumPlayers is the number of Players in the game, excluding
	// the Environment.
	NumPlayers int `json:"numPlayers"`
}

// DefaultOptions are the options used if none are specified.
var DefaultOptions = Options{
	PlanetOptions: PlanetOptions{
		NumPlanets: 60,
		Radius:     10,
		MinRadius:  4,
	},
	PlayerOptions: PlayerOptions{
		NumPlayers: 4,
	},
	RotationSpeed: 2.0,
}
