package example

import (
	"github.com/MoyunRz/grpc_impl/grpcs"
	"log"
)

// PubService
// 该结构体的创建是为了进行注册时可以扫描到该结构体下的方法函数
type PubService struct{}

// SendMessage
// 发送信息
// receiver *grpcs.DealService 该参数每个函数都要带上，该参数用于接收调用方传入的数据
func (p PubService) SendMessage(receiver *grpcs.DealService) {
	type KeyData struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	// 你自己定义的接收结构体
	var kd KeyData
	// 将数据绑定到你的结构体中
	err := receiver.JsonBind(&kd)
	log.Println("已经收到消息")
	if err != nil {
		receiver.ErrResult(err.Error())
		return
	}
	receiver.Result("success")
}
