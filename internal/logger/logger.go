package logger

import (
	"go-board-api/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

var ProductionEncoderConfig = zapcore.EncoderConfig{
	TimeKey:				"timestamp",
	LevelKey:				"level",
	NameKey:				"logger",
	CallerKey:			"caller",
	MessageKey:			zapcore.OmitKey,
	StacktraceKey: 	"stacktrace",
	LineEnding:     zapcore.DefaultLineEnding,
	EncodeLevel:    zapcore.CapitalLevelEncoder,
	EncodeTime:     zapcore.ISO8601TimeEncoder, 
	EncodeDuration: zapcore.SecondsDurationEncoder,
	EncodeCaller:   zapcore.ShortCallerEncoder,
}

var DevelopmentEncoderConfig = zapcore.EncoderConfig{
	TimeKey:				"timestamp",
	LevelKey:				"level",
	NameKey:				"logger",
	CallerKey:			zapcore.OmitKey,
	MessageKey:			zapcore.OmitKey,
	StacktraceKey: 	zapcore.OmitKey,
	LineEnding:     zapcore.DefaultLineEnding,
	EncodeLevel:    zapcore.CapitalLevelEncoder,
	EncodeTime:     zapcore.ISO8601TimeEncoder, 
	EncodeDuration: zapcore.SecondsDurationEncoder,
	EncodeCaller:   zapcore.FullCallerEncoder,
}

func Init() {
	zapConfig := zap.Config{
		Level:						zap.NewAtomicLevelAt(zapcore.DebugLevel),
		Development: 			false,
		Encoding:					"json",
		EncoderConfig: 		ProductionEncoderConfig,
		OutputPaths: 			[]string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	if config.EnvVar.AppEnv == "development" {
		zapConfig.Development = true
		zapConfig.EncoderConfig = DevelopmentEncoderConfig
	}

	log, _ = zapConfig.Build(zap.AddCallerSkip(1))
	defer log.Sync()
}

func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

func Warn(message string, fields ...zap.Field) {
	log.Warn(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}

func Fatal(message string, fields ...zap.Field) {
	log.Fatal(message, fields...)
}

func Panic(message string, fields ...zap.Field) {
	log.Panic(message, fields...)
}