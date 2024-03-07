package main

import (
	"flag"
	"fmt"
	"sync"

	"github.com/KM911/demo/network"
)

var (
	MainThreadWaitGroup = sync.WaitGroup{}
	Type                string
)

func init() {
	//MainThreadWaitGroup.Add(1)
	flag.StringVar(&Type, "type", "tcp", "tcp or udp or atp")
	flag.Parse()
}

//因为你需要即转发udp和tcp的流量啊 不然你怎么进行中间件的开发呢? 不是吗 ?

func TCP() {
	switch len(flag.Args()) {
	case 1:
		// network.TCPClientStart("127.0.0.1:3000")
		network.TCPClientSendFile("127.0.0.1:3000", "readme.mp4")
	default:
		network.TCPServerReceiveFile("0.0.0.0:3000")
	}

}

func UDP() {
	switch len(flag.Args()) {
	case 1:
		network.UDPClientSendFile("127.0.0.1:3000", "readme.mp4")
	default:
		network.UDPServerReceiveFile("0.0.0.0:3000")
	}
}

func UDPPingPong() {
	switch len(flag.Args()) {
	case 1:
		network.UDPPingServer("127.0.0.1:3000")
	default:
		network.UDPPangServer("0.0.0.0:3000")
	}
}
func ATP() {
	switch len(flag.Args()) {
	case 1:
		network.ATPClientSendFile("127.0.0.1:3000", "readme.mp4")
	default:
		network.ATPServerReceiveFile("0.0.0.0:3000")
	}

}

func Http() {
	switch len(flag.Args()) {
	case 1:
		network.TCPServer("0.0.0.0:7890")
	default:
		network.HttpServer(":7890")
	}

}

func TCPFile() {
	switch len(flag.Args()) {
	case 1:
		network.TCPClientFile("127.0.0.1:3000", "readme.mp4")
	default:
		network.TCPServerFile("0.0.0.0:3000")
	}
}

// 为了表示其稳定性 我们进行发送文件的测试
func main() {
	fmt.Println("flag.Args()", flag.Args())
	switch Type {
	case "tcp-file":
		TCPFile()
	case "tcp":
		TCP()
	case "udp":
		UDP()
		// UDPPingPong()
	case "atp":
		ATP()
	case "http":
		Http()
	default:
		panic("type error")
	}
}
