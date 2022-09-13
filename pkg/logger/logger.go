package logger

import (
	"github.com/bulatok/denet_task/internal/config"

	"go.uber.org/zap"
)

var (
	zapLogger = new(zap.Logger)
)

func Init(conf *config.Config) error {
	switch conf.Mode {
	case "PROD":
		return initProd()
	default:
		return initDev()
	}
}

func initDev() (err error) {
	zapLogger, err = zap.NewDevelopment()
	return err
}

func initProd() (err error) {
	zapLogger, err = zap.NewProduction()
	return err
}

func Flush() error {
	return zapLogger.Sync()
}

func Debug(msg string, fields ...zap.Field) {
	zapLogger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	zapLogger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	zapLogger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	zapLogger.Warn(msg, fields...)
}

func Fatal(err error, fields ...zap.Field) {
	zapLogger.Fatal(err.Error(), fields...)
}
