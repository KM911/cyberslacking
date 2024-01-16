package network

import (
	"fmt"
	"net"
	"os"
	"time"
)

func CreateUDPServer(_address string) *net.UDPConn {

	udpAddress, err := net.ResolveUDPAddr("udp", _address)
	if err != nil {
		fmt.Println("无法解析UDP地址:", err)
		return nil
	}

	// 创建的UPD地址内容为
	fmt.Println("创建的UPD地址内容为", udpAddress)

	// 创建UDP连接
	// 需要一个udpAddress对象 其实就是 ip + port
	conn, err := net.ListenUDP("udp", udpAddress)
	if err != nil {
		fmt.Println("无法监听UDP地址:", err)
		return nil
	}
	return conn

}

func ServerStart() {
	server := CreateUDPServer("0.0.0.0:3000")
	buffer := make([]byte, 2048)
	for {
		n, address, _ := server.ReadFromUDP(buffer)
		// 你可能会有多个客户端连接到你的服务端,所以你应该根据其address创建一个map
		fmt.Println("remote address", address)
		// fmt.Println("message", string(buffer[:n]))

		// 这里第一次读取到的是我们的control
		controlMessage := ResolveControlMessage(buffer[:n])
		ReceiveFile(server, controlMessage)
	}
}

func ReceiveFile(_conn *net.UDPConn, fileInfo *ControlMessage) {
	start := time.Now()

	file, _ := os.Create(fileInfo.FileName)
	buffer := make([]byte, fileInfo.ChunkSize)
	index := 0
	for {
		if index >= int(fileInfo.ChunkNum) {
			break
		}
		n, _ := _conn.Read(buffer)
		// fmt.Print("\r收到第", fmt.Sprintf("%09d", index), "数据包")
		file.Write(buffer[:n])
		index++

		// 一段时间后,停止该loop
	}
	file.Close()
	fmt.Println("数据传输完成", fileInfo)
	// 校验文件的md5值
	md5 := getFileMd5(fileInfo.FileName)
	if md5 == fileInfo.FileMd5 {
		fmt.Println("md5校验通过 文件一致")
	} else {
		// 文件传输失败
		fmt.Println("文件传输失败")
	}
	// 等效传输速度为
	cost := time.Now().After(start)
	fmt.Println("花费时间为 ", cost, " 等效传输速度为", fileInfo.ChunkSize*fileInfo.ChunkNum/1024, "KB/S")

}
