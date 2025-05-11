package _case

import (
	"fmt"
	"io"
	"log"
	"os"
)

func FileCase() {
	//writeFile("F:\\系统\\go基础学习\\IO\\test\\test.txt")
	//copyFile("F:\\系统\\go基础学习\\IO\\test\\test.txt", "F:\\系统\\go基础学习\\IO\\test\\test2.txt")
	createfile("test\\create_file")
}

// 创建文件
//
//	创建的路径是项目的根路径加文件名
func createfile(name string) {

	newfile, _ := os.Create(name)
	defer newfile.Close()

}

// 一次性写入
func writeFile(path string) {
	err := os.WriteFile(path, []byte("这是一次写入测试"), os.ModePerm)
	fmt.Println(err)
}

// 文件复制
func copyFile(src string, dest string) {
	file, _ := os.Open(src)
	write, _ := os.OpenFile(dest, os.O_RDWR, 0655)
	written, err := io.Copy(write, file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(written)
}
