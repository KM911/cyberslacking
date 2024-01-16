package main

import (
	"os"
	"sync"

	"github.com/KM911/demo/network"
)

var (
	MainThreadWaitGroup = sync.WaitGroup{}
)

func init() {
	MainThreadWaitGroup.Add(1)
}

func Server() {
	switch {

	case os.Args[1] == "s" || os.Args[1] == "server":
		network.ServerStart()

	// case os.Args[1] == "c" || os.Args[1] == "client":
	// 	network.ClientStart()
	default:
		network.ClientStart()
		// os.Exit(1)
	}
}

func main() {
	Server()

	// msg := network.GetControlMessage("sekiro.exe")
	// fmt.Println(msg)
	// 可以开始尝试进行法宝了
	// buffer := msg.ToBuffer()
	// // 进行反序列化

	// network.ResolveControlMessage(buffer)
}
