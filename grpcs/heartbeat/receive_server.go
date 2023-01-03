package heartbeat

import (
	"log"

	"github.com/MoyunRz/grpc_impl/grpcs"
)

// PubService
// 该结构体的创建是为了进行注册时可以扫描到该结构体下的方法函数
type ReceiveService struct{}

// ReceiveData
// 接收信息
// receiver *grpcs.DealService 该参数每个函数都要带上，该参数用于接收调用方传入的数据
func (p ReceiveService) ReceiveData(receiver *grpcs.DealService) {
	// 你自己定义的接收结构体
	var rData ReceiveHeartbeatData
	// 将数据绑定到你的结构体中
	err := receiver.JsonBind(&rData)
	log.Println("已经收到消息")
	if err != nil {
		receiver.ErrResult(err.Error())
		return
	}
	receiver.Result("success")
}

// ReceiveHeartbeat
// 接收心跳
// receiver *grpcs.DealService 该参数每个函数都要带上，该参数用于接收调用方传入的数据
func (p ReceiveService) ReceiveHeartbeat(receiver *grpcs.DealService) {
	// 你自己定义的接收结构体
	type Heartbeat struct {
		Host      string `json:"host"`
		TimeStamp int64  `json:"timestamp"`
	}
	var hb Heartbeat
	// 将数据绑定到你的结构体中
	err := receiver.JsonBind(&hb)
	log.Printf("this host:[%s]  is well, time: %d  /n", hb.Host, hb.TimeStamp)

	if err != nil {
		receiver.ErrResult(err.Error())
		return
	}
	//if PubLocalHostInfo.IpAddr == "" || PubLocalHostInfo.Domain == "" {
	//	PubLocalHostInfo = GetHostInfo()
	//}
	receiver.Result("SUCCESS")
}
