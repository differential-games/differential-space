package game

import (
	"math"

	"github.com/pkg/errors"
)

type Game struct {
	PlayerTurn int

	Planets []Planet
	Players []Player

	Scores []int
	Winner int

	// RotationSpeed is how quickly the galaxy rotates.
	//
	// The rotation, in radians, of each planet between consecutive turns for the same player is equal to
	// (RotationSpeed / planet radius).
	RotationSpeed float64
}

// New creates a new Game according to the passed Opts.
func New(options Options) (*Game, error) {
	if options.PlayerOptions.NumPlayers > options.PlanetOptions.NumPlanets {
		return nil, errors.Errorf("NumPlayers %d greater than NumPlanets %d",
			options.PlayerOptions.NumPlayers,
			options.PlanetOptions.NumPlanets,
		)
	}

	// Generate a galaxy.
	planets, err := GeneratePlanets(options.PlanetOptions)
	if err != nil {
		return nil, err
	}

	// Generate Players and their starting positions.
	players := GeneratePlayers(options.PlayerOptions)
	GenerateStarts(planets, options.PlanetOptions.Radius, []PlayerStart{
		{
			Owner: 1,
			Starts: []Start{
				{Strength: 3},
			},
		},
		{
			Owner: 2,
			Starts: []Start{
				{Strength: 1},
				{Strength: 0},
			},
		},
	})

	scores := make([]int, 1+options.NumPlayers)
	scores[0] = options.NumPlanets - 3
	scores[1] = 1
	scores[2] = 2

	result := Game{
		PlayerTurn:    0,
		Planets:       planets,
		Players:       players,
		Scores:        scores,
		RotationSpeed: options.RotationSpeed,
	}
	return &result, nil
}

// NextTurn ends the current Player's turn and moves to the next one.
//
// Returns the id of the next player, and whether they have any planets.
func (g *Game) NextTurn() (int, bool) {
	g.PlayerTurn++
	if g.PlayerTurn == len(g.Players) {
		// The environment Player, 0, doesn't get a turn.
		g.PlayerTurn = 1
	}

	score := 0
	for i, p := range g.Planets {
		if p.Owner == g.PlayerTurn {
			score++
			if g.Planets[i].Ready {
				// The Planet didn't do anything last turn, so build a ship.
				g.Planets[i].Strength++
				if g.Planets[i].Strength > 8 {
					// A Planet can't have more than 8 ships.
					g.Planets[i].Strength = 8
				}
			} else {
				// The Planet is not ready, so ready it.
				g.Planets[i].Ready = true
			}
		}
	}

	// Advance time in the galaxy.
	g.Rotate(g.RotationSpeed / float64(len(g.Players)))

	return g.PlayerTurn, score == 0
}

func (g *Game) Rotate(t float64) {
	for i := range g.Planets {
		g.Planets[i].Angle += t * g.Planets[i].InvRadius
		g.Planets[i].X = g.Planets[i].Radius * math.Sin(g.Planets[i].Angle)
		g.Planets[i].Y = g.Planets[i].Radius * math.Cos(g.Planets[i].Angle)
	}
}
