package environment

type Environment struct {
	Version  string `env:"VERSION"`
	HostName string `env:"HOSTNAME"`
}
