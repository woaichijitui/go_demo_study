package _case

import (
	"io"
	"log"
	"os"
	"path"
)

func CopyDirtoDir() {
	//	获取源文件文件名字的切片
	list := getFiles(sourceDir)

	for _, f := range list {
		//Split函数将路径从最后一个斜杠后面位置分隔为两个部分（dir和file）并返回。如果路径中没有斜杠，函数返回值dir会设为空字符串，file会设为path。两个返回值满足path == dir+file
		_, name := path.Split(f)
		destFileName := destDir + "copy/" + name
		CopyFile(f, destFileName)
	}
}
func CopyFile(srcName, destName string) (int64, error) {
	//	打开源文件IO流
	src, err := os.Open(srcName)
	if err != nil {
		log.Fatal(err)
	}
	// 	用后关闭io流
	defer src.Close()

	dest, err := os.OpenFile(destName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer dest.Close()

	return io.Copy(dest, src)
}
