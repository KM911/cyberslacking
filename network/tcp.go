package network

import (
	"fmt"
	"net"
	"time"
)

type TCPConn struct {
	conn *net.TCPConn
}

func (_conn *TCPConn) EstablishConnection(_address string) {
	tcpAddress, err := net.ResolveTCPAddr("tcp", _address)
	if err != nil {
		fmt.Println("无法解析TCP地址:", err)
	}
	fmt.Println("tcp address", tcpAddress)
	conn, err := net.DialTCP("tcp", nil, tcpAddress)
	if err != nil {
		fmt.Println("无法连接到服务器:", err)
		return
	}
	_conn.conn = conn

}
func (_conn *TCPConn) SendData(_data []byte) int {
	n, err := _conn.conn.Write(_data)
	if err != nil {
		panic(err)
	}
	return n
}

func (_conn *TCPConn) ReceiveData() []byte {
	buffer := make([]byte, 1500)
	n, err := _conn.conn.Read(buffer)
	if err != nil {
		return nil
	}
	return buffer[:n]
}

func (_conn *TCPConn) Listener(_address string) {
	tcpAddress, err := net.ResolveTCPAddr("tcp", _address)
	if err != nil {
		fmt.Println("无法解析TCP地址:", err)
		return
	}
	fmt.Println("start listen ", tcpAddress)
	listener, err := net.ListenTCP("tcp", tcpAddress)
	if err != nil {
		fmt.Println("无法监听TCP地址:", err)
		return
	}
	_conn.conn, err = listener.AcceptTCP()
	if err != nil {
		panic(err)
	}
}

func TCPServerStart(_address string) {
	conn := &TCPConn{}
	conn.Listener(_address)
	//	TODO
	for {
		data := conn.ReceiveData()
		fmt.Println(string(data))
	}
}

func TCPClientStart(_address string) {
	conn := &TCPConn{}
	conn.EstablishConnection(_address)
	// TODO
	for {
		conn.SendData([]byte("hello world"))
		time.Sleep(1 * time.Second)
	}
}

func TCPServerReceiveFile(_address string) {
	conn := &TCPConn{}
	conn.Listener(_address)
	fileInfo := &FileInfoMessage{}
	ParseGob(conn.ReceiveData(), fileInfo)
	ReceiveFile(conn.conn, fileInfo)
}

func TCPClientSendFile(_address string, _src string) {
	conn := &TCPConn{}
	conn.EstablishConnection(_address)
	fileInfo := GetFileInfoMessage(_src)
	conn.SendData(fileInfo.ToBuffer())
	SendFile(conn.conn, _src, fileInfo)
}
