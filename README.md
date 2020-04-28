# differential-space
Differential Space is a minimalist turn-based space strategy game I made for my video series.

## How do I play this?
The game has two binaries - one for the server, and the other for a client.
I haven't yet written code to combine these two binaries into one.

### Launch the Server
If you're running on Windows, you can
[download the v0.1.0 binary](https://drive.google.com/file/d/1YbToM27Yhorl9Fa0C1bP2sWDIkx2t688/view?usp=sharing)
instead of building it.

1. [Install Go](https://golang.org/doc/install). Remember to add GOPATH to your PATH.
1. Build and launch the server.

   ```
   cd server
   go install differential-space.go && differential-space
   ```

Alternatively, after `go install` you can execute the binary locally from GOPATH.

### Install and Launch the Client
1. Install `differential-space-client.exe`. You'll have the option to immediately launch
differential-space-client, and add it to your Start Menu.
1. Otherwise, you'll be able to find it in `C:\Program Files (x86)\Made in GameMaker Studio 2`.

## Configuring the Game
Configure the game with a JSON configuration file.

For examples, see sample [easy](server/easy.json) and [hard](server/hard.json) configurations.
The server prints out the configuration it is using on startup.

Option definitions:
- **difficulty** `[0.0, 1.0]`, is how intelligent the AI players are.
Higher is more difficult. 
AI opponents on harder difficulties receive no benefits whatsoever - they just play better.
I haven't won against an AI at `1.0` yet.
- **humanPlayers** is the IDs of the human players.
The server automatically acts for non-human Players.
- **minRadius** is the minimum radius of Planets in the system.
- **numPlayers** is the number of Players in the game.
The server supports as many Players as you want, but the client only supports up to 4.
From 1 to 4, the player colors are Green, Red, Blue, Purple.
- **numPlanets** is the number of Planets in the game.
- **radius** is the maximum distance a Planet maybe from the center of the map.

### Play
Play occurs in a sequence of turns where each Player acts independently.
At the start of the game, each player owns one planet, and the rest are unowned.

The rectangle turn button in the upper-right-hand corner indicates whose turn it is.
At the beginning of the game, the rectangle is grey to indicate that it is no one's turn.
Click the button to advance the turn.

The goal of the game is to take over the galaxy by eliminating your opponents.

On your turn, each planet under your control with at least one ship on it may make a move.
The number of white dots on each planet indicates how many ships are there.
A planet may have at most eight ships.
A black square on a planet indicates that it has already used its move this turn.
All ships present on the planet participate in the move.

If a planet does not use its move, next turn it will build one ship.

Moves are broken down as follows:
#### Colonization
Colonization is from an owned planet to an unowned planet.
Colonization consumes one ship.
#### Reinforcement
Reinforcement is from an owned planet to an owned planet by the same player.
Excess ships stay on the original planet.
#### Attack
An attack is from an owned planet to a planet owned by another player.
Battle takes place in rounds, up to one for each ship on the attacker.
If you mouse over enemy planets, the game displays the odds of doing damage each round.
To win, you must win a number of rounds equal to the number of enemy ships, plus one to take over the planet.
The further away the enemy planet, the lower your chance of doing damage each round.

##### Example 1
- An enemy planet contains 2 ships.
- The win rate is 80% per round.
- You are attacking with 3 ships.

Then to take over the planet you must win all 3 rounds at 80% each.
Thus, you have a 51.2% chance of taking over the enemy planet.

#### Example 2
- An enemy planet contains 5 ships.
- The win rate is 70% per round.
- You are attacking with 2 ships.

You can't win, but you can still do damage.
In this case, you have the following probabilities of doing damage:
- 9% - 0 damage
- 42% - 1 damage
- 49% - 2 damage

## Troubleshooting

### The map is empty
If you just see an empty map, the game isn't connecting to the server properly.
Ensure the server is running.
If it is, port 8080 may be in use.
The game doesn't yet support modifying what port the server/client communicate through.

## TODO
These are not in order.

- **Color-blind support**.
The color scheme should be deuteronamoly- and deuteranopia-friendly, and use shapes to indicate ownership as well as colors.
- **Battle prediction**. Show estimated damage dealt to enemy planets.
- **Window resolution**. Allow setting the game window resolution.
- **Linux and Mac installation**.
- **Rotation speed**. Allow configuring how quickly stars move.
- **More players**. Support up to 16 players.

## Non-Goals
This was a fun project to work on for my video series, but I'm not going to do much else on it.
If you'd like to contribute code to do these, feel free!

- Remote multiplayer
