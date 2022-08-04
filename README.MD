# grpc_impl

## 介绍

该项目是为了golang项目中为了能更好的进行grpc的使用，去掉繁琐proto重零开始搭建的繁琐事情，
直接专注于service服务的开发而写的一个grpc模块工具

## 源码测试

 请看 `example`目录下的 `grpcs_test.go`


## 使用教程-Demo

```go

package tests

// 导入框架包
import (
	"testing"
	"fmt"
	"github.com/MoyunRz/grpc_impl/grpcs"
)

// TestServer
// 启动服务端测试
func TestServer(t *testing.T) {

    // 注册服务的结构体
    // 结构体含有调用的服务方法
	var rservice = gService.PubService{}
	// ....
	grpcs.ConnServiceConfig("127.0.0.1", "8087").
	RegisterService(&rservice).    // 注册服务
	StartServer()   // 启动GRPC服务端
}

// TestClient
// 启动客户端测试
func TestClient(t *testing.T) {
    // 发生数据
	kd := map[string]interface{}{
		"user_id":    1,
		"device_id":  "",
		"quantity":   1,
		"integral":   "0",
		"calorie":    "0",
		"remark":     "",
		"start_time": "2022-08-04 18:40:20:08:00",
	}
	// 回调数据
	kds := ""
	grpcs.
	ConnClientConfig("127.0.0.1", "8087").  // 传入服务端的地址、端口
	NewGrpcClient("FindStepsByUserId", kd, &kds) // 调用服务端的方法
	fmt.Println(kds)
}

```

## 服务方法和结构体的创建-Demo

```go
package gService

// PubService
// 该结构体的创建是为了进行注册时可以扫描到该结构体下的方法函数
type PubService struct {}

// SendMessage
// 发送信息
// receiver *grpcs.DealService 该参数每个函数都要带上，该参数用于接收调用方传入的数据
func (p PubService) SendMessage(receiver *grpcs.DealService) {
    // 你自己定义的接收结构体
	res := domain.DeviceInfo{}
	// 将数据绑定到你的结构体中
	err := receiver.JsonBind(&res)
	log.Println("已经收到消息")
	if err != nil {
		receiver.ErrResult(err.Error())
		return
	}
	
	// 处理逻辑
	// ...
    // 返回数据，该函数是封装好数据返回给调用的客户端
	receiver.Result("success")
}

```

## 服务端启动注册

以上面的结构体 `PubService` 为列子进行注册
```
    // 注册服务的结构体
    // 结构体含有调用的服务方法
	var rservice = PubService{}
	grpcs.ConnServiceConfig("127.0.0.1", "8087"). // 服务端地址、端口号
	RegisterService(&rservice).     // 注册服务
	StartServer()                   // 启动GRPC服务端
	
```

## 客户调用参数

以上面的结构体 `PubService`下的 `SendMessage(receiver *grpcs.DealService)` 为列子进行远程调用

```
    // 要传给服务端的数据信息
    kd := map[string]interface{}{
		"msg":    "hello",
	}
    // 定义接收的数据字段
    // 现在是string，因为那边服务端是返回'success' --string类型
    // 你也可以自行定义struct，但要和服务端协定好
    kds := ""
    // 开始调用
	grpcs.
	ConnClientConfig("127.0.0.1", "8087").  // 传入服务端的地址、端口
	NewGrpcClient("SendMessage", kd, &kds) // 调用服务端的方法
	fmt.Println(kds)
```


