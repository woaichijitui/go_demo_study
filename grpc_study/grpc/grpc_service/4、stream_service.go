package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc_study/grpc/grpc_bp/proto"
	"io"
	"log"
	"net"
	"os"
)

type StreamServer struct {
}

func (s StreamServer) DownLoadFile(req *proto.Request, stream proto.StreamService_DownLoadFileServer) error {
	fmt.Println(req)

	//读取文件
	file, err := os.Open("D:\\study\\grpc\\【枫枫知道】手把手系列之gRPC基础教程  Go语言 P2 2.grpc工具安装.mp4")
	if err != nil {
		log.Fatal(err)
	}
	//分片读取
	buf := make([]byte, 1024)
	var count = 0
	for {
		count++
		_, err := file.Read(buf)
		if err != nil && err == io.EOF {
			break

		}
		stream.Send(&proto.Response{
			FileName: req.Name,
			Content:  buf,
		})

		fmt.Printf("读取了%d次文件", count)

	}

	return nil
}

func main() {
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	proto.RegisterStreamServiceServer(s, &StreamServer{})
	fmt.Println("grpc run time at :8081")

	//处理客户端请求
	s.Serve(listen)

}
