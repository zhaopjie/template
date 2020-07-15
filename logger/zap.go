/**
 * @Author: duke
 * @Description:
 * @File:  zap
 * @Version: 1.0.0
 * @Date: 2020/6/18 4:56 下午
 */

package logger

import (
	"os"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	once    sync.Once
	Lever   string
	OutFile string
}

var lg *zap.Logger

// NewLogger
func NewLogger(level, outFile string) *zap.Logger {
	lg := &Logger{
		Lever:   level,
		OutFile: outFile,
	}
	lg.once.Do(func() {
		lg.initLogger()
	})
	return GetLogger()
}

func GetLogger() *zap.Logger {
	return lg
}

// initLogger 启动日志服务
func (l *Logger) initLogger() {
	//日志输出结构
	enc := zapcore.NewJSONEncoder(l.NewEncoderConfig())
	//标准输出
	ws := zapcore.Lock(os.Stdout)
	enab := zap.NewAtomicLevelAt(getLevel(l.Lever))
	core := zapcore.NewCore(enc, ws, enab)
	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	lg = zap.New(core, caller, development)
}

// NewEncoderConfig 新的编码器配置
func (l *Logger) NewEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "log",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, //小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,    //标准时间
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
}

// getLevel 获取日志输出等级
func getLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

func String(k, v string) zap.Field {
	return zap.String(k, v)
}

func Duration(k string, d time.Duration) zap.Field {
	return zap.Duration(k, d)
}

func Float64(key string, val float64) zap.Field {
	return zap.Float64(key, val)
}

func Time(key string, val time.Time) zap.Field {
	return zap.Time(key, val)
}

func Int(k string, i int) zap.Field {
	return zap.Int(k, i)
}

func Array(key string, val zapcore.ArrayMarshaler) zap.Field {
	return zap.Array(key, val)
}

func Int64(k string, i int64) zap.Field {
	return zap.Int64(k, i)
}

func Error(v error) zap.Field {
	return zap.Error(v)
}

func Object(key string, val zapcore.ObjectMarshaler) zap.Field {
	return zap.Object(key, val)
}
