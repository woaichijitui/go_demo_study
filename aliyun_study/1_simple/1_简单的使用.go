package _simple

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"path/filepath"
)

// OSS配置
const (
	Endpoint        = "" //地域
	AccessKeyId     = ""
	AccessKeySecret = ""
	BucketName      = ""
)

func UploadImage(c *gin.Context) {
	//接收文件
	form, err := c.FormFile("image")
	if err != nil {
		fmt.Println(err)
		return
	}
	file, err := form.Open()
	if err != nil {
		fmt.Println(err)
		return
	}

	//获取阿里云存储空间
	bucket := CreateAliyunClient()

	// 生成上传路径和文件名
	objectName := filepath.Join("test1", "images", form.Filename)

	// 上传文件到OSS
	err = bucket.PutObject(objectName, file)
	if err != nil {
		fmt.Printf("Failed to upload file to OSS: %v", err)
		return
	}

	err = c.SaveUploadedFile(form, `./uploads/1.jpg`)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// 创建aliyun客户端和容器空间
func CreateAliyunClient() (bucket *oss.Bucket) {
	// 创建OSS客户端
	client, err := oss.New(Endpoint, AccessKeyId, AccessKeySecret)
	if err != nil {
		fmt.Printf("Failed to create OSS client: %v", err)
		return
	}

	// 获取存储空间
	bucket, err = client.Bucket(BucketName)
	if err != nil {
		fmt.Printf("Failed to get OSS bucket: %v", err)
		return
	}
	return bucket

}

// GinServer webserver
func GinServer() {
	router := gin.Default()

	router.POST("/images", UploadImage)

	router.Run("0.0.0.0:8091")
}
