package _case

import (
	"log"
	"os"
	"strings"
)

// 源文件和目标目录，相对于这个项目来说
const (
	sourceDir = "source-file/"
	destDir   = "dest-file/"
)

// 获取一个文件夹里所有文件的路径
func getFiles(dir string) []string {
	//	读取文件
	fs, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	list := make([]string, 0)
	for _, f := range fs {
		if f.IsDir() {
			// 若是是文件夹 就跳过
			continue
		}

		//	获得文件路径
		fullname := strings.Trim(dir, "/") + "/" + f.Name()
		list = append(list, fullname)
	}
	return list
}
