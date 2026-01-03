package config

import "time"

type Service struct {
	Name        string `yaml:"name" validate:"required"`
	Environment string `yaml:"environment" validate:"oneof=testing development staging production"`
}

type Logger struct {
	Developement bool   `yaml:"development"`
	StdoutOnly   bool   `yaml:"stdout_only"`
	OutputFile   string `yaml:"output_file" validate:"required"`

	MaxFileSizeMB  int  `yaml:"max_file_size_mb" validate:"gte=0"`
	MaxFilesAmount int  `yaml:"max_backups_amount" validate:"gte=1"`
	MaxFileAgeDays int  `yaml:"max_file_age_days" validate:"gte=1"`
	Compression    bool `yaml:"compression"`

	TimeKey       string `yaml:"time_key"`
	LevelKey      string `yaml:"level_key"`
	NameKey       string `yaml:"name_key"`
	CallerKey     string `yaml:"caller_key"`
	MessageKey    string `yaml:"message_key"`
	StacktraceKey string `yaml:"stacktrace_key"`
}

type GRPC struct {
	Host string `yaml:"host" validate:"ip"`
	Port uint   `yaml:"port" validate:"port"`

	EnableTLS bool `yaml:"enable_tls"`

	EnableHealth     bool `yaml:"enable_health"`
	EnableReflection bool `yaml:"enable_reflection"`

	KeepaliveMaxConnectionIdle     time.Duration `yaml:"keepalive_max_connection_idle"`
	KeepaliveMaxConnectionAge      time.Duration `yaml:"keepalive_max_connection_age"`
	KeepaliveMaxConnectionAgeGrace time.Duration `yaml:"keepalive_max_connection_age_grace"`
	KeepaliveTime                  time.Duration `yaml:"keepalive_time"`
	KeepaliveTimeout               time.Duration `yaml:"keepalive_timeout"`

	KeepaliveMinTime             time.Duration `yaml:"keepalive_min_time"`
	KeepalivePermitWithoutStream bool          `yaml:"keepalive_permit_without_stream"`

	MaxReceiveMessageSizeMB int `yaml:"max_receive_message_size_mb" validate:"gte=4"`
	MaxSendMessageSizeMB    int `yaml:"max_send_message_size_mb" validate:"gte=4"`

	Timeout time.Duration `yaml:"timeout"`
}

type Config struct {
	Service Service `yaml:"service"`
	Logger  Logger  `yaml:"logger"`
	GRPC    GRPC    `yaml:"grpc"`
}
