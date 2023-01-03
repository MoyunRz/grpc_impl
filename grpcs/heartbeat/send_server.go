package heartbeat

import (
	"fmt"
	"github.com/MoyunRz/grpc_impl/grpcs"
	"time"
)

// SendService
// 该结构体的创建是为了进行注册时可以扫描到该结构体下的方法函数
type SendService struct{}

// SendHeartbeat
// 发送心跳
// receiver *grpcs.DealService 该参数每个函数都要带上，该参数用于接收调用方传入的数据
func sendHeartbeat(host, port string) bool {
	type Heartbeat struct {
		Host      string `json:"host"`
		TimeStamp int64  `json:"timestamp"`
	}
	var rhd = Heartbeat{
		Host:      PubLocalHostInfo.IpAddr,
		TimeStamp: time.Now().Unix(),
	}
	var res = ""
	grpcs.ConnClientConfig(host, port).NewGrpcClient("ReceiveHeartbeat", rhd, &res)
	if res != "SUCCESS" {
		fmt.Println("心跳发送失败")
		return false
	} else {
		fmt.Println("心跳发送 SUCCESS ")
		return true
	}
}
