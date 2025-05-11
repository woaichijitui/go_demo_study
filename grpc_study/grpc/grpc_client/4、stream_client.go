package main

import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc_study/grpc/grpc_bp/proto"
	"io"
	"log"
	os "os"
)

func main() {

	conn, err := grpc.NewClient(":8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	//	初始化
	streamServiceClient := proto.NewStreamServiceClient(conn)
	stream, err := streamServiceClient.DownLoadFile(context.Background(), &proto.Request{
		Name: "电影",
	})

	file, _ := os.OpenFile("电影.mp4", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	writer := bufio.NewWriter(file)
	for {
		response, err := stream.Recv()
		if err != nil && err == io.EOF {
			fmt.Println(err)
			//	客户端流结束
			break
		}

		_, err = writer.Write(response.GetContent())
		if err != nil {
			log.Fatal(err)
		}

	}
	writer.Flush()

}
