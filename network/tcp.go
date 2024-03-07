package network

import (
	"fmt"
	"net"
	"os"
	"time"
)

const (
	// 相当于 2MB
	MTU = 1024 * 1024 * 4
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
	buffer := make([]byte, MTU)
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
	fmt.Println("tcp listen ", tcpAddress)
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
	fileInfo.FileName = "TCP_" + fileInfo.FileName
	ReceiveFile(conn.conn, fileInfo)
}

func TCPClientSendFile(_address string, _src string) {
	conn := &TCPConn{}
	conn.EstablishConnection(_address)
	fileInfo := GetFileInfoMessage(_src)
	conn.SendData(fileInfo.ToBuffer())
	SendFile(conn.conn, _src, fileInfo)
}

func TCPServerFile(_address string) {

	conn := &TCPConn{}
	conn.Listener(_address)

	file, err := os.Create("test.file")
	if err != nil {
		panic(err)
	}

	buffer := make([]byte, getFileSize("readme.mp4"))
	conn.conn.Read(buffer)
	file.Write(buffer)
	file.Close()

	md5_new := getFileMd5("test.file")
	md5_old := getFileMd5("readme.mp4")

	if md5_new == md5_old {
		fmt.Println("md5校验通过 文件一致")
	} else {
		fmt.Println("文件传输失败")
	}

	// fileInfo := &FileInfoMessage{}
	// ParseGob(conn.ReceiveData(), fileInfo)
	// fileInfo.FileName = "TCP_" + fileInfo.FileName
	// ReceiveFile(conn.conn, fileInfo)
}

func TCPClientFile(_address string, _src string) {
	start := time.Now()
	conn := &TCPConn{}
	conn.EstablishConnection(_address)
	// file, _ := os.Open(_src)
	// io.Copy(conn.conn, file)

	fileSize := getFileSize(_src)
	fmt.Println("file size is ", fileSize)

	buffer := make([]byte, fileSize)
	file, _ := os.Open(_src)
	file.Read(buffer)
	conn.conn.Write(buffer)

	cost := time.Since(start)

	fmt.Println("花费时间为 ", cost, " 等效传输速度为", float64(fileSize)/1024/1024/cost.Seconds(), "MB/S")

}
