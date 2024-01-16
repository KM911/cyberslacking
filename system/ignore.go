package system

import (
	"os"
	"os/signal"
	"syscall"
)

func DisableInterpret() {
	c := make(chan os.Signal, 10)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	for {
		<-c
	}
}
