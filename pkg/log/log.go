package log

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack2 "gopkg.in/natefinch/lumberjack.v2"
)

const (
	Console = "console"
	File    = "file"
)

var (
	Level  = zap.DebugLevel
	Target = Console
)

var (
	Logger *zap.Logger
	Sugar  *zap.SugaredLogger
)

func NewEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func NewLogConfig() *lumberjack2.Logger {
	return &lumberjack2.Logger{
		Filename:   "log/im.log", // 日志文件的位置
		MaxSize:    1024,         // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: 10,           // 保留旧文件的最大个数
		MaxAge:     7,            // 保留旧文件的最大天数
		Compress:   false,        // 是否压缩/归档旧文件
	}
}

func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func Init() {
	var writeSyncer zapcore.WriteSyncer
	writeSyncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(NewLogConfig()), zapcore.AddSync(os.Stdout))
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(NewEncoderConfig()),
		writeSyncer,
		Level,
	)
	Logger = zap.New(core, zap.AddCaller())
	Sugar = Logger.Sugar()
}
