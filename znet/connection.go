package znet

import (
	"errors"
	"io"
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
	//消息管理模块
	Handler ziface.IMessageHandler
}

//初始化连接的方法
func NewConnection(conn *net.TCPConn, connId uint32, messageHandler ziface.IMessageHandler) *Connection {
	return &Connection{
		Conn:    conn,
		ConnId:  connId,
		IsClose: false,
		//handleApi: callback,
		ExitChan: make(chan bool, 1),
		Handler:  messageHandler,
	}
}

func (c *Connection) StartReader() {
	log.Printf("conn {%d} reader is running\n", c.ConnId)
	//完成业务之后关闭连接
	defer c.Stop()
	//从客户端读取数据
	for {
		//创建一个拆包、解包的类
		dp := NewSerialization()

		//读取消息头
		messageHead := make([]byte, dp.GetDataLength())
		if _, err := io.ReadFull(c.GetConnection(), messageHead); err != nil {
			log.Printf("read from client error: {%s}\n", err)
			c.ExitChan <- true
			break
		}
		//读取消息长度和消息id
		clientMessage, err := dp.UnPack(messageHead)
		if err != nil {
			log.Printf("data unpack error: {%s}\n", err.Error())
			c.ExitChan <- true
			break
		}

		//根据消息长度，再次读取data内容
		var data []byte
		if clientMessage.GetLen() > 0 {
			data = make([]byte, clientMessage.GetLen())
			if _, err := io.ReadFull(c.GetConnection(), data); err != nil {
				log.Printf("data unpack error: {%s}\n", err.Error())
				c.ExitChan <- true
				break
			}
		}
		clientMessage.SetData(data)
		//得到当前Conn的Request数据
		req := Request{
			conn:    c,
			message: clientMessage,
		}
		//根据绑定号的messageId调度对应的处理方法
		go c.Handler.DoMessageHandler(&req)
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

//发送数据时，按照封包发送
func (c *Connection) Send(messageId uint32, data []byte) error {
	if c.IsClose {
		return errors.New("connection is closed")
	}
	//将消息进行封装
	dp := NewSerialization()
	messagePack, err := dp.Pack(NewMessage(messageId, data))
	if err != nil {
		log.Printf("pack message error, message id: [%d]\n", messageId)
		return errors.New("message package error")
	}
	//将数据发送给客户端
	if _, err = c.Conn.Write(messagePack); err != nil {
		log.Printf("message write error, message id: [%d]\n", messageId)
		return errors.New("connection write error")
	}
	return nil
}
