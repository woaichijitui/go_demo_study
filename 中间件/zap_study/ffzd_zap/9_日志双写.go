package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func MultiWriteSyncLogger1() {
	// 使用 zap 的 NewDevelopmentConfig 快速配置
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05") // 替换时间格式化方式
	// 创建自定义的 Encoder
	encoder := &ZapCore{
		Encoder: zapcore.NewConsoleEncoder(cfg.EncoderConfig), // 使用 Console 编码器
	}
	// 创建 输出控制台 Core
	consoleCore := zapcore.NewCore(
		encoder,
		os.Stdout,
		zap.InfoLevel,
	)

	// 创建 输出控制台 Core
	file, _ := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	fileCore := zapcore.NewCore(
		encoder,
		file,
		zap.InfoLevel,
	)
	core := zapcore.NewTee(consoleCore, fileCore)
	//创建 logger
	logger := zap.New(core, zap.AddCaller())
	logger.Info("xxx")

}

func MultiWriteSyncLogger2() {
	// 使用 zap 的 NewDevelopmentConfig 快速配置
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05") // 替换时间格式化方式
	// 创建自定义的 Encoder
	encoder := &ZapCore{
		Encoder: zapcore.NewConsoleEncoder(cfg.EncoderConfig), // 使用 Console 编码器
	}
	// 创建 输出控制台和file Core
	file, _ := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(file, os.Stdout),
		zap.InfoLevel,
	)

	//创建 logger
	logger := zap.New(core, zap.AddCaller())
	logger.Info("xxx")
	logger.Warn("xxx")

}
