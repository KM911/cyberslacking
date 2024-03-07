package network

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func HttpServer(_address string) {
	// 创建一个http server 支持用post接受文件

	http.HandleFunc("/ip", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method)
		if r.Method == "GET" {
			w.Header().Set("Access-Control-Allow-Origin", "*")

			w.Write([]byte(r.RemoteAddr))

		}
	})
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method)
		if r.Method == "GET" {
			w.Header().Set("Access-Control-Allow-Origin", "*")

			w.Write([]byte("hello world"))

		}
	})

	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method)
		if r.Method == "POST" {
			// 接受文件
			file, fileHeader, err := r.FormFile("file")

			if err != nil {
				panic(err)
			}
			fmt.Println(fileHeader.Filename)
			fmt.Println(fileHeader.Size)
			fmt.Println(fileHeader.Header)
			fmt.Println(fileHeader.Header["Content-Type"])
			// 创建文件
			f, err := os.Create("Http_" + fileHeader.Filename)
			if err != nil {
				panic(err)
			}
			defer f.Close()
			// 写入文件
			_, err = io.Copy(f, file)
			if err != nil {
				panic(err)
			}
			// 返回文件
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Write([]byte("上传成功"))
		}

	})
	// 设置允许跨域访问

	http.ListenAndServe(_address, nil)
}

// 尝试使用TCP来模拟http服务器

func TCPServer(_address string) {
	conn := &TCPConn{}
	conn.Listener(_address)
	buffer := make([]byte, MTU)
	for {
		// buffer := conn.ReceiveData()
		// fmt.Println("buffer : ", string(buffer))
		// 模拟http的报文格式需要你

		n, err := conn.conn.Read(buffer)
		if err != nil {
			fmt.Println("read error : ", err)
			// panic(err)读取完成了不是
			// fmt.Println("err buffer : ", string(buffer[:n]))
		}
		fmt.Println("buffer : ", string(buffer[:n]))

		conn.conn.Write([]byte("HTTP/1.1 200 OK\r\nHost: localhost:3000\r\nContent-length: " + strconv.Itoa(n) + "\r\nContent-Type: text/html\r\n\r\n" + string(buffer[:n])))
		fmt.Println("写入请求成功")
		time.Sleep(1 * time.Second)

		// n, err := conn.conn.Read(buffer)
		// if err != nil {
		// 	panic(err)
		// }
		// fmt.Println("buffer : ", string(buffer[:n]))

	}
}
