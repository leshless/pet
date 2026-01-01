package telemetry

import (
	"github.com/leshless/golibrary/xslices"
	"go.uber.org/zap"
)

var (
	ServiceKey     = "service"
	VersionKey     = "version"
	HostKey        = "host"
	EnvironmentKey = "environment"
)

// @PublicValueInstance
type Field struct {
	key   string
	value any
}

func Service(value string) Field {
	return NewField(ServiceKey, value)
}

func Version(value string) Field {
	return NewField(VersionKey, value)
}

func Host(value string) Field {
	return NewField(HostKey, value)
}

func Environment(value string) Field {
	return NewField(EnvironmentKey, value)
}

func Any(key string, value any) Field {
	return NewField(key, value)
}

func Error(value error) Field {
	return NewField("", value)
}

func fieldsToZap(fields []Field) []zap.Field {
	return xslices.Map(fields, func(field Field) zap.Field {
		return zap.Any(field.key, field.value)
	})
}
