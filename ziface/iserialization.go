package ziface

//数据序列化接口 面向TCP连接，解决TCP传输过程中的粘包问题
type Serialization interface {
	//获取数据包长度的方法
	GetDataLength() int
	//序列化数据的方法(封装数据包的方法)
	Pack(message IMessage) ([]byte, error)
	//反序列化的方法(拆包的方法)
	UnPack([]byte) (IMessage, error)
}
