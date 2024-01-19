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

// func SendFileControled(_conn *net.UDPConn, _src string, fileInfo *FileInfoMessage) {
// 	task := CreateTaskQueen(100)
// 	fmt.Println("开始发送文件")
// 	go func() {
// 		// 第一次正常发包
// 		for i := 1; i <= int(fileInfo.ChunkNum); i++ {
// 			// 这里速度需要控制一下,避免太快了不是吗?

// 			task.queen <- int64(i)
// 		}
// 		// 从_conn中读取数据,补充到剩余的任务队列中
// 		// task.queen.Close()
// 		close(task.queen)
// 		// 这里的0值感觉不太好
// 	}()

// 	file, _ := os.Open(_src)
// 	buffer := make([]byte, fileInfo.ChunkSize)
// 	var index int64
// 	for {

// 		index = <-task.queen
// 		// 正在发送数据包
// 		// fmt.Println("index : ", index)
// 		if index == 0 {
// 			break
// 		}
// 		n, _ := file.ReadAt(buffer, (index-1)*int64(fileInfo.ChunkSize))
// 		_conn.Write(ToByte(Payload{
// 			Index: uint64(index - 1),
// 			Data:  buffer[:n],
// 		}))
// 	}
// 	fmt.Println("包发送完毕")
// }

//func ClientStart() {
//	conn := EstablishConnection("127.0.0.1:3000")
//	// conn.Write([]byte("hello world"))
//	fileInfo := GetFileInfoMessage(SendFileName)
//	conn.Write(fileInfo.ToBuffer())
//
//	// SendFile(conn, SendFileName, fileInfo)
//	SendFileControled(conn, SendFileName, fileInfo)
//}
