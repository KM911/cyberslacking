package module

import (
	"fmt"
	"net"
)

func ForLoopChek() {
	for {
		for _, client := range ClientPool {
			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			if err != nil {
				delete(ClientPool, client.Conn.RemoteAddr().String())
				BoardCastMessage([]byte(conn.RemoteAddr().String()+"quit"), "")
				continue
			}
			BoardCastMessage(buf[:n], "")
		}
	}
}

func ListenSingleConnection(conn net.Conn) {
	buf := make([]byte, 1024)
	client := ClientPool[conn.RemoteAddr().String()]
	key := ClientPool[conn.RemoteAddr().String()].Key
	for {
		n, err := conn.Read(buf)
		if err != nil {
			delete(ClientPool, conn.RemoteAddr().String())
			BoardCastMessage([]byte(conn.RemoteAddr().String()+"quit"), "")
			return
		}
		// decry
		bs, err := decrypt(buf[:n], key)
		if err != nil {
			fmt.Println("decrypt error", err)
			continue
		}
		msg := ResolveMessage(bs)
		// fmt.Println("msg is ", msg.Action, msg.Content)
		switch msg.Action {
		case "chat":
			bs := append(client.Username, msg.Content...)
			MessageQueue.Enqueue(bs)
			BoardCastMessage(bs, conn.RemoteAddr().String())
		case "list":
			// conn.Write([]byte("list"))
			users := []byte{}
			for _, v := range ClientPool {
				// connection two slice
				fmt.Println("username is ", string(v.Username))
				users = append(v.Username, []byte("  ")...)
				users = append(users, []byte(v.Conn.RemoteAddr().String())...)
				users = append(users, []byte("\t")...)
			}
			bs, err := encrypt(users, key)
			if err != nil {
				fmt.Println("encrypt error", err)
			}
			conn.Write(bs)
		case "rename":
			// ClientPool[conn.RemoteAddr().String()].Username = string(msg.Content)
			fmt.Println("rename")
			client.Username = msg.Content
			ClientPool[conn.RemoteAddr().String()] = client
			// for _, v := range ClientPool {
			// if v.Conn.RemoteAddr().String() == conn.RemoteAddr().String() {
			// v.Username = msg.Content
			// break
			// }
			// }

		}
	}
}

func ThreadsCheck() {
	// 针对每一个连接都创建一个线程/协程来处理
	// 理由是goroutine相对比较轻量级

}
