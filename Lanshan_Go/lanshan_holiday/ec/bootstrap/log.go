package bootstrap

import (
	"demo/ec/global"
	"demo/ec/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var (
	level   zapcore.Level // zap 日志等级
	options []zap.Option  // zap 配置项
)

func InitializeLog() *zap.Logger {
	// 创建根目录
	createRootDir()

	// 设置日志等级
	setLogLevel()

	if global.Ec.Config.Log.ShowLine {
		options = append(options, zap.AddCaller())
	}

	// 初始化 zap
	return zap.New(getZapCore(), options...)
}

func createRootDir() {
	if ok, _ := utils.PathExists(global.Ec.Config.Log.RootDir); !ok {
		_ = os.Mkdir(global.Ec.Config.Log.RootDir, os.ModePerm)
	}
}

func setLogLevel() {
	switch global.Ec.Config.Log.Level {
	case "debug":
		level = zap.DebugLevel
		options = append(options, zap.AddStacktrace(level))
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
		options = append(options, zap.AddStacktrace(level))
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
}

// 扩展 Zap
func getZapCore() zapcore.Core {
	var encoder zapcore.Encoder

	// 调整编码器默认配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("[" + "2006-01-02 15:04:05.000" + "]"))
	}
	encoderConfig.EncodeLevel = func(l zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(global.Ec.Config.Ec.Env + "." + l.String())
	}

	// 设置编码器
	if global.Ec.Config.Log.Format == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	return zapcore.NewCore(encoder, getLogWriter(), level)
}

// 使用 lumberjack 作为日志写入器
func getLogWriter() zapcore.WriteSyncer {
	file := &lumberjack.Logger{
		Filename:   global.Ec.Config.Log.RootDir + "/" + global.Ec.Config.Log.Filename,
		MaxSize:    global.Ec.Config.Log.MaxSize,
		MaxBackups: global.Ec.Config.Log.MaxBackups,
		MaxAge:     global.Ec.Config.Log.MaxAge,
		Compress:   global.Ec.Config.Log.Compress,
	}

	return zapcore.AddSync(file)
}
