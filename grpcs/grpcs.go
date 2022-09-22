package grpcs

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/MoyunRz/grpc_impl/grpcs/gutils"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"time"
)

var SGCf *GrpcConfig
var CGCf *GrpcConfig

type GrpcConfig struct {
	Host     string
	Port     string
	RService []interface{}
}

type Data struct {
	Body interface{} `json:"body"`
}

func ConnClientConfig(host, port string) *GrpcConfig {
	if CGCf == nil {
		CGCf = &GrpcConfig{
			Host: host,
			Port: port,
		}
	}
	return CGCf
}

func ConnServiceConfig(host, port string) *GrpcConfig {
	if SGCf == nil {
		SGCf = &GrpcConfig{
			Host: host,
			Port: port,
		}
	}
	return SGCf
}
func (g *GrpcConfig) RegisterService(v interface{}) *GrpcConfig {
	g.RService = append(g.RService, v)
	return g
}

func (g *GrpcConfig) StartServer() {

	log.Printf("grpc server start run:%s/%s", g.Host, g.Port)
	// 监听本地的8972端口
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", g.Host, g.Port))
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}

	// 创建gRPC服务器,需要验证
	//s := grpc.NewServer(
	//	grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
	//		grpc_ctxtags.StreamServerInterceptor(),
	//		grpc_opentracing.StreamServerInterceptor(),
	//		grpc_zap.StreamServerInterceptor(grpc_mw2.ZapInterceptor()),
	//		grpc_auth.StreamServerInterceptor(grpc_mw2.AuthInterceptor),
	//		grpc_recovery.StreamServerInterceptor(grpc_mw2.RecoveryInterceptor()),
	//	)),
	//	grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
	//		grpc_ctxtags.UnaryServerInterceptor(),
	//		grpc_opentracing.UnaryServerInterceptor(),
	//		grpc_zap.UnaryServerInterceptor(grpc_mw2.ZapInterceptor()),
	//		grpc_auth.UnaryServerInterceptor(grpc_mw2.AuthInterceptor),
	//		grpc_recovery.UnaryServerInterceptor(grpc_mw2.RecoveryInterceptor()),
	//	)),
	//)
	s := grpc.NewServer()
	// 在gRPC服务端注册服务
	RegisterGreeterServer(s, &GrpcService{})
	//在给定的gRPC服务器上注册服务器反射服务
	reflection.Register(s)
	// Serve方法在lis上接受传入连接，为每个连接创建一个ServerTransport和server的goroutine。
	// 该goroutine读取gRPC请求，然后调用已注册的处理程序来响应它们。
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}

func (g *GrpcConfig) NewGrpcClient(method string, v interface{}, reply interface{}) {
	// 连接服务器
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", g.Host, g.Port), grpc.WithInsecure())
	if err != nil {
		fmt.Printf("faild to connect: %v", err)
	}
	defer conn.Close()

	mp := Data{
		Body: v,
	}

	mb := map[string][]byte{}
	bj, err := json.Marshal(mp)

	// 序列化
	mb["key_data"] = bj
	// 调用服务端的SayHello
	timeStamp := time.Now().Unix()
	c := NewGreeterClient(conn)
	r, err := c.SendRequest(context.Background(), &ParamRequest{
		Method:    method,
		Params:    mb,
		TimeStamp: timeStamp,
	})
	if err != nil {
		log.Errorf("grpc post客户端出错: could not greet: %v", err)
		return
	}
	if r.Code == -1 {
		log.Errorf("服务端错误提示: %s ", r.Msg)
		return
	}
	// 反序列化
	data := r.RpcReply["data"]
	log.Errorf("客户端返回数据: %v ", string(data))
	if data != nil {
		err = gutils.Deserialize(data, reply)
		if err != nil {
			log.Errorf("grpc 客户端反序列化出错: %v ", err)
			return
		}
	}
	return
}
