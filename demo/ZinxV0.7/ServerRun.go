package main

import (
	"log"
	"zinx/ziface"
	"zinx/znet"
)

//zinx应用程序

func main() {
	iServer := znet.NewServer("[zinx-v5.0]")
	//给服务器添加自定义router
	iServer.AddRouter(0, &PingRouter{})
	iServer.AddRouter(1, &HelloRouter{})
	//启动服务器
	iServer.Serve()
}

//自定义路由
type PingRouter struct {
	znet.BaseRouter
}

func (p *PingRouter) Handle(request ziface.IRequest) {
	log.Println("call ping router handle...")
	log.Printf("read from client: messageId={%d}, messageData={%s}\n",
		request.GetMessageId(), string(request.GetData()))

	err := request.GetConnection().Send(200, []byte("ping...ping...ping..."))
	if err != nil {
		log.Printf("write error: {%s}\n", err.Error())
	}
}

type HelloRouter struct {
	znet.BaseRouter
}

func (h *HelloRouter) Handle(request ziface.IRequest) {
	log.Println("call hello router handle...")
	log.Printf("read from client: messageId={%d}, messageData={%s}\n",
		request.GetMessageId(), string(request.GetData()))

	err := request.GetConnection().Send(201, []byte("hello server..."))
	if err != nil {
		log.Printf("write error: {%s}\n", err.Error())
	}
}
