package module

import "flag"

var (
	ServerMode = false
	ServerPort string
	Nickname   string
)

func init() {
	// flag parse
	flag.BoolVar(&ServerMode, "server", false, "Run as server mode")
	flag.StringVar(&ServerPort, "port", "localhost:4399", "Server address")
	flag.StringVar(&Nickname, "nickname", "alice", "Your nickname")
	flag.Parse()
}
