package main

import (
	"log"
	"net"
	"time"
)

func main() {
	//1. 连接远程服务器，得到conn连接4
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Fatalln("server conn error: ", err)
		return
	}
	//2. 向连接中写入数据
	cnt := 0
	for {
		buffer := []byte("Hello Server")
		_, err := conn.Write(buffer)
		if err != nil {
			log.Fatalln("write error: ", err)
			return
		}
		time.Sleep(1 * time.Second)

		//接受服务器响应
		receiveBuf := make([]byte, 521)
		read, err := conn.Read(receiveBuf)
		if err != nil {
			log.Fatalln("receive from server error: ", err)
			return
		}
		log.Printf("server: %s\n", string(receiveBuf[:read]))
		cnt++
		if cnt >= 10 {
			break
		}
	}

}
