package module

import (
	"fmt"
	"net"
	"sync"
	"time"
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

type Client struct {
	Conn     net.Conn
	Username []byte
	Key      []byte
}

func NewClient(conn net.Conn) Client {
	return Client{
		Conn:     conn,
		Username: []byte{},
		Key:      generateKey(),
	}
}

var (
	ClientPool   = make(map[string]Client)
	Pause        = sync.WaitGroup{}
	Listener     net.Listener
	MessageQueue = NewRingQueue(10)
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
		fmt.Println("new conn ", conn.RemoteAddr().String())
		client := NewClient(conn)
		ClientPool[conn.RemoteAddr().String()] = client
		conn.Write(client.Key)
		// send history
		MessageQueue.Each(func(msg []byte) {
			bs, err := encrypt(msg, client.Key)
			time.Sleep(1 * time.Second)
			if err != nil {
				fmt.Println(err)
			}
			conn.Write(bs)
		})
		go ListenSingleConnection(conn)
	}
}

func ConnectClient() {
	for {
		conn, err := Listener.Accept()
		Must(err)
		fmt.Println("new conn ", conn.RemoteAddr().String())
		newClient := NewClient(conn)
		ClientPool[conn.RemoteAddr().String()] = newClient
		conn.Write(newClient.Key)

	}
}

func BoardCastMessage(message []byte, sneder string) {
	for _, client := range ClientPool {
		if client.Conn.RemoteAddr().String() == sneder {
			continue
		}
		bs, err := encrypt(message, client.Key)
		if err != nil {
			fmt.Println(err)
		}
		client.Conn.Write(bs)
	}
}
