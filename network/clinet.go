package network

import (
	"fmt"
	"net"
	"os"
	"time"
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

func SendFile(_conn *net.UDPConn, _src string, fileInfo *ControlMessage) {
	file, _ := os.Open(_src)

	buffer := make([]byte, fileInfo.ChunkSize)
	index := 0
	for {
		if index >= int(fileInfo.ChunkNum) {
			break
		}
		if index%90 == 0 {
			time.Sleep(100 * time.Millisecond)
		}
		n, _ := file.Read(buffer)
		_conn.Write(buffer[:n])
		index++
		// 降低发送的频率 , 进行限流, 不然会将包全部丢弃

	}
	fmt.Println("共发送", index, "个数据包")
}

func ClientStart() {
	conn := EstablishConnection("127.0.0.1:3000")
	// conn.Write([]byte("hello world"))
	fileInfo := GetControlMessage(SendFileName)
	conn.Write(fileInfo.ToBuffer())

	SendFile(conn, SendFileName, fileInfo)

}
