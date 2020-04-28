# differential-space
Differential Space is a minimalist turn-based space strategy game I made for my video series.

## How do I play this?
For now, it isn't human-playable without modifying the server code, but I'll be adding a way to soon.
All you can do now is watch AI opponents play against each other. Again, this will change very soon.

You'll need two windows, one to launch the server, and the other to launch the client.

### Launch the Server
1. [Install Go](https://golang.org/doc/install). Remember to add GOPATH to your PATH.
1. Build and launch the server.

   ```
   cd server
   go install differential-space.go && differential-space
   ```

### Install and Launch the Client
1. Install `differential-space-client.exe`. You'll have the option to immediately launch
differential-space-client, and add it to your Start Menu.
1. Otherwise, you'll be able to find it in `C:\Program Files (x86)\Made in GameMaker Studio 2`.

### Play
To end the current turn, click the box in the upper-right hand corner. The color of the box indicates
whose turn it is.

I'll add instructions and rules once the game is configurable and human-playable.

### Troubleshooting
If you just see an empty starfield, the game isn't connecting to the server properly. This may be
because port 8080 is in use. The game doesn't yet support modifying what port the server/client
communicate through.
