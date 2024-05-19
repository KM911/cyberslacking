package main

import (
	"cyberslacking/app"
	"sync"
)

var (
	MainThreadWaitGroup = sync.WaitGroup{}
)

func main() {
	app.Start()
}
