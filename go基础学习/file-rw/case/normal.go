package _case

import (
	"log"
	"os"
	"path"
)

//一般的读取文件
//os包提供的

func ReadWriteFiles() {
	//	读取文件列表
	files := getFiles(sourceDir)
	//	循环读取文件并写入到目标文件夹中文件夹
	for _, f := range files {
		//	读取操作
		//	一次性读取文件内容并写入新文件
		bytes, err := os.ReadFile(f)
		if err != nil {
			log.Fatal(err)
		}
		_, name := path.Split(f)
		//	目标文件夹路径 + 文件名字 = 文件路径
		destName := destDir + "normal/" + name
		//	写入操作
		err = os.WriteFile(destName, bytes, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}
