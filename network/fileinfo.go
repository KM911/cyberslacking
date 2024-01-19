package network

import (
	"bytes"
	"crypto/md5"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

type FileInfoMessage struct {
	ChunkSize uint64
	ChunkNum  uint64
	FileName  string
	FileMd5   string
}

func getFileMd5(filename string) string {
	// 文件全路径名
	path := fmt.Sprintf("./%s", filename)
	// 打开文件
	pFile, err := os.Open(path)
	if err != nil {
		fmt.Println("打开文件失败 ,", err.Error())
		return ""
	}
	defer pFile.Close()
	// 创建 MD5 哈希对象
	md5h := md5.New()
	// 读取文件内容并计算 MD5 哈希值
	io.Copy(md5h, pFile)
	// 返回 MD5 哈希值
	return hex.EncodeToString(md5h.Sum(nil))
}

func GetFileInfoMessage(_src string) *FileInfoMessage {
	msg := &FileInfoMessage{
		ChunkSize: 1400,
	}
	fileInfo, _ := os.Stat(_src)
	if uint64(fileInfo.Size())%msg.ChunkSize != 0 {
		msg.ChunkNum = uint64(fileInfo.Size())/msg.ChunkSize + 1
	} else {
		msg.ChunkNum = uint64(fileInfo.Size()) / msg.ChunkSize
	}

	// 获取文件的md5
	msg.FileMd5 = getFileMd5(_src)
	msg.FileName = "udp_" + _src

	fmt.Println("your message is ", msg)
	return msg
}

func (_msg *FileInfoMessage) ToBuffer() []byte {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(_msg)
	if err != nil {
		panic(err.Error())
	}
	bf := buf.Bytes()
	return bf
}

func ResolveControlMessage(_data []byte) *FileInfoMessage {
	_msg := &FileInfoMessage{}
	reader := bytes.NewReader(_data)

	decoder := gob.NewDecoder(reader)
	err := decoder.Decode(_msg)
	if err != nil {
		panic(err.Error())
	}
	return _msg
}
