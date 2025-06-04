package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_study/grpc/grpc_bp/proto"
	"net"
)

type DuoSayService struct {
}

//实现

func (h DuoSayService) Say(c context.Context, req *proto.HelloRequest) (*proto.HellResponse, error) {
	fmt.Println("客户端消息：", req)
	return &proto.HellResponse{Name: "htt"}, nil
}

type DuoLookService struct {
}

//实现

func (h DuoLookService) Look(c context.Context, req *proto.HelloRequest) (*proto.HellResponse, error) {
	fmt.Println("客户端消息：", req)
	return &proto.HellResponse{Name: "htt"}, nil
}

/*
编写一个结构体，名字叫什么不重要
重要的是得实现protobuf中的所有方法
监听端口
注册服务
*/
func main() {
	//监听端口
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println(err)
	}
	// 创建一个gRPC服务器实例。
	s := grpc.NewServer()
	service1 := DuoSayService{}
	service2 := DuoLookService{}
	//	注册服务
	proto.RegisterSayServer(s, &service1)
	proto.RegisterLookServer(s, &service2)
	fmt.Println("grpc run time at :8081")

	//处理客户端请求
	s.Serve(listen)
}
