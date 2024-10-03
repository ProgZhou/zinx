package main

import (
	"log"
	"zinx/ziface"
	"zinx/znet"
)

//zinx应用程序

func main() {
	iServer := znet.NewServer("[zinx-v4.0]")
	//给服务器添加自定义router
	iServer.AddRouter(&PingRouter{})
	//启动服务器
	iServer.Serve()
}

//自定义路由
type PingRouter struct {
	znet.BaseRouter
}

func (p *PingRouter) PreHandle(request ziface.IRequest) {
	log.Println("call router pre handle...")
	_, err := request.GetConnection().GetConnection().Write([]byte("before ping...\n"))
	if err != nil {
		log.Println("call router pre error.." + err.Error())
	}
}

func (p *PingRouter) Handle(request ziface.IRequest) {
	log.Println("call router handle...")
	_, err := request.GetConnection().GetConnection().Write([]byte("ping...ping...ping...\n"))
	if err != nil {
		log.Println("call router handle error.." + err.Error())
	}
}

func (p *PingRouter) PostHandle(request ziface.IRequest) {
	log.Println("call router post handle...")
	_, err := request.GetConnection().GetConnection().Write([]byte("after ping...\n"))
	if err != nil {
		log.Println("call router post error.." + err.Error())
	}
}
