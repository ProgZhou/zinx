package ziface

//将请求的消息定义成一个抽象的消息模块
type IMessage interface {
	SetData(data []byte)
	GetData() []byte

	GetLen() uint32
	SetLen(length uint32)

	GetId() uint32
	SetId(id uint32)
}
