package _case

import (
	"bufio"
	"log"
	"net"
	"os"
	"strings"
)

func Client() {
	// 	发送请求
	conn, err := net.Dial("tcp", "127.0.0.1:2000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	log.Printf("发送请求\n")

	//	读取终端数据
	stdRead := bufio.NewReader(os.Stdin)
	//	从标准输入中 发送请求数据
	for {
		str, err := stdRead.ReadString('\n')
		if err != nil {
			log.Println(err)
			continue
		}

		str = strings.Trim(str, "\r\n")
		// 若是输入Q 退出
		if str == "Q" {
			break
		}
		//	发送请求数据
		//	接受响应数据

		_, err = conn.Write([]byte(str))
		if err != nil {
			log.Fatal(err)
		}

		buf := make([]byte, 152)
		//	接受响应
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("接收到服务端消息： %s \n", string(buf[:n]))

	}
}
