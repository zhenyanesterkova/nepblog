package config

type DataBaseConfig struct {
	PostgresConfig *PostgresConfig
}

type PostgresConfig struct {
	DSN string
}
