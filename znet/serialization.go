package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"zinx/utils"
	"zinx/ziface"
)

type Serialization struct {
}

func NewSerialization() *Serialization {
	return &Serialization{}
}

//获取数据包长度的方法
func (s *Serialization) GetDataLength() int {
	//消息头均为8个字节 长度4byte 消息id 4byte
	return 8
}

//序列化数据的方法(封装数据包的方法)
func (s *Serialization) Pack(message ziface.IMessage) ([]byte, error) {
	//创建一个存放字节数据的buffer
	buffer := bytes.NewBuffer([]byte{})
	//将data长度写入buffer中
	//使用二进制写入，小端方式(与拆包时对应即可)
	err := binary.Write(buffer, binary.LittleEndian, message.GetLen())
	if err != nil {
		return nil, err
	}
	//将messageId写入buffer中
	err = binary.Write(buffer, binary.LittleEndian, message.GetId())
	if err != nil {
		return nil, err
	}
	//将data具体数据写入buffer中
	err = binary.Write(buffer, binary.LittleEndian, message.GetData())
	if err != nil {
		return nil, err
	}
	//返回封装好的数据包
	return buffer.Bytes(), nil
}

//反序列化的方法(拆包的方法)
func (s *Serialization) UnPack(binaryData []byte) (ziface.IMessage, error) {
	//从输入的数据中创建一个ioReader
	reader := bytes.NewReader(binaryData)
	message := &Message{}
	//读取消息头长度
	err := binary.Read(reader, binary.LittleEndian, &message.DataLength)
	if err != nil {
		return nil, err
	}
	//读取messageId
	err = binary.Read(reader, binary.LittleEndian, &message.Id)
	if err != nil {
		return nil, err
	}
	//判断数据包长度是否超过配置的最大数据包大小
	if utils.GlobalProperty.MaxPackageSize > 0 && message.DataLength > uint32(utils.GlobalProperty.MaxPackageSize) {
		return nil, errors.New("data package too large")
	}
	return message, nil
}
