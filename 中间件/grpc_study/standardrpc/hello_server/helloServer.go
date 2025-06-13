package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

type Hello struct {
}

type Res struct {
	Name string
}

type Req struct {
	ID int
}

func (h Hello) SayHello(r Req, res *Res) error {
	res.Name = "htt" + string(r.ID)

	fmt.Println(r)
	return nil
}

func main() {
	//注册rpc服务
	rpc.Register(new(Hello))
	//调用http处理程序
	rpc.HandleHTTP()

	//监听
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	http.Serve(listen, nil)
}
