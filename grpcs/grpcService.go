package grpcs

import (
	"encoding/json"
	"github.com/MoyunRz/grpc_impl/grpcs/gutils"
	"github.com/goinggo/mapstructure"
	log "github.com/sirupsen/logrus"
)

type DealService struct {
	Data  interface{} `json:"data"`
	Reply *ParamReply `json:"reply"`
}

type KeyData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (receiver *DealService) Result(value interface{}) {
	reply := &ParamReply{
		Code: 1,
		Msg:  "success",
	}
	send := map[string][]byte{}
	send["data"] = nil
	if value != nil {
		data, err := gutils.Serialize(value)
		if err != nil {
			reply.Code = -1
			reply.Msg = err.Error()
			receiver.Reply = reply
			return
		}
		send["data"] = data
	}
	reply.RpcReply = send
	receiver.Reply = reply
}

func (receiver *DealService) ErrResult(msg string) {
	send := map[string][]byte{}
	send["data"] = nil
	receiver.Reply = &ParamReply{
		Code:     -1,
		Msg:      msg,
		RpcReply: send,
	}
}

func (receiver *DealService) Bind(v interface{}) error {
	if receiver.Data != nil {
		data := receiver.Data.(map[string]interface{})
		if err := mapstructure.Decode(data, v); err != nil {
			log.Error(err.Error())
			return err
		}
	}
	return nil
}

func (receiver DealService) JsonBind(v interface{}) error {
	b, err := json.Marshal(receiver.Data)
	err = json.Unmarshal(b, v)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}
