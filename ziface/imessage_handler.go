package ziface

//消息管理抽象层
type IMessageHandler interface {

	//调度消息处理方法
	DoMessageHandler(request IRequest)
	//为消息添加处理路由
	AddRouter(messageId uint32, router IRouter)
}
