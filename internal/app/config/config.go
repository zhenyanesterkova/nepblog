package config

type Config struct {
	DBConfig    DataBaseConfig
	SConfig     ServerConfig
	LConfig     LoggerConfig
	RetryConfig RetryConfig
}

func New() *Config {
	return &Config{
		SConfig: ServerConfig{
			Address: DefaultServerAddress,
		},
		LConfig: LoggerConfig{
			Level: DefaultLogLevel,
		},
	}
}

func (c *Config) Build() {
	c.flagBuild().envBuild()
}
