package example

import (
	"fmt"
	"github.com/MoyunRz/grpc_impl/grpcs"
	"reflect"
	"testing"
)

type KeyData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
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
	var rservice = grpcs.DealService{}
	grpcs.ConnServiceConfig("127.0.0.1", "8087").RegisterService(&rservice).StartServer()
}

func TestClient(t *testing.T) {
	kd := KeyData{Name: "myPost", Age: 12}
	kds := ""
	grpcs.ConnClientConfig("127.0.0.1", "8087").NewGrpcClient("SendMessage", kd, &kds)
	fmt.Println(kds)
}
