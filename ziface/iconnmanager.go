package ziface

//连接管理模块 抽象层
type IConnManager interface {
	//添加连接
	AddConnection(conn IConnection)
	//删除连接
	RemoveConnection(conn IConnection)
	//根据connId查询连接
	GetConnection(connId uint32) (IConnection, error)
	//获取当前连接总数
	Size() int
	//清楚并终止所有连接
	ClearConnection()
}
