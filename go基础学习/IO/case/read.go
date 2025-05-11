package _case

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func Read() {
	stringReadCase("hello world!!")
	OSReadCase("F:\\系统\\coding\\go基础学习\\IO\\case\\read.go")

	seekRead("F:\\系统\\coding\\go基础学习\\IO\\test\\test.txt")
}

// strings.NewRead 实现了reader
func stringReadCase(str string) {
	reader := strings.NewReader(str)
	//	创建一个缓冲的切片、
	buf := make([]byte, 1024)
	//	循环读取数据
	for {
		//将数据读取到切片里 并返回切片的长度（数据的长度）和错误
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("读完了")
				break
			}
			log.Fatal(err)
		}
		fmt.Println(string(buf[:n]))
	}

}

// os.Open 实现了 reader
func OSReadCase(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//	创建一个缓冲的切片、
	buf := make([]byte, 1024)
	//创建一个切片储存数据
	r := make([]byte, 0)
	//	循环读取数据
	for {
		//将数据读取到切片里 并返回切片的长度（数据的长度）和错误
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("读完了")
				break
			}
			log.Fatal(err)
		}
		//buf...将切片给的元素追加r后面
		//	但是会出现错误
		r = append(r, buf[:n]...)

		//这种性能怎么样?
		//for i := 0; i < n; i++ {
		//	r = append(r, buf[i])
		//}
	}
	fmt.Println(string(r))

}

// seek 用于指定下次读取或者写入时的偏移量
func seekRead(path string) {
	file, _ := os.Open(path)
	defer file.Close()

	//SeekCurrent offset相对位置 坐标起点
	file.Seek(0, io.SeekEnd)
	buf := make([]byte, 1024)
	_, _ = file.Read(buf)
	fmt.Println(string(buf))
}
