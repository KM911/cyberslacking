package network

import (
	"fmt"
	"net"
	"os"
	"time"
)

func SendFile(_conn net.Conn, _src string, fileInfo *FileInfoMessage) {
	file, _ := os.Open(_src)
	buffer := make([]byte, fileInfo.ChunkSize)
	index := 0
	for {
		if index >= int(fileInfo.ChunkNum) {
			break
		}
		n, _ := file.Read(buffer)
		_conn.Write(buffer[:n])
		index++
	}
	fmt.Println("共发送", index, "个数据包")
}

func ReceiveFile(_conn net.Conn, fileInfo *FileInfoMessage) {
	go func() {
		time.Sleep(5 * time.Second)
		panic("file transfer failed")
	}()
	start := time.Now()
	file, _ := os.Create(fileInfo.FileName)
	buffer := make([]byte, fileInfo.ChunkSize)
	index := 0
	for {
		if index >= int(fileInfo.ChunkNum) {
			break
		}
		n, _ := _conn.Read(buffer)
		file.Write(buffer[:n])
		index++
	}
	file.Close()

	// 校验文件的md5值
	CheckFileMd5(fileInfo.FileName, fileInfo.FileMd5)

	// 等效传输速度为
	cost := time.Since(start)
	fmt.Println("花费时间为 ", cost, " 等效传输速度为", float64(fileInfo.ChunkSize*fileInfo.ChunkNum)/1024/1024/cost.Seconds(), "MB/S")
}

func CheckFileMd5(_src, _md5 string) {
	md5 := getFileMd5(_src)
	if md5 == _md5 {
		fmt.Println("md5校验通过 文件一致")
	} else {
		// 文件传输失败
		fmt.Println("文件传输失败")
	}
}
func fileCheck(fileInfo *FileInfoMessage) {
	fmt.Println("数据传输完成", fileInfo)
	// 校验文件的md5值
	md5 := getFileMd5(fileInfo.FileName)
	if md5 == fileInfo.FileMd5 {
		fmt.Println("md5校验通过 文件一致")
	} else {
		// 文件传输失败
		fmt.Println("文件传输失败")
	}
}
