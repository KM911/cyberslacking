package network

import (
	"fmt"
	"net"
)

func EstablishConnection(_serverAddress string) *net.UDPConn {
	udpAddress, err := net.ResolveUDPAddr("udp", _serverAddress)
	if err != nil {
		fmt.Println("无法解析UDP地址:", err)
		return nil
	}

	// 创建UDP连接 其实这个是就是无意义的 并不存在udp连接这种说法
	conn, err := net.DialUDP("udp", nil, udpAddress)
	if err != nil {
		fmt.Println("无法连接到服务器:", err)
		return nil
	}
	return conn
}
