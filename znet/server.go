package znet

import (
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
		//3. 阻塞等待客户端连接，处理客户端连接业务
		for {
			//等待客户端连接，如果有连接，则返回   阻塞方法
			conn, err := listener.AcceptTCP()
			if err != nil {
				log.Println("[server listening]server listen error: ", err)
				continue
			}
			//对已经连接的客户端做一些业务
			go func() {
				//另起一个协程读取客户端传送的数据
				for {
					buffer := make([]byte, 512)
					//cnt表示读取的数据大小
					cnt, err2 := conn.Read(buffer)
					if err2 != nil {
						log.Println("[server receive]server receive from client error: ", err)
						continue
					}
					log.Printf("receive from client: %s\n", string(buffer))
					//发送给客户端
					if _, err2 := conn.Write(buffer[:cnt]); err2 != nil {
						log.Println("[server write]server write error: ", err2)
						continue
					}
				}
			}()
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

//初始化server的方法
func NewServer(name string) ziface.IServer {
	return &Server{
		Name:      name,
		IpVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8888,
	}
}
