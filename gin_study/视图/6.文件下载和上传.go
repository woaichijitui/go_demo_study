package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
)

// 文件上传
func _fileDown(c *gin.Context) {

	//	获取 fileHearder
	fileHearder, _ := c.FormFile("file1")
	//	获取	file
	// 	file 实现了reader writer
	file, _ := fileHearder.Open()
	//	保存文件
	//	1.正常保存 字节缓存
	//	2.SaveUploadedFile
	all, _ := io.ReadAll(file)
	fmt.Println(len(all) / 1024)

	dst := "./uploads/" + fileHearder.Filename
	c.SaveUploadedFile(fileHearder, dst)

}

////文件下载

func _FileUpload(c *gin.Context) {
	c.File("uploads/golang.jpg")
}
func main() {

	router := gin.Default()

	//文件上传
	router.POST("/fileDown", _fileDown)

	//文件下载
	router.POST("/fileUpload", _FileUpload)
	router.Run(":8080")
}
