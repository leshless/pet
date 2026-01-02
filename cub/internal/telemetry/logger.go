package telemetry

import (
	"go.uber.org/zap"
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
