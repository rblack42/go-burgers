package config

var defaultConfig = &UiConfig{
	Width:	800,
	Height:	600,
	Title:	"Burgers Explorer",
}

type UiConfig struct {
	Width	int
	Height	int
	Title	string
}

func GetConfig() (*UiConfig) {
	cfg := defaultConfig
	return cfg
}
