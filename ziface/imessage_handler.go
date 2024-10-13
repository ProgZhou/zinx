package ziface

//消息管理抽象层
type IMessageHandler interface {

	//调度消息处理方法
	DoMessageHandler(request IRequest)
	//为消息添加处理路由
	AddRouter(messageId uint32, router IRouter)
	//启动worker工作池
	StartWorkPool()
	//将request发送给对应的消息队列，由worker进行处理
	SendMessageToTaskQueue(request IRequest)
}
