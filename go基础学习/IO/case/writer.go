package _case

import (
	"log"
	"os"
)

func WriterCase() {
	writer("F:\\系统\\go基础学习\\IO\\test\\test.txt")
}

// os.openfile 可以实现io.writer接口
func writer(path string) {

	//	以可读可写模式打开file
	file, err := os.OpenFile(path, os.O_RDWR, 655)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 和读取一样需要一个缓冲取

	// 写入
	file.Write([]byte("这是writer 写入测试"))

}
