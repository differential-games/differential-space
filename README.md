# differential-space
Differential Space is a minimalist turn-based space strategy game I made for my video series.

## How do I play this?
You'll need two windows, one to launch the server, and the other to launch the client.

### Launch the Server
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

### Play
To end the current turn, click the box in the upper-right hand corner. The color of the box indicates
whose turn it is.

I'll add instructions and rules once the game is configurable and human-playable.

## Configuring the Game
Configure the game with a JSON configuration file.

For examples, see sample [easy](server/easy.json) and [hard](server/hard.json) configurations.
The server prints out the configuration it is using on startup.

Option definitions:
- **difficulty** `[0.0, 1.0]`, is how intelligent the AI players are.
AI opponents receive no benefits whatsoever.
Higher is more difficult. I haven't won against an AI at `1.0` yet.
- **humanPlayers** is the IDs of the human players.
The server automatically acts for non-human Players.
- **minRadius** is the minimum radius of Planets in the system.
- **numPlayers** is the number of Players in the game.
The server supports as many Players as you want, but the client only supports up to 4.
- **numPlanets** is the number of Planets in the game.
- **radius** is the maximum distance a Planet maybe from the center of the map.

## Troubleshooting
If you just see an empty map, the game isn't connecting to the server properly. This may be
because port 8080 is in use. The game doesn't yet support modifying what port the server/client
communicate through.