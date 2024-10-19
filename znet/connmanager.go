package znet

import (
	"errors"
	"log"
	"sync"
	"zinx/ziface"
)

//连接管理模块实现层
type ConnManager struct {
	//已经建立的连接集合
	connections map[uint32]ziface.IConnection
	//针对map更改的互斥锁
	mapLock sync.RWMutex
}

func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint32]ziface.IConnection),
	}
}

func (c *ConnManager) AddConnection(conn ziface.IConnection) {
	//保护共享资源，加写锁
	c.mapLock.Lock()
	//释放锁
	defer c.mapLock.Unlock()
	c.connections[conn.GetConnID()] = conn
	log.Printf("connection[%d] add to manager success.\n", conn.GetConnID())
}

func (c *ConnManager) RemoveConnection(conn ziface.IConnection) {
	//保护共享资源，加写锁
	c.mapLock.Lock()
	defer c.mapLock.Unlock()
	delete(c.connections, conn.GetConnID())
	log.Printf("connection[%d] remove from manager.\n", conn.GetConnID())
}

func (c *ConnManager) GetConnection(connId uint32) (ziface.IConnection, error) {
	//加读锁
	c.mapLock.RLock()
	defer c.mapLock.RUnlock()
	if conn, ok := c.connections[connId]; ok {
		return conn, nil
	} else {
		return nil, errors.New("connection not found")
	}
}

func (c *ConnManager) Size() int {
	return len(c.connections)
}

func (c *ConnManager) ClearConnection() {
	c.mapLock.Lock()
	defer c.mapLock.Unlock()
	for connId, conn := range c.connections {
		conn.Stop()
		delete(c.connections, connId)
	}
	log.Printf("clear all connections success\n")
}
