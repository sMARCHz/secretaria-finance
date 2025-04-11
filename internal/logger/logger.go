package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Debug(args ...any)
	Info(args ...any)
	Warn(args ...any)
	Error(args ...any)
	Fatal(args ...any)

	Debugf(format string, args ...any)
	Infof(format string, args ...any)
	Warnf(format string, args ...any)
	Errorf(format string, args ...any)
	Fatalf(format string, args ...any)
}

type logger struct {
	*zap.SugaredLogger
}

func New() Logger {
	logger := newZapProduction()
	return NewWithZap(logger)
}

func NewWithZap(l *zap.Logger) Logger {
	return &logger{l.Sugar()}
}

func newZapProduction() *zap.Logger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config := zap.NewProductionConfig()
	config.EncoderConfig = encoderConfig
	config.OutputPaths = []string{"logs/secretaria.log", "stderr"}

	log, err := config.Build()
	if err != nil {
		panic(err)
	}
	return log
}
