package conf

import "github.com/caarlos0/env/v6"

type AppConfig struct {
	Port string `env:"PORT" envDefault:"8001"`
	//DB CONFIG
	LogFormat string `env:"LOG_FORMAT" envDefault:"127.0.0.1"`
	DBHost    string `env:"DB_HOST" envDefault:"localhost"`
	DBPort    string `env:"DB_PORT" envDefault:"5432"`
	DBUser    string `env:"DB_USER" envDefault:"postgres"`
	DBPass    string `env:"DB_PASS" envDefault:"postgres"`
	DBName    string `env:"DB_NAME" envDefault:"postgres"`
	DBSchema  string `env:"DB_SCHEMA" envDefault:"public"`
	EnableDB  string `env:"ENABLE_DB" envDefault:"true"`
	// ENV
	EnvName string `env:"ENV_NAME" envDefault:"dev"`
	// Elasticsearch
	ESAddress string `env:"ES_ADDRESSS" envDefault:"http://localhost:9200"`
}

var config AppConfig

func LoadConfig() {
	_ = env.Parse(&config)
}

func GetConfig() AppConfig {
	return config
}

func LoadEnv() AppConfig {
	return config
}
