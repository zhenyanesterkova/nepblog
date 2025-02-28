package config

import (
	"os"
)

func (c *Config) setEnvServerConfig() {
	if envEndpoint, ok := os.LookupEnv("ADDRESS"); ok {
		c.SConfig.Address = envEndpoint
	}
}

func (c *Config) setEnvLoggerConfig() {
	if envLogLevel, ok := os.LookupEnv("LOG_LEVEL"); ok {
		c.LConfig.Level = envLogLevel
	}
}

func (c *Config) setDBConfig() {
	if dsn, ok := os.LookupEnv("DATABASE_DSN"); ok {
		if c.DBConfig.PostgresConfig == nil {
			c.DBConfig.PostgresConfig = &PostgresConfig{}
		}
		c.DBConfig.PostgresConfig.DSN = dsn
	}
}

func (c *Config) envBuild() *Config {
	c.setEnvServerConfig()
	c.setEnvLoggerConfig()
	c.setDBConfig()

	return c
}
