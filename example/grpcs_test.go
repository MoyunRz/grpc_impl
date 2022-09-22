package example

import (
	"fmt"
	"github.com/MoyunRz/grpc_impl/grpcs"
	"github.com/shopspring/decimal"
	"reflect"
	"testing"
)

type KeyData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type HeartAndBloodOxygenInfo struct {
	UserId         int64           `json:"user_id"`
	DeviceId       string          `json:"device_id"`
	HeartFrequency decimal.Decimal `json:"heart_frequency"`
	Hblood         decimal.Decimal `json:"hblood"`
	Lblood         decimal.Decimal `json:"lblood"`
	SpO2           decimal.Decimal `json:"spO2"`
	Remark         string          `json:"remark"`
	StartTime      string          `json:"start_time"`
}

func TestRpc1(t *testing.T) {
	myType := &grpcs.DealService{}
	mtV := reflect.ValueOf(&myType).Elem()
	call := mtV.MethodByName("StringService")
	if call.Kind() == reflect.Func {
		fmt.Println("true")
	}
	methods := call.Call(nil)
	str := methods[0].Interface().(string)
	fmt.Println(str)
	err := methods[1].Interface()
	if err != nil {
		fmt.Println(err.(error))
	}
}

func TestServer(t *testing.T) {
	var rservice = PubService{}
	grpcs.ConnServiceConfig("127.0.0.1", "8087").RegisterService(&rservice).StartServer()
}

func TestClient(t *testing.T) {
	kd := KeyData{Name: "myPost", Age: 12}
	kds := ""
	grpcs.ConnClientConfig("127.0.0.1", "8087").NewGrpcClient("SendMessage", kd, &kds)
	fmt.Println(kds)
}
