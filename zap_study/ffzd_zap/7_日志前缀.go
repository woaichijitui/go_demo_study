package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
	"os"
)

type ZapCore struct {
	zapcore.Encoder
}

// EncodeEntry 实现了 zapcore.Core 接口中的 EncodeEntry 方法。
// 它负责将日志条目和关联的字段编码成带有前缀的日志行。
// 参数:
// - Entry: 待编码的日志条目，包含日志的时间戳、级别、消息等信息。
// - fields: 与日志条目相关的一组字段，用于提供更多上下文信息。
// 返回值:
// - *buffer.Buffer: 包含编码后日志行的缓冲区。
// - error: 如果编码过程中发生错误，则返回该错误
func (m *ZapCore) EncodeEntry(Entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	buf, err := m.Encoder.EncodeEntry(Entry, fields)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func Prefix(prefix string) {
	// 使用 zap 的 NewDevelopmentConfig 快速配置
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05") // 替换时间格式化方式
	// 创建自定义的 Encoder
	encoder := &ZapCore{
		Encoder: zapcore.NewConsoleEncoder(cfg.EncoderConfig), // 使用 Console 编码器
	}
	// 创建 Core
	core := zapcore.NewCore(
		encoder,                    // 使用自定义的 Encoder
		zapcore.AddSync(os.Stdout), // 输出到控制台
		zapcore.DebugLevel,         // 设置日志级别
	)
	logger := zap.New(core, zap.AddCaller())
	logger.Info("添加了前缀")
}
