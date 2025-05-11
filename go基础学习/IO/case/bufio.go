package _case

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func BufioCase() {
	//bufioReadCase("F:\\系统\\coding\\go基础学习\\IO\\case\\read.go")
	scannerRead("F:\\系统\\coding\\go基础学习\\IO\\case\\read.go")
	//bufioWriteCase("F:\\系统\\coding\go基础学习\\IO\\test\\bufiotest.txt")
	//bufReadAndWrite("F:\\系统\\coding\\go基础学习\\IO\\test\\test.txt", "F:\\系统\\coding\\go基础学习\\IO\\test\\bufiotest.txt")
}

// bufio 带缓冲的读
// 读一行
func bufioReadCase(path string) {

	file, _ := os.Open(path)
	bufRead := bufio.NewReader(file)
	defer file.Close()
	for {
		line, _, err := bufRead.ReadLine()
		fmt.Println(string(line))
		if err != nil {
			break
		}
	}
}

// Scanner 指定分隔符读
func scannerRead(path string) {
	file, _ := os.Open(path)
	defer file.Close()

	/*func NewScanner(r io.Reader) *Scanner {
		return &Scanner{
		r:            r,
		split:        ScanLines,
		maxTokenSize: MaxScanTokenSize,
	}
	}*/
	scanner := bufio.NewScanner(file)
	//	按行读(默认按行读
	//scanner.Split(bufio.ScanLines)
	//	按word读
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// bufio 带缓冲的写
// WriteString 是从头开始覆盖的写 但并不清空文件
func bufioWriteCase(path string) {

	file, _ := os.OpenFile(path, os.O_RDWR, 0655)
	bufWrite := bufio.NewWriter(file)
	defer file.Close()
	for i := 0; i < 10; i++ {
		n, err := bufWrite.WriteString("hello\n")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("写了%d个字符\n", n)
	}
	bufWrite.Flush()

}

// 文件读写案列
func bufReadAndWrite(src string, desk string) {

	//	bufio，Newread
	file, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	bufRead := bufio.NewReader(file)

	//
	deskFile, err := os.OpenFile(desk, os.O_RDWR, 0655)
	if err != nil {
		log.Fatal(err)
	}
	defer deskFile.Close()
	bufWriter := bufio.NewWriter(deskFile)
	//all := make([]byte, 0)
	//	读写文件
	//	无需缓冲
	for {
		//ReadLine 低水平
		//line, prefix, err := bufRead.ReadLine()
		//	改用 ReadBytes delim:遇到这个字节时返回切片
		bytes, err := bufRead.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("读完了")
				break
			}
			log.Fatal(err)
		}
		//all = append(all, bytes...)
		nn, err := bufWriter.WriteString(string(bytes))

		fmt.Printf("读取了%d个字字符\n", nn)

	}
	//将滞留在内部缓冲区的数据写入文件
	//	例如 此案例中 若末尾不写这个方法 会导致最后一行数据不能写入文件 留在缓冲区里
	bufWriter.Flush()

}
