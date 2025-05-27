package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Color() {
	cfg := zap.NewDevelopmentConfig()
	// 设置日志等级
	cfg.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	// 设置日志时间格式化
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	// 设置颜色
	cfg.EncoderConfig.EncodeLevel = coloredLevelEncoder
	logger, _ := cfg.Build()

	logger.Debug("this is Debug logger")
	logger.Info("this is Info logger")
	logger.Warn("this is Warn logger")
	logger.Error("this is Error logger")
	//logger.Fatal("this is Fatal logger")
	logger.Info("this is info",
		zap.String("username", "admin"),
		zap.Int("user_id", 42),
		zap.Bool("active", true))

	// 打印颜色
	logger.Debug("this is \033[31mDebug\033[0m logger")
	logger.Info("this is \033[32mInfo\033[0m logger")
	logger.Warn("this is \033[33mWarn\033[0m logger")
	logger.Error("this is \033[34mError\033[0m logger")
	logger.Fatal("this is \033[35mError\033[0m logger")

}
