package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc_study/grpc/grpc_bp/hello_grpc"
)

/*建立连接
调用方法*/

func main() {

	addr := ":8081"

	// 使用 grpc.Dial 创建一个到指定地址的 gRPC 连接。
	// 此处使用不安全的证书来实现 SSL/TLS 连接
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 初始化客户端
	client := hello_grpc.NewHelloClient(conn)
	res, err := client.Hello(context.Background(), &hello_grpc.HelloRequest{
		ID: 1,
	})

	fmt.Println(res, err)

}
