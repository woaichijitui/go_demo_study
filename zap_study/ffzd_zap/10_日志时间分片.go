package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"sync"
	"time"
)

type myWrite struct {
	mu         sync.Mutex
	file       *os.File
	logDir     string //日志目录
	currentDay string
}

// Write 实现了 io.Writer 接口，用于将数据写入到日志文件中。
// 它根据当前日期决定是否需要创建新的日志文件。
// 参数 p 是要写入的日志数据。
// 返回值 n 是成功写入的字节数，err 是可能发生的错误。
func (m *myWrite) Write(p []byte) (n int, err error) {
	// 使用互斥锁确保并发安全
	m.mu.Lock()
	defer m.mu.Unlock()

	// 检查是否需要写入新的文件中
	currentDay := time.Now().Format("2006-06-02 15-04-05")
	if m.currentDay != currentDay {
		// 如果当前有打开的文件，则先关闭
		if m.file != nil {
			m.file.Close()
		}

		// 确保日志目录存在，如果不存在则创建
		if err := os.MkdirAll(m.logDir, os.ModePerm); err != nil {
			return 0, err
		}

		// 创建新的log文件
		filePath := fmt.Sprintf("%s/app-%s.log", m.logDir, currentDay)
		m.file, err = os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return 0, err
		}
		m.currentDay = currentDay
	}

	// 如果是同一天，则直接写入当前日志文件
	return m.file.Write(p)
}

func dynamicLog() {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")

	write := &myWrite{
		logDir: "./logs",
	}

	ConsoleCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(cfg.EncoderConfig),
		zapcore.AddSync(os.Stdout),
		zap.InfoLevel,
	)

	fileCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(cfg.EncoderConfig),
		zapcore.AddSync(write),
		zap.InfoLevel,
	)

	core := zapcore.NewTee(ConsoleCore, fileCore)
	log := zap.New(core, zap.AddCaller()).Sugar()

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		log.Infof("这是时间日志%d", i+1)
	}
}
