package module

import (
	"fmt"
	"net"
	"sync"
)

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func pause() {
	Pause.Add(1)
	Pause.Wait()
}

var (
	ConnectionPool = make(map[string]net.Conn)
	Pause          = sync.WaitGroup{}
	Listener       net.Listener
)

func ServerStart() {
	println("Server start ,for loop check")
	Listener, _ = net.Listen("tcp", ServerPort)
	// Must(err)
	go ConnectClient()
	go ForLoopChek()
	pause()
}

func ServerStartGoroutine() {
	println("Server start,goroutine for every network")
	Listener, _ = net.Listen("tcp", ServerPort)
	// Must(err)
	for {
		conn, err := Listener.Accept()
		Must(err)
		ConnectionPool[conn.RemoteAddr().String()] = conn
		go HandleMessageGoroutine(conn)
	}
	// go ConnectClient()
	// go HandleMessage()
	// pause()
}

func HandleMessageGoroutine(conn net.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			delete(ConnectionPool, conn.RemoteAddr().String())
			BoardCastMessage(conn.RemoteAddr().String() + "quit")
			// free the resource
			return
		}
		fmt.Println("receive message:", string(buf[:n]))
		BoardCastMessage(string(buf[:n]))
	}
}

func HandleMessage() {
	// 这里涉及到多路复用的问题,我们这里是对每一个连接都进行了监听
	// TODO 优化nio
	ForLoopChek()
}

func ConnectClient() {
	for {
		conn, err := Listener.Accept()
		Must(err)
		fmt.Println("new connetion:", conn.RemoteAddr())
		ConnectionPool[conn.RemoteAddr().String()] = conn
	}
}

func BoardCastMessage(message string) {
	for _, conn := range ConnectionPool {
		conn.Write([]byte(message))
	}
}
