package znet

import (
	"fmt"
	"log"
	"zinx/ziface"
)

//消息管理模块具体实现层
type MessageHandler struct {

	//消息id与路由的映射
	messageApi map[uint32]ziface.IRouter
}

//初始化MessageHandler的方法
func NewMessageHandler() *MessageHandler {
	return &MessageHandler{messageApi: make(map[uint32]ziface.IRouter)}
}

func (m *MessageHandler) DoMessageHandler(request ziface.IRequest) {
	//从request中找到messageId
	messageId := request.GetMessageId()
	handler, ok := m.messageApi[messageId]
	if !ok {
		fmt.Printf("messageId [%d] don't have handler\n", messageId)
		return
	}
	//如果有对应的消息处理路由，则处理
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}

func (m *MessageHandler) AddRouter(messageId uint32, router ziface.IRouter) {
	//判断当前messageId是否有绑定的router
	if _, ok := m.messageApi[messageId]; ok {
		log.Printf("messageId [%d] already has router\n", messageId)
		return
	}
	//添加messageId对应的router映射关系
	m.messageApi[messageId] = router
	log.Printf("messageId [%d] add router success\n", messageId)
}
