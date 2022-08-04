package grpcs

import (
	"context"
	"encoding/json"
	"log"
	"reflect"
)

type GrpcService struct{}

func (g *GrpcService) SendRequest(ctx context.Context, params *ParamRequest) (*ParamReply, error) {
	values := CallMethodByName(params)
	if values != nil {
		v := values.(*ParamReply)
		return v, nil
	}

	return &ParamReply{
		Code: 1,
		Msg:  "success",
	}, nil

}

type Response struct {
	Body interface{} `json:"body"`
}

func CallMethodByName(param *ParamRequest) interface{} {
	// 要获取的结构体
	// 获取结构体内部信息
	for _, item := range SGCf.RService {
		mtV := reflect.ValueOf(item).Elem()
		// 根据名称反射获取结构体内的方法
		call := mtV.MethodByName(param.Method)
		// 判断方法是否存在
		if call.IsNil() || call.Kind() != reflect.Func {
			continue
		}
		// 传入一个参数
		params := make([]reflect.Value, 0)
		var rp Response
		json.Unmarshal(param.Params["key_data"], &rp)
		ds := DealService{}
		ds.Data = rp.Body
		params = append(params, reflect.ValueOf(ds))
		// 调用该方法
		call.Call(params)
		// 返回方法的处理结果
		return ds.Reply
	}
	log.Print("不存在")
	return nil
}
