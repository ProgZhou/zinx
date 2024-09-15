package main

import "zinx/znet"

//zinx应用程序

func main() {
	iServer := znet.NewServer("[zinx-v1.0]")
	//启动服务器
	iServer.Serve()
}
