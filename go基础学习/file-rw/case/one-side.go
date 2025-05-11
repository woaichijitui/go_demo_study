package _case

import (
	"io"
	"log"
	"os"
	"path"
)

// 文件复制
func OneSideReadWriteToDest() {
	files := getFiles(sourceDir)
	for _, f := range files {
		_, name := path.Split(f)
		destName := destDir + "one-side/" + name
		oneSideReadWrite(destName, f)
	}

}
func oneSideReadWrite(destName, src string) {
	srcfile, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer srcfile.Close()

	destFile, err := os.OpenFile(destName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer destFile.Close()

	// 用于读取时的缓存
	buf := make([]byte, 1024)

	for {
		// 利用缓冲读取
		n, err := srcfile.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		//	当n == 0 ，读取结束
		if n == 0 {
			break
		}

		// 其实buf就是搬运工，在源文件和目标文件之间搬运数据
		destFile.Write(buf[:n])

	}
}
