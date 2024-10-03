package main

import (
	"io"
	"log"
	"net"
	"time"
	"zinx/znet"
)

func main() {
	//1. 连接远程服务器，得到conn连接4
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		log.Fatalln("server conn error: ", err)
		return
	}
	//2. 向连接中写入数据
	cnt := 0
	for {
		dp := znet.NewSerialization()
		packMessage, _ := dp.Pack(znet.NewMessage(0, []byte("zinx-v0.5 client test message")))
		time.Sleep(1 * time.Second)
		if _, err := conn.Write(packMessage); err != nil {
			log.Printf("client write error: " + err.Error())
		}
		//接受服务器响应
		headMessage := make([]byte, dp.GetDataLength())
		if _, err := io.ReadFull(conn, headMessage); err != nil {
			log.Println("client read error: " + err.Error())
		}
		message, _ := dp.UnPack(headMessage)
		var data []byte
		if message.GetLen() > 0 {
			data = make([]byte, message.GetLen())
			if _, err := io.ReadFull(conn, data); err != nil {
				log.Printf("read message data error: " + err.Error())
				break
			}

		}
		message.SetData(data)
		log.Printf("receive message: messageId=[%d], messageData=[%s]\n", message.GetId(), string(message.GetData()))
		cnt++
		if cnt >= 10 {
			break
		}
	}

}
