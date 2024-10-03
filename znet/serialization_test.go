package znet

import (
	"fmt"
	"testing"
)

//单元测试
func TestSerialization(t *testing.T) {
	msg1 := &Message{
		Id:         0,
		DataLength: 5,
		Data:       []byte{'H', 'e', 'l', 'l', 'o'},
	}

	dp := NewSerialization()
	dataPack, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(dataPack)

	message, err := dp.UnPack(dataPack)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(message)
}
