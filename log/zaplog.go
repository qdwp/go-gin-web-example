package log

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger

func InitLogger() {
	var writeSyncer zapcore.WriteSyncer
	logMode := os.Getenv("LOG_MODE")
	if "file" == logMode {
		hook := lumberjack.Logger{
			Filename:   "app.log",
			MaxSize:    100, // MB
			MaxBackups: 10,
			MaxAge:     7, // days
			LocalTime:  true,
			Compress:   true, // disabled by default
		}
		writeSyncer = zapcore.AddSync(&hook)
	} else {
		writeSyncer = zapcore.AddSync(os.Stdout)
	}

	// 动态调整日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = customTimeEncoder

	encoder := zapcore.NewJSONEncoder(encoderConfig)

	core := zapcore.NewCore(
		encoder,
		writeSyncer,
		atomicLevel,
	)

	logger = zap.New(core)
}

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

// Info

func Info(args ...interface{}) {
	logger.Sugar().Info(args)
}

func Infof(template string, args ...interface{}) {
	logger.Sugar().Infof(template, args)
}

func InfoWithFields(msg string, fields map[string]string) {
	zapFields := make([]zap.Field, 0, len(fields))
	for k, v := range fields {
		zapFields = append(zapFields, zap.String(k, v))
	}
	logger.Info(msg, zapFields...)
}

// Error

func Error(args ...interface{}) {
	logger.Sugar().Error(args)
}

func Errorf(template string, args ...interface{}) {
	logger.Sugar().Errorf(template, args)
}

func ErrorWithFields(msg string, fields map[string]string) {
	zapFields := make([]zap.Field, 0, len(fields))
	for k, v := range fields {
		zapFields = append(zapFields, zap.String(k, v))
	}
	logger.Error(msg, zapFields...)
}

// Warn

func Warn(args ...interface{}) {
	logger.Sugar().Warn(args)
}

func Warnf(template string, args ...interface{}) {
	logger.Sugar().Warnf(template, args)
}

func WarnWithFields(msg string, fields map[string]string) {
	zapFields := make([]zap.Field, 0, len(fields))
	for k, v := range fields {
		zapFields = append(zapFields, zap.String(k, v))
	}
	logger.Warn(msg, zapFields...)
}

// Debug

func Debug(args ...interface{}) {
	logger.Sugar().Debug(args)
}

func Debugf(template string, args ...interface{}) {
	logger.Sugar().Debugf(template, args)
}

func DebugWithFields(msg string, fields map[string]string) {
	zapFields := make([]zap.Field, 0, len(fields))
	for k, v := range fields {
		zapFields = append(zapFields, zap.String(k, v))
	}
	logger.Debug(msg, zapFields...)
}

// Fatal

func Fatal(args ...interface{}) {
	logger.Sugar().Fatal(args)
}

func Fatalf(template string, args ...interface{}) {
	logger.Sugar().Fatalf(template, args)
}

func FatalWithFields(msg string, fields map[string]string) {
	zapFields := make([]zap.Field, 0, len(fields))
	for k, v := range fields {
		zapFields = append(zapFields, zap.String(k, v))
	}
	logger.Fatal(msg, zapFields...)
}
