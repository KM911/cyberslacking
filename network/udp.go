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
	buffer := make([]byte, 1500)
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

	// 创建的UPD地址内容为
	fmt.Println("创建的UPD地址内容为", udpAddress)

	// 创建UDP连接
	// 需要一个udpAddress对象 其实就是 ip + port
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
	ReceiveFile(conn.conn, fileInfo)
}

func UDPClientSendFile(_address string, _src string) {
	conn := &UDPConn{}
	conn.EstablishConnection(_address)
	fileInfo := GetFileInfoMessage(_src)
	conn.SendData(fileInfo.ToBuffer())
	SendFile(conn.conn, _src, fileInfo)
}
