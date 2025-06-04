package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
	"os"
)

// errorEncoder 自定义 errorEncoder
type errorEncoder struct {
	zapcore.Encoder
	errFile    *os.File
	file       *os.File
	currentDay string
	logRootDir string
}

func (e *errorEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (buf *buffer.Buffer, err error) {
	buf, err = e.Encoder.EncodeEntry(entry, fields)
	if err != nil {
		return nil, err
	}
	//
	switch entry.Level {
	case zap.ErrorLevel:

		if e.errFile == nil {

			file, err := os.OpenFile("err.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
			if err != nil {
				return nil, err
			}
			e.errFile = file
		}

		e.errFile.Write(buf.Bytes())
	}

	return buf, nil
}

func level() {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05") // 替换时间格式化方式
	// 创建自定义的 Encoder
	encoder := &errorEncoder{
		Encoder: zapcore.NewConsoleEncoder(cfg.EncoderConfig), // 使用 Console 编码器
	}

	// 自定义日志输出
	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zap.InfoLevel)

	logger := zap.New(core, zap.AddCaller())

	for i := 0; i < 10; i++ {
		logger.Info("日志级别分片 ")
		logger.Error("日志级别分片")
	}

}
