# Differential Space
Differential Space is a minimalist turn-based space strategy game I made for [my video series](https://www.youtube.com/channel/UC29CSoWw7ASMCTcjzc-7k0w).
The purpose is to illustrate how strategy in terrain that moves changes how players behave.

## How do I play this?
The game has two binaries - one for the server, and the other for a client.
The server actually runs the game, and the client is the UI players interact with.

I haven't yet written code to combine these two binaries into one.

### Server
You've got two options here.
The game will ask to communicate through your network.
Feel free to forbid this since it doesn't keep the client and server from communicating on your computer.
I don't know how to disable this message.

#### 1. Download the Windows Server Executable
If you're okay running arbitrary code compiled by a stranger on the internet, [download it here](bin/differential-space-server.exe).

#### 2. Compile the Server
1. [Install Go](https://golang.org/doc/install). Remember to add GOPATH to your PATH.
1. Build the server.

   ```
   cd server
   go build differential-space.go
   ```

### Client
1. Download and Run [`differential-space-client.exe`](bin/differential-space-client.exe).
This is a Game Maker Studio installer.
You'll have the option to immediately launch `differential-space-client`, and add it to your Start Menu.
1. Otherwise, you'll be able to find it in `C:\Program Files (x86)\Made in GameMaker Studio 2`.

## Configuring the Game
Configure the game with a JSON configuration file.

For examples, see sample [easy](server/easy.json) and [hard](server/hard.json) configurations.
The server prints out the configuration it is using on startup.

Option definitions:
- **difficulty** is how intelligent the AI players are on a scale from 0.0 to 1.0.
Higher is more difficult. 
AI opponents on harder difficulties receive no benefits whatsoever - they just play better.
I haven't won against an AI at 1.0 yet.
- **humanPlayers** is the IDs of the human players.
The server automatically acts for non-human players.
- **minRadius** is the minimum radius of planets in the system.
- **numPlayers** is the number of players in the game.
The server supports as many players as you want, but the client only supports up to 4.
From 1 to 4, the player colors are Green, Red, Blue, Purple.
- **numPlanets** is the number of planets in the game.
- **radius** is the maximum distance a planet maybe from the center of the map.

### Play
**Controls**
- Left mouse button: Select planet, next turn button.
- Right mouse button: Execute move to planet.
- Middle mouse button: Pan map.
- Scroll wheel: Zoom map in/out.

Play occurs in a sequence of turns where each player acts independently.
Each circle on the screen represents a single planet.
At the start of the game, each player owns one planet, and the rest are unowned.

The rectangle turn button in the upper-left-hand corner indicates whose turn it is.
At the beginning of the game, the rectangle is grey to indicate that it is no one's turn.
Click the button to advance the turn.

The goal of the game is to take over the galaxy by eliminating your opponents.

On your turn, each planet under your control with at least one ship on it may make a move.
Planets you can issue moves from have a white glowing dot at their center.
To move, left-click on one of your planets, and right-click on a target.
The game will let you know if a move is invalid with a message at the top of the screen.

The number of white dots on each planet indicates how many ships are there.
A planet may have at most eight ships.
All ships present on the planet participate in the move.

If a planet does not use its move, next turn it will build one ship.

Moves fall into one of three categories:
#### Colonization
Colonization is from an owned planet to an unowned planet.
Colonization consumes one ship.
#### Reinforcement
Reinforcement is from an owned planet to an owned planet by the same player.
Excess ships stay on the original planet.
#### Attack
An attack is from an owned planet to a planet owned by another player.
Battle takes place in rounds, up to one for each ship on the attacker.
If you mouse over enemy planets, the game displays either:
- the probability you'll take over the planet, or (if the probability is 0%)
- the expected damage you'll do.

The exact damage dealt is random, so you may deal more or less damage than the estimate.

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

The UI will show that you are expected to do 1 damage, but you could do 0 or 2.

## Troubleshooting

### The map is empty
If you just see an empty map, the game isn't connecting to the server properly.
Ensure the server is running.
If it is, port 8080 may be in use (the server should show an error on startup if this is the case).
The game doesn't yet support modifying what port the server/client communicate through.

## TODO
These are not in order.

- **Window resolution**.
Allow setting the game window resolution.
- **Linux and Mac installation**.
I don't have access to Mac/Linux machines in quarantine, so I can't build for systems other than Windows. :/
- **Rotation speed**.
Allow configuring how quickly stars move.
- **Changing the port**
To help resolve 'map is empty' problem for players who want to use a different port than 8080.
- **Planet status box**
Not only will this help those with vision impairments, but it can help the game communicate its rules better.
- **Planet names**
This makes it easier to remember the lay of the map (especially as it is moving), and talk about it.
- **Magnified mode**
Enlarged sprites for those with vision impairments.
Possibly toggleable on the keyboard?
Zoom in on moused-over planet?

## Non-Goals
This was a fun project to work on for my video series, but I'm not going to do much else on it.
If you'd like to contribute code to do these, feel free!

- Online multiplayer
- Custom skins
