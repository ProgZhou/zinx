package znet

import (
	"log"
	"net"
	"zinx/ziface"
)

//具体的连接实现
type Connection struct {
	//socket TCP套接字
	Conn *net.TCPConn
	//连接id
	ConnId uint32
	//连接状态
	IsClose bool
	//当前连接绑定的处理业务的api方法
	//handleApi ziface.HandleFunc
	//等待连接被动推出的channel
	ExitChan chan bool
	//该链接处理的方法router
	Router ziface.IRouter
}

//初始化连接的方法
func NewConnection(conn *net.TCPConn, connId uint32, router ziface.IRouter) *Connection {
	return &Connection{
		Conn:    conn,
		ConnId:  connId,
		IsClose: false,
		//handleApi: callback,
		ExitChan: make(chan bool, 1),
		Router:   router,
	}
}

func (c *Connection) StartReader() {
	log.Printf("conn {%d} reader is running\n", c.ConnId)
	//完成业务之后关闭连接
	defer c.Stop()
	//从客户端读取数据
	for {
		buffer := make([]byte, 512)
		_, err := c.Conn.Read(buffer)
		if err != nil {
			log.Printf("read from client error: {%s}\n", err)
			c.ExitChan <- true
			break
		}

		//if err = c.handleApi(c.Conn, buffer, cnt); err != nil {
		//	log.Printf("conn {%d} handle error: %s\n", c.ConnId, err)
		//	c.ExitChan <- true
		//	break
		//}

		//得到当前Conn的Request数据
		req := Request{
			conn: c,
			data: buffer,
		}
		//从路由中找到注册绑定的Conn对应的路由调度
		go func(request ziface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)

	}
}

//启动连接，使当前连接开始工作
func (c *Connection) Start() {
	log.Printf("conn start... conn id={%d}\n", c.ConnId)
	//启动从当前连接读数据的业务
	go c.StartReader()
	//TODO 启动从当前连接写数据的业务

	for {
		select {
		case <-c.ExitChan:
			return //如果得到退出消息，则不再阻塞
		}
	}
}

func (c *Connection) Stop() {
	log.Printf("conn stop... conn id={%d}\n", c.ConnId)
	//如果当前连接已经关闭，则返回
	if c.IsClose {
		return
	}

	//将连接状态置为已关闭
	c.IsClose = true
	//关闭tcp连接
	c.Conn.Close()
	//通知从缓冲队列读数据的业务，该链接已经关闭
	c.ExitChan <- true
	//关闭管道 回收资源
	close(c.ExitChan)
}

func (c *Connection) GetConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnId
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) Send(data []byte) error {
	return nil
}
