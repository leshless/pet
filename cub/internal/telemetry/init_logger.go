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

func InitLogger(
	configHolder config.Holder,
	environmentHolder environment.Holder,
	gracefulRegistrator graceful.Registrator,
) (*logger, error) {
	config := configHolder.Config().Logger
	serviceConfig := configHolder.Config().Service
	environment := environmentHolder.Environment()

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

	var level zapcore.Level
	if config.Developement {
		level = zapcore.DebugLevel
	} else {
		level = zapcore.InfoLevel
	}

	core := zapcore.NewCore(
		encoder,
		writer,
		level,
	)

	baseFields := []Field{
		Service(serviceConfig.Name),
		Environment(serviceConfig.Environment),
		Host(environment.HostName),
	}

	zapLogger := zap.New(core, zap.Fields(fieldsToZap(baseFields)...))

	logger := NewLogger(zapLogger)

	logger.Info("logger successfully initialized")

	gracefulRegistrator.Register(func(_ context.Context) error {
		return logger.Sync()
	})

	return logger, nil
}
