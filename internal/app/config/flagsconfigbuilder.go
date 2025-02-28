package config

import (
	"flag"
)

func (c *Config) setFlagsVariables() {
	flag.StringVar(
		&c.SConfig.Address,
		"a",
		c.SConfig.Address,
		"address and port to run server",
	)

	flag.StringVar(
		&c.LConfig.Level,
		"l",
		c.LConfig.Level,
		"log level",
	)

	dsn := ""
	flag.StringVar(
		&dsn,
		"d",
		dsn,
		"database dsn",
	)

	flag.Parse()

	if isFlagPassed("d") {
		if c.DBConfig.PostgresConfig == nil {
			c.DBConfig.PostgresConfig = &PostgresConfig{}
		}
		c.DBConfig.PostgresConfig.DSN = dsn
	}
}

func (c *Config) flagBuild() *Config {
	c.setFlagsVariables()

	return c
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
