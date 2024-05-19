//go:build default
// +build default

package app

import (
	"cyberslacking/module"
)

func Start() {
	// module.ServerStart()
	module.ServerStartGoroutine()

}
