package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"zinx/ziface"
)

//全局配置模块，用于读取用户编写的json配置文件
type GlobalProperties struct {
	TcpServer ziface.IServer //全局server对象
	Host      string         `json:"host"` //服务器主机监听的ip地址
	Port      int            `json:"port"` //服务器端口号
	Name      string         `json:"name"` //服务器名称

	Version   string `json:"version"`   //当前服务器版本号
	MaxBuffer int    `json:"maxBuffer"` //当前服务器能够读取的缓冲区大小
}

//定义一个全局的对外对象
var GlobalProperty *GlobalProperties

//加载配置文件的方法
func (g *GlobalProperties) Load() {
	dir, _ := os.Getwd()
	data, err := os.ReadFile(fmt.Sprintf("%s/utils/conf/zinx.json", dir))
	if err != nil {
		log.Fatalln("properties read error: " + err.Error())
	}
	err = json.Unmarshal(data, GlobalProperty)
	if err != nil {
		log.Fatalln("json unmarshal error: " + err.Error())
	}
}

//提供一个初始化方法
func init() {
	//默认配置
	GlobalProperty = &GlobalProperties{
		Host:      "127.0.0.1",
		Port:      8888,
		Name:      "Zinx-Server",
		Version:   "v0.4",
		MaxBuffer: 512,
	}
	GlobalProperty.Load()
}
