package main

import (
	"fmt"
	"net/rpc"
)

type Req struct {
	Id int
}

type Res struct {
	Name string
}

func main() {
	req := Req{
		Id: 1,
	}

	dial, err := rpc.DialHTTP("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	var res Res
	s := dial.Call("HelloService.SayHello", req, &res)
	fmt.Println(res, s)
}
