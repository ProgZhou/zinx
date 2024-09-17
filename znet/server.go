package znet

import (
	"errors"
	"fmt"
	"log"
	"net"
	"zinx/ziface"
)

//具体实现层：IServer接口的实现，定义一个server服务器
type Server struct {
	//服务器名称
	Name string
	//服务器绑定的ip版本
	IpVersion string
	//服务器监听的ip地址
	IP string
	//服务器监听的端口号
	Port int
	//路由
	Router ziface.IRouter
}

//定义当前客户端连接所绑定的api TODO 由用户自定义
func Callback(conn *net.TCPConn, data []byte, cnt int) error {
	log.Println("[conn handle] callback client...")
	if _, err := conn.Write(data[:cnt]); err != nil {
		log.Printf("write buffer error: {%s}\n", err)
		return errors.New("callback error: " + err.Error())
	}
	return nil
}

func (s *Server) Start() {
	log.Printf("[server starting]server listen on ip: %s, port: %d\n", s.IP, s.Port)
	go func() {
		//1. 获取一个TCP的地址
		addr, err := net.ResolveTCPAddr(s.IpVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			log.Println("resolve tcp addr error: ", err)
			return
		}
		//2. 监听服务器地址
		listener, err := net.ListenTCP(s.IpVersion, addr)
		if err != nil {
			log.Println("listen error: ", err)
			return
		}
		log.Printf("[server start]server {%s} start success\n", s.Name)
		var cid uint32
		cid = 0
		//3. 阻塞等待客户端连接，处理客户端连接业务
		for {
			//等待客户端连接，如果有连接，则返回   阻塞方法
			conn, err := listener.AcceptTCP()
			if err != nil {
				log.Println("[server listening]server listen error: ", err)
				continue
			}
			//将得到的TCP连接封装成自定义的Connection
			clientConn := NewConnection(conn, cid, Callback)
			cid++
			//启动当前的连接业务处理
			go clientConn.Start()
		}
	}()

}

func (s *Server) Serve() {
	s.Start()
	//TODO 做一些启动服务器之后的额外业务

	//serve方法本身应阻塞
	select {}
}

func (s *Server) Stop() {
	//TODO 停止服务器，将一些服务器的资源、状态或者一些已经开辟的连接进行回收
}

func (c *Server) AddRouter(router ziface.IRouter) {
	c.Router = router
}

//初始化server的方法
func NewServer(name string) ziface.IServer {
	return &Server{
		Name:      name,
		IpVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8888,
		Router:    nil,
	}
}
