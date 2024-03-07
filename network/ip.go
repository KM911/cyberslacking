package network

import (
	"fmt"
	"net"
	"time"
)

// 这个IP conn 又是谁的部下 也太离谱了吧

type IPConn struct {
	conn *net.IPConn
}

func (_conn *IPConn) EstablishConnection(_address string) {
	ipAddress, err := net.ResolveIPAddr("ip", _address)
	if err != nil {
		fmt.Println("无法解析IP地址:", err)
	}

	// 创建的UPD地址内容为
	fmt.Println("创建的UPD地址内容为", ipAddress)

	// 创建IP连接
	// 需要一个ipAddress对象 其实就是 ip + port
	conn, err := net.ListenIP("ip", ipAddress)
	if err != nil {
		fmt.Println("无法监听IP地址:", err)
	}
	_conn.conn = conn
}
func (_conn *IPConn) SendData(_data []byte) int {
	n, err := _conn.conn.Write(_data)
	if err != nil {
		panic(err)
	}
	return n
}

func (_conn *IPConn) ReceiveData() []byte {
	buffer := make([]byte, MTU)
	n, err := _conn.conn.Read(buffer)
	if err != nil {
		return nil
	}
	return buffer[:n]
}

func (_conn *IPConn) Listener(_address string) {
	ipAddress, err := net.ResolveIPAddr("ip", _address)
	if err != nil {
		fmt.Println("无法解析IP地址:", err)
		return
	}

	// 创建的UPD地址内容为
	fmt.Println("创建的UPD地址内容为", ipAddress)

	// 创建IP连接
	// 需要一个ipAddress对象 其实就是 ip + port
	conn, err := net.ListenIP("ip", ipAddress)
	if err != nil {
		fmt.Println("无法监听IP地址:", err)
		return
	}
	_conn.conn = conn
}

func IPServerStart(_address string) {
	conn := &IPConn{}
	conn.Listener(_address)
	//	TODO
	for {
		data := conn.ReceiveData()
		fmt.Println(string(data))
	}
}

func IPClientStart(_address string) {
	conn := &IPConn{}
	conn.EstablishConnection(_address)
	// TODO
	for {
		conn.SendData([]byte("hello world"))
		time.Sleep(2 * time.Second)
	}
}
