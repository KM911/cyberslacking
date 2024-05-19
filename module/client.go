package module

import (
	"fmt"
	"net"
	"os"
)

var (
	conn                 net.Conn
	HistoryMessage       []byte
	HistoryMessageLength int
)

func init() {
	HistoryMessage = make([]byte, 1024)
	HistoryMessageLength = 0
}

func ClientStart() {
	println("Client start")

	tcpAddress, err := net.ResolveTCPAddr("tcp", ServerPort)
	if err != nil {
		fmt.Println("can not resolve tcp address", err)
	}
	fmt.Println("tcp address", tcpAddress)
	conn, err = net.DialTCP("tcp", nil, tcpAddress)
	if err != nil {
		fmt.Println("failed to connect server:", err)
		return
	}
	fmt.Println("Connected to server:", conn.RemoteAddr())
	fmt.Print(">")
	// 先告知其他人自己的昵称
	// conn.Write([]byte(Nickname))
	go ListenBroadcast()
	go ListenUserInput()
	pause()
}
func ListenBroadcast() {
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Server closed")
			os.Exit(0)
		}
		if n != HistoryMessageLength && string(buf[:n]) != string(HistoryMessage[:HistoryMessageLength]) {
			fmt.Print("\rnickname>", string(buf[:n]))
			fmt.Print(">")
		}
	}
}

func SendMessages() {
	conn.Write(HistoryMessage[:HistoryMessageLength])
}

func ListenUserInput() {
	for {
		HistoryMessageLength, _ = os.Stdin.Read(HistoryMessage)
		SendMessages()
		fmt.Print(">")
	}
}
