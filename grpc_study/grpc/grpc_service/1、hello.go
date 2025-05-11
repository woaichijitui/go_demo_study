package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_study/grpc/grpc_bp/hello_grpc"
	"net"
)

type HelloService struct {
}

//实现

func (h HelloService) Hello(c context.Context, req *hello_grpc.HelloRequest) (*hello_grpc.HellResponse, error) {
	fmt.Println("客户端消息：", req)
	return &hello_grpc.HellResponse{Name: "htt"}, nil
}

/*
编写一个结构体，名字叫什么不重要
重要的是得实现protobuf中的所有方法
监听端口
注册服务
*/
func main() {
	//监听端口
	listen, _ := net.Listen("tcp", ":8081")

	// 创建一个gRPC服务器实例。
	server := HelloService{}
	s := grpc.NewServer()

	//	注册服务
	hello_grpc.RegisterHelloServer(s, &server)
	fmt.Println("grpc run time at :8081")

	//处理客户端请求
	s.Serve(listen)
}
