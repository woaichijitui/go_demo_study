package _case

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const README = "README.md"

// 一次性读取
// 按行打印
// 适合小文件
func ReadLine1() {
	readme, err := os.Open(README)
	if err != nil {
		log.Fatal(err)
	}
	defer readme.Close()

	//	io.readAlL 方法读取全部内容
	bytes, err := io.ReadAll(readme)
	if err != nil {
		log.Fatal(err)
	}

	//	strings.spilt 分割内容
	//fmt.Println(string(bytes))
	lines := strings.Split(string(bytes), "\n")

	for _, line := range lines {
		fmt.Print(line)
	}
}

// 通过bufio 按行读取
// bufio 通过对io模块的封装，提供了数据缓冲的功能，能一定程度上减少大数据读写带来的开销
// 当发起读写操作时，会尝试从缓冲区读取数据，缓冲区没有数据后，才会从数据源获取
// 默认大小为4k
func ReadLine2() {
	fileHandle, err := os.OpenFile(README, os.O_RDONLY, 0444)
	if err != nil {
		log.Fatal(err)
	}
	defer fileHandle.Close()

	// 	bufio.NewReader 带缓冲的io流
	reader := bufio.NewReader(fileHandle)
	for {
		//	 reader.ReadString 切片操作
		//	byte 可以用 单引号
		//	每行自带换行符
		readString, err := reader.ReadString(byte('\n'))

		if err == io.EOF {
			break
		}
		fmt.Print(readString)
	}
}

// 通过scanner 按行读取
// 默认缓冲大小64k
func ReadLine3() {
	openFile, err := os.OpenFile(README, os.O_RDONLY, 0444)
	if err != nil {
		log.Fatal(err)
	}
	defer openFile.Close()
	//	bufio.NewScanner 获取
	scanner := bufio.NewScanner(openFile)

	// 这是什么我忘了？
	for scanner.Scan() {
		//	scanner.Text() 方法来获取每一行
		line := scanner.Text()

		fmt.Println(line)
	}

}
