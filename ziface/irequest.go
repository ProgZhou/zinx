package ziface

//IRequest接口：将客户端请求的连接信息和请求数据包装到一个Request中
type IRequest interface {
	//得到当前连接
	GetConnection() IConnection
	//得到请求的消息数据
	GetData() []byte
	//得到请求消息的id
	GetMessageId() uint32
}
