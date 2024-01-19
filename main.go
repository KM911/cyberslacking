package main

import (
	"os"
	"sync"

	"github.com/KM911/demo/network"
)

var (
	MainThreadWaitGroup = sync.WaitGroup{}
)

func init() {
	//MainThreadWaitGroup.Add(1)

}

//因为你需要即转发udp和tcp的流量啊 不然你怎么进行中间件的开发呢? 不是吗 ?

func TCP() {
	switch len(os.Args) {
	case 2:
		// network.TCPClientStart("127.0.0.1:3000")
		network.TCPClientSendFile("127.0.0.1:3000", "readme.mp4")
	default:
		network.TCPServerReceiveFile("0.0.0.0:3000")
	}

}

func UDP() {
	switch len(os.Args) {
	case 2:
		network.UDPClientSendFile("127.0.0.1:3000", "readme.mp4")
	default:
		network.UDPServerReceiveFile("0.0.0.0:3000")
	}
}

func ATP() {
	switch len(os.Args) {
	case 2:
		network.ATPClientSendFile("127.0.0.1:3000", "readme.mp4")
	default:
		network.ATPServerReceiveFile("0.0.0.0:3000")
	}

}

// 为了表示其稳定性 我们进行发送文件的测试
func main() {
	//TCP()
	//UDP()
	ATP()
}
