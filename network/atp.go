package network

import (
	"fmt"
	"net"
	"os"
	"time"
)

type ATPConn struct {
	conn *net.UDPConn
}

func (_conn *ATPConn) EstablishConnection(_address string) {
	atpAddress, err := net.ResolveUDPAddr("udp", _address)
	if err != nil {
		fmt.Println("无法解析ATP地址:", err)
	}

	// 创建的UPD地址内容为
	fmt.Println("创建的UPD地址内容为", atpAddress)

	// 创建ATP连接
	// 需要一个atpAddress对象 其实就是 atp + port
	conn, err := net.DialUDP("udp", nil, atpAddress)
	if err != nil {
		fmt.Println("无法监听ATP地址:", err)
	}
	_conn.conn = conn
}
func (_conn *ATPConn) SendData(_data []byte) int {
	n, err := _conn.conn.Write(_data)
	if err != nil {
		panic(err)
	}
	return n
}

func (_conn *ATPConn) ReceiveData() []byte {
	buffer := make([]byte, MTU)
	n, err := _conn.conn.Read(buffer)
	if err != nil {
		return nil
	}
	return buffer[:n]
}

var (
	ControlChan = make(chan int, 1)
)

func (_conn *ATPConn) Listener(_address string) {
	atpAddress, err := net.ResolveUDPAddr("udp", _address)
	if err != nil {
		fmt.Println("无法解析ATP地址:", err)
		return
	}
	fmt.Println("atpAddress", atpAddress)
	conn, err := net.ListenUDP("udp", atpAddress)
	if err != nil {
		fmt.Println("无法监听ATP地址:", err)
		return
	}
	_conn.conn = conn
}

func ATPSendFile(_conn *ATPConn, _src string, _fileInfo *FileInfoMessage) {
	buffer := make([]byte, _fileInfo.ChunkSize)
	file, err := os.Open(_fileInfo.FileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for i := 0; i < int(_fileInfo.ChunkNum); i++ {
		read, err := file.Read(buffer)
		if err != nil {
			panic(err)
		}

		_conn.conn.Write(buffer[:read])
		_conn.conn.Read(buffer)
	}
}

func ATPReceiveFile(_conn *ATPConn, _fileInfo *FileInfoMessage) {
	go func() {
		time.Sleep(5 * time.Second)
		panic("file transfer failed")
	}()
	start := time.Now()

	buffer := make([]byte, _fileInfo.ChunkSize)
	file, err := os.Create(_fileInfo.FileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Println("ready to receive file")
	for i := 0; i < int(_fileInfo.ChunkNum); i++ {
		n, address, err := _conn.conn.ReadFromUDP(buffer)
		file.Write(buffer[:n])
		if err != nil {
			panic(err)
		}
		_conn.conn.WriteToUDP([]byte("1"), address)
	}

	fileCheck(_fileInfo)
	cost := time.Since(start)
	fmt.Println("花费时间为 ", cost, " 等效传输速度为", float64(_fileInfo.ChunkSize*_fileInfo.ChunkNum)/1024/1024/cost.Seconds(), "MB/S")
}

func ATPServerReceiveFile(_address string) {
	conn := &ATPConn{}
	conn.Listener(_address)
	fileInfo := &FileInfoMessage{}
	ParseGob(conn.ReceiveData(), fileInfo)
	fileInfo.FileName = "atp_" + fileInfo.FileName
	ATPReceiveFile(conn, fileInfo)
}

func ATPClientSendFile(_address string, _src string) {
	conn := &ATPConn{}
	conn.EstablishConnection(_address)
	fileInfo := GetFileInfoMessage(_src)
	conn.SendData(fileInfo.ToBuffer())
	ATPSendFile(conn, _src, fileInfo)
}
