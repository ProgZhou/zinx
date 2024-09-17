package znet

import "zinx/ziface"

//实现router先嵌入这个baseRouter基类，用户可自定义baseRouter 如果继承baseRouter类，即可随机选择其中的某个方法进行重写
type BaseRouter struct {
}

func (b *BaseRouter) PreHandle(request ziface.IRequest) {

}

func (b *BaseRouter) Handle(request ziface.IRequest) {

}

func (b *BaseRouter) PostHandle(request ziface.IRequest) {

}
