package config

const (
	DefaultListenAddress = "0.0.0.0:8123"
)

type App struct {
	ListenAddress string `validate:"required,hostname_port"`
}
