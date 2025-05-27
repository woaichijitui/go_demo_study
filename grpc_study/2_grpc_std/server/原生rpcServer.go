package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

type HelloService struct{}

type Req struct {
	Id int
}

type Res struct {
	Name string
}

func (h HelloService) SayHello(r Req, res *Res) error {
	res.Name = "hello"
	fmt.Println("rpc 请求")
	return nil
}

func main() {
	rpc.Register(new(HelloService))
	rpc.HandleHTTP()

	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	http.Serve(listen, nil)
}
