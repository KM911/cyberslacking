package network

import (
	"fmt"
	"net"
	"time"
)

type UDPConn struct {
	conn *net.UDPConn
}

func (_conn *UDPConn) EstablishConnection(_address string) {
	udpAddress, err := net.ResolveUDPAddr("udp", _address)
	if err != nil {
		fmt.Println("无法解析UDP地址:", err)
	}

	fmt.Println("udpAddress", udpAddress)
	conn, err := net.DialUDP("udp", nil, udpAddress)
	if err != nil {
		fmt.Println("无法连接到服务器:", err)
		return
	}
	_conn.conn = conn
}
func (_conn *UDPConn) SendData(_data []byte) int {
	n, err := _conn.conn.Write(_data)
	if err != nil {
		panic(err)
	}
	return n
}

func (_conn *UDPConn) ReceiveData() []byte {
	buffer := make([]byte, MTU)
	n, err := _conn.conn.Read(buffer)
	if err != nil {
		return nil
	}
	return buffer[:n]
}

func (_conn *UDPConn) Listener(_address string) {
	udpAddress, err := net.ResolveUDPAddr("udp", _address)
	if err != nil {
		fmt.Println("无法解析UDP地址:", err)
		return
	}
	fmt.Println("udp listen ", udpAddress)
	conn, err := net.ListenUDP("udp", udpAddress)
	if err != nil {
		fmt.Println("无法监听UDP地址:", err)
		return
	}
	_conn.conn = conn
}

func UDPServerStart(_address string) {
	conn := &UDPConn{}
	conn.Listener(_address)
	//	TODO
	for {
		data := conn.ReceiveData()
		fmt.Println(string(data))
	}
}

func UDPClientStart(_address string) {
	conn := &UDPConn{}
	conn.EstablishConnection(_address)
	// TODO

	for {
		conn.SendData([]byte("hello world"))
		time.Sleep(1 * time.Second)
	}
}

func UDPServerReceiveFile(_address string) {
	conn := &UDPConn{}
	conn.Listener(_address)
	fileInfo := &FileInfoMessage{}
	ParseGob(conn.ReceiveData(), fileInfo)
	fileInfo.FileName = "UDP_" + fileInfo.FileName
	ReceiveFile(conn.conn, fileInfo)
}

func UDPClientSendFile(_address string, _src string) {
	conn := &UDPConn{}
	conn.EstablishConnection(_address)
	fileInfo := GetFileInfoMessage(_src)
	conn.SendData(fileInfo.ToBuffer())
	SendFile(conn.conn, _src, fileInfo)
}

func UDPPingServer(_address string) {

	conn := &UDPConn{}
	buffer := make([]byte, MTU)
	conn.EstablishConnection(_address)
	for {
		conn.SendData([]byte("ping"))
		n, err := conn.conn.Read(buffer)
		if err != nil {
			panic(err)
		}
		fmt.Println("receive from ", conn.conn.RemoteAddr(), " data is ", string(buffer[:n]))
		time.Sleep(1 * time.Second)
	}

}

func UDPPangServer(_address string) {
	conn := &UDPConn{}
	buffer := make([]byte, MTU)
	conn.Listener(_address)
	for {
		n, address, err := conn.conn.ReadFromUDP(buffer)
		if err != nil {
			panic(err)
		}
		fmt.Println("receive from ", address, " data is ", string(buffer[:n]))
		conn.conn.WriteToUDP([]byte("pang"), address)
	}
}
