package znet

import (
	"fmt"
	"log"
	"math/rand"
	"time"
	"zinx/utils"
	"zinx/ziface"
)

//消息管理模块具体实现层
type MessageHandler struct {

	//消息id与路由的映射
	messageApi map[uint32]ziface.IRouter

	//消息任务队列
	taskQueue []chan ziface.IRequest
	//messageHandler协程池的大小
	workerPoolSize uint32
}

//初始化MessageHandler的方法
func NewMessageHandler() *MessageHandler {
	return &MessageHandler{
		messageApi:     make(map[uint32]ziface.IRouter),
		workerPoolSize: uint32(utils.GlobalProperty.WorkPoolSize), //从全局配置中获取
		taskQueue:      make([]chan ziface.IRequest, utils.GlobalProperty.WorkPoolSize),
	}
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

//启动一个worker协程池(只会触发一次)
func (m *MessageHandler) StartWorkPool() {
	//根据工作池的大小分别开启worker，每个worker用一个go协程承载
	for i := 0; i < int(m.workerPoolSize); i++ {
		//给当前worker对应的channel开辟空间
		m.taskQueue[i] = make(chan ziface.IRequest, utils.GlobalProperty.MaxWorkTaskSize)
		//启动当前的worker
		go m.startWorker(i, m.taskQueue[i])
	}
}

//启动一个worker流程
func (m *MessageHandler) startWorker(id int, queue chan ziface.IRequest) {
	log.Printf("worker[%d] start...\n", id)
	//不断阻塞等待对应消息队列的消息
	for {
		select {
		//如果有消息过来，则处理当前request所绑定的业务
		case request := <-queue:
			m.DoMessageHandler(request)
		}
	}
}

//将request发送给对应的消息队列，由worker进行处理
func (m *MessageHandler) SendMessageToTaskQueue(request ziface.IRequest) {
	//轮询，将消息平均分配给taskQueue 暂时使用随机数
	rand.Seed(time.Now().UnixNano())
	workerId := rand.Intn(int(m.workerPoolSize)) % int(m.workerPoolSize)
	log.Printf("add request[%d] to worker[%d]\n", request.GetMessageId(), workerId)
	m.taskQueue[workerId] <- request
}
