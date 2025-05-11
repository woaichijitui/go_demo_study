package main

import (
	"fmt"
	"net/rpc"
)

type Request struct {
	ID int
}

type Response struct {
	Name string
}

func main() {

	req := Request{

		ID: 12,
	}

	client, err := rpc.DialHTTP("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
	}

	var res Response
	s := client.Call("Hello.SayHello", req, &res)
	fmt.Println(res, s)
}
