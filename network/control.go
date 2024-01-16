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

type ControlMessage struct {
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

func GetControlMessage(_src string) *ControlMessage {
	msg := &ControlMessage{
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

func (_msg *ControlMessage) ToBuffer() []byte {
	// 将其序列化
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(_msg)
	if err != nil {
		panic(err.Error())
	}
	bf := buf.Bytes()
	fmt.Println("序列化后的数据 : ", bf)
	return bf
}

func ResolveControlMessage(_data []byte) *ControlMessage {
	_msg := &ControlMessage{}
	// fmt.Println("your data is ", _data)
	// 我们需要将 _data 变成一个io.Reader
	reader := bytes.NewReader(_data)
	fmt.Println(reader)

	decoder := gob.NewDecoder(reader)
	err := decoder.Decode(_msg)
	if err != nil {
		fmt.Println(err.Error())
	}
	// fmt.Println("反序列化后的数据", _msg)
	return _msg
}
