package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// 定义日志前缀
const (
	LOGPREFIX = "\033[35m[myapp]\033[0m"
)

// 定义打印的颜色常量
const (
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorReset  = "\033[0m"
)

// 自定义 EncodeLevel 函数，为日志级别添加颜色
func coloredLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	switch level {
	case zapcore.DebugLevel:
		enc.AppendString(colorBlue + "DEBUG" + colorReset)
	case zapcore.InfoLevel:
		enc.AppendString(colorGreen + "INFO" + colorReset)
	case zapcore.WarnLevel:
		enc.AppendString(colorYellow + "WARN" + colorReset)
	case zapcore.ErrorLevel, zapcore.DPanicLevel, zapcore.PanicLevel, zapcore.FatalLevel:
		enc.AppendString(colorRed + "ERROR" + colorReset)
	default:
		enc.AppendString(level.String()) // 默认行为
	}
}

// 自定义 Encoder 结构体
type MyEncoder struct {
	zapcore.Encoder
	errFile    *os.File
	file       *os.File
	currentDay string
	logRootDir string
	mu         sync.Mutex
}

// EncodeEntry 方法实现了对日志条目的自定义编码
func (e *MyEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (buf *buffer.Buffer, err error) {
	buf, err = e.Encoder.EncodeEntry(entry, fields)
	if err != nil {
		return nil, err
	}
	// 在日志行的最前面添加前缀
	data := buf.String()
	buf.Reset()
	buf.AppendString(LOGPREFIX + " " + data)

	// 检查是否需要写入新的文件中
	currentDay := time.Now().Format("2006-01-02")
	logDir := filepath.Join(e.logRootDir, "/", currentDay)
	if currentDay != e.currentDay {
		if e.file != nil {
			e.file.Close()
		}
		// 确保日志目录存在，如果不存在则创建
		logDir = e.logRootDir + "/" + currentDay
		if err := os.MkdirAll(e.logRootDir+"/"+currentDay, os.ModePerm); err != nil {
			return nil, err
		}
		// 创建新的文件夹
		filePath := fmt.Sprintf("%s/%s.log", logDir, currentDay)
		file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			return nil, err
		}
		e.file = file
		e.currentDay = currentDay
	}
	e.file.Write(buf.Bytes())

	// err日志输出
	switch entry.Level {
	case zap.ErrorLevel:
		if e.errFile == nil {
			file, err := os.OpenFile(logDir+"/"+"err.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
			if err != nil {
				return nil, err
			}
			e.errFile = file
		}
		e.errFile.Write(buf.Bytes())
	}

	return buf, nil
}

// logInit 函数用于初始化日志系统
func logInit() {
	cfg := zap.NewDevelopmentConfig()
	// 设置日志等级
	cfg.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	// 设置日志时间格式化
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	// 设置颜色
	//cfg.EncoderConfig.EncodeLevel = coloredLevelEncoder
	//
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	// 创建自定义的 Encoder
	encoder := &MyEncoder{
		logRootDir: "./logs",
		Encoder:    zapcore.NewConsoleEncoder(cfg.EncoderConfig), // 使用 Console 编码器
	}

	// 自定义日志输出
	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zap.InfoLevel)

	logger := zap.New(core, zap.AddCaller()).Sugar()

	for i := 0; i < 10; i++ {
		logger.Infof("日志整合test%d", i)
		logger.Warnf("日志整合test%d", i)
		logger.Errorf("日志整合test%d", i)
	}
}
