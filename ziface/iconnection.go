package ziface

import "net"

//定义抽象的连接模块
type IConnection interface {
	//启动连接 让当前的连接开始工作
	Start()
	//停止连接 结束当前连接的工作
	Stop()
	//获取当前连接的conn套接字
	GetConnection() *net.TCPConn
	//获取当前连接模块的连接id
	GetConnID() uint32
	//获取远程客户端的TCP状态
	RemoteAddr() net.Addr
	//发送数据 将数据发送给远程的客户端
	Send(data []byte) error
}

//定义一个处理连接的业务方法
type HandleFunc func(*net.TCPConn, []byte, int) error
