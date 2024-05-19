//go:build client
// +build client

package app

import (
	"cyberslacking/module"
)

func Start() {
	module.ClientStart()
}
