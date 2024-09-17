package ziface

//抽象路由接口：路由里的数据都是IRequest
type IRouter interface {

	//处理业务之前的方法(前处理)
	PreHandle(request IRequest)

	//处理业务的主方法
	Handle(request IRequest)

	//处理业务之后的方法(后处理)
	PostHandle(request IRequest)
}
