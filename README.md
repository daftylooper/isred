# isred
my barebones implementation of redis.

# what is this
isred is a simple key-value store hosted up on server. It can cache data in-memory and can connect to other services via a TCP connection.

# how to use

Requirements - `go(>=1.21.3)`, `git`

1) Clone the repository - `git clone https://github.com/daftylooper/isred`
2) Run `main.go` - `go run main.go`. spins up a TCP server on port 3000.
3) Connect to the TCP server, say, telnet - `telnet localhost 3000`
4) Enjoy! send commands back and get reply
5) Optional: star this repo.

