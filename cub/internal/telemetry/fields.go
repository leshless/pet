package telemetry

import (
	"github.com/leshless/golibrary/xslices"
	"go.uber.org/zap"
)

var (
	ErrorKey = "error"

	ServiceKey     = "service"
	HostKey        = "host"
	EnvironmentKey = "environment"

	CallKey   = "call"
	MethodKey = "method"
)

// @PublicValueInstance
type Field struct {
	key   string
	value any
}

func Service(value string) Field {
	return NewField(ServiceKey, value)
}

func Host(value string) Field {
	return NewField(HostKey, value)
}

func Environment(value string) Field {
	return NewField(EnvironmentKey, value)
}

func Call(value string) Field {
	return NewField(CallKey, value)
}

func Method(value string) Field {
	return NewField(MethodKey, value)
}

func Error(value error) Field {
	return NewField(ErrorKey, value)
}

func Any(key string, value any) Field {
	return NewField(key, value)
}

func fieldsToZap(fields []Field) []zap.Field {
	return xslices.Map(fields, func(field Field) zap.Field {
		return zap.Any(field.key, field.value)
	})
}
