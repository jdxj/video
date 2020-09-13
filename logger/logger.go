package logger

import (
	"os"
	"video/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	sugar *zap.SugaredLogger
)

func Init() {
	core := zapcore.NewCore(
		encoder(),
		syncer(),
		level(),
	)
	sugar = zap.New(core).Sugar()
}

func syncer() zapcore.WriteSyncer {
	switch config.Mod {
	case "debug":
		return zapcore.AddSync(os.Stdout)

	case "release":
		log := config.Log
		rotation := &lumberjack.Logger{
			Filename:   log.Path,
			MaxSize:    log.MaxSize,
			MaxAge:     log.MaxAge,
			MaxBackups: log.MaxBackups,
			LocalTime:  false,
			Compress:   false,
		}
		return zapcore.AddSync(rotation)
	}
	return nil
}

func encoder() zapcore.Encoder {
	switch config.Mod {
	case "debug":
		return zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	case "release":
		return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	}
	return nil
}

func level() zapcore.Level {
	switch config.Mod {
	case "release":
		return zap.InfoLevel
	}
	return zap.DebugLevel
}

func Debug(template string, args ...interface{}) {
	sugar.Debugf(template, args...)
}

func Info(template string, args ...interface{}) {
	sugar.Infof(template, args...)
}

func Warn(template string, args ...interface{}) {
	sugar.Warnf(template, args...)
}

func Error(template string, args ...interface{}) {
	sugar.Errorf(template, args...)
}

func Sync() {
	_ = sugar.Sync()
}
