package telemetry

import (
	"context"
	"os"

	"github.com/leshless/golibrary/graceful"
	"github.com/leshless/pet/cub/internal/config"
	"github.com/leshless/pet/cub/internal/environment"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/lumberjack.v2"
)

type Logger interface {
	Sync() error
	Debug(message string, fields ...Field)
	Info(message string, fields ...Field)
	Warn(message string, fields ...Field)
	Error(message string, fields ...Field)
	With(fields ...Field) Logger
}

// @PublicPointerInstance
type logger struct {
	zapLogger *zap.Logger
}

var _ Logger = (*logger)(nil)

func InitLogger(
	configHolder config.Holder,
	environmentHolder environment.Holder,
	gracefulRegistrator graceful.Registrator,
) (*logger, error) {
	config := configHolder.Config().Logger
	environment := environmentHolder.Environment()

	level := zap.NewAtomicLevel()

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        config.TimeKey,
		LevelKey:       config.LevelKey,
		NameKey:        config.NameKey,
		CallerKey:      config.CallerKey,
		MessageKey:     config.MessageKey,
		StacktraceKey:  config.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	var encoder zapcore.Encoder
	if config.Developement {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	var writer zapcore.WriteSyncer
	if config.StdoutOnly {
		writer = zapcore.AddSync(os.Stdout)
	} else {
		writer = zapcore.AddSync(&lumberjack.Logger{
			Filename:   config.OutputFile,
			MaxSize:    config.MaxFileSizeMB,
			MaxBackups: config.MaxFilesAmount,
			MaxAge:     config.MaxFileAgeDays,
			Compress:   config.Compression,
		})
	}

	core := zapcore.NewCore(
		encoder,
		writer,
		level,
	)

	baseFields := []Field{
		Service(config.Service),
		Environment(config.Environment),
		Host(environment.HostName),
		Version(environment.Version),
	}

	zapLogger := zap.New(core, zap.Fields(fieldsToZap(baseFields)...))

	logger := NewLogger(zapLogger)

	logger.Info("logger successfully initialized")

	gracefulRegistrator.Register(func(_ context.Context) error {
		return logger.Sync()
	})

	return logger, nil
}

func (l *logger) Sync() error {
	return l.zapLogger.Sync()
}

func (l *logger) Debug(message string, fields ...Field) {
	l.zapLogger.Debug(message, fieldsToZap(fields)...)
}

func (l *logger) Info(message string, fields ...Field) {
	l.zapLogger.Info(message, fieldsToZap(fields)...)
}

func (l *logger) Warn(message string, fields ...Field) {
	l.zapLogger.Warn(message, fieldsToZap(fields)...)
}

func (l *logger) Error(message string, fields ...Field) {
	l.zapLogger.Error(message, fieldsToZap(fields)...)
}

func (l *logger) With(fields ...Field) Logger {
	return NewLogger(l.zapLogger.With(fieldsToZap(fields)...))
}
