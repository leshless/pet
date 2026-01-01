package config

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

	Service     string `yaml:"service" validate:"required"`
	Environment string `yaml:"environment" validate:"required"`
}

type Config struct {
	Logger Logger `yaml:"logger"`
}
