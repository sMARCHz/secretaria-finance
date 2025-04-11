package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewZapLogger() *zap.Logger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config := zap.NewProductionConfig()
	config.EncoderConfig = encoderConfig
	config.OutputPaths = []string{"/var/log/secretaria/app.log", "stdout"}
	config.ErrorOutputPaths = []string{"stderr"}

	log, err := config.Build()
	if err != nil {
		panic(err)
	}

	return log.WithOptions(zap.AddCaller(), zap.AddCallerSkip(1))
}
