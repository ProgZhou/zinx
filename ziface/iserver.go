package ziface

//抽象层：定义一个服务器接口
type IServer interface {
	//启动服务器
	Start()
	//运行服务器
	Serve()
	//停止服务器
	Stop()
	//路由功能：给当前的服务注册一个路由方法，供客户端的连接处理使用
	AddRouter(messageId uint32, router IRouter)
	//获取当前服务器的连接管理
	GetConnManager() IConnManager
	//设置连接建立时的钩子函数
	SetConnectionStart(hookFunc func(conn IConnection))
	//设置连接释放时的钩子函数
	SetConnectionStop(hookFunc func(conn IConnection))
	//调用连接建立时的钩子函数
	CallConnStart(conn IConnection)
	//调用连接释放时的钩子函数
	CallConnStop(conn IConnection)
}
