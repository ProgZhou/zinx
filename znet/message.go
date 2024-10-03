package znet

import "zinx/ziface"

type Message struct {
	ziface.IMessage
	Id         uint32 //消息id
	DataLength uint32 //消息长度
	Data       []byte //消息具体内容
}

func NewMessage(messageId uint32, data []byte) *Message {
	return &Message{
		Id:         messageId,
		DataLength: uint32(len(data)),
		Data:       data,
	}
}

func (m *Message) SetData(data []byte) {
	m.Data = data
}

func (m *Message) GetData() []byte {
	return m.Data
}

func (m *Message) GetLen() uint32 {
	return m.DataLength
}

func (m *Message) SetLen(length uint32) {
	m.DataLength = length
}

func (m *Message) GetId() uint32 {
	return m.Id
}

func (m *Message) SetId(id uint32) {
	m.Id = id
}
