package _case

import (
	"log"
	"net"
)

// 调用一个简单的tcp server 服务
func Server() {
	//
	listen, err := net.Listen("tcp", "127.0.0.1:2000")
	if err != nil {
		log.Fatal(err)
	}
	defer listen.Close()
	log.Println("开始监听")

	// 循环监听并处理数据
	for {
		//监听
		// conn接口 继承了reader writer 接口
		// 接受请求 阻塞等待 建立连接（接受请求之后会建立连接
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go process(conn)
	}

}

// 处理函数
func process(conn net.Conn) {
	// 避免资源浪费
	defer conn.Close()
	for {
		buf := make([]byte, 152)
		n, err := conn.Read(buf)
		if err != nil {
			log.Println(err)
			break
		}
		s := string(buf[:n])
		log.Printf("客户端接受到消息 %s\n", s)

		// 响应
		_, err = conn.Write(buf[:n])
		if err != nil {
			log.Println(err)
			break
		}
	}

}
