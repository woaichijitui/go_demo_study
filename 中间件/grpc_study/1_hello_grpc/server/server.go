package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc_study/1_hello_grpc/hello_grpc"
	"net"
)

type HelloServer struct {
}

func (receiver HelloServer) SayHello(ctx context.Context, in *hello_grpc.HelloRequest) (*hello_grpc.HelloResponse, error) {
	if in.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name is empty")
	}
	return &hello_grpc.HelloResponse{
		Message: "Hello " + in.Name,
	}, nil
}

func main() {

	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}
	s := HelloServer{}

	server := grpc.NewServer()

	hello_grpc.RegisterHelloServiceServer(server, s)

	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}
}
