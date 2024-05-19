package module

import "fmt"

func ForLoopChek() {
	for {
		for _, conn := range ConnectionPool {
			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			if err != nil {
				delete(ConnectionPool, conn.RemoteAddr().String())
				BoardCastMessage(conn.RemoteAddr().String() + "quit")
				continue
			}
			fmt.Println("receive message:", string(buf[:n]))
			BoardCastMessage(string(buf[:n]))
		}
	}
}

func ThreadsCheck() {
	// 针对每一个连接都创建一个线程/协程来处理
	// 理由是goroutine相对比较轻量级

}
