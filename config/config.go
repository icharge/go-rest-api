package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Database struct {
		Host     string `env:"DB_HOST" env-default:"localhost" yaml:"host"`
		Username string `env:"DB_USERNAME" env-default:"go_user" yaml:"username"`
		Password string `env:"DB_PASSWORD" env-default:"go_password" yaml:"password"`
		DBName   string `env:"DB_NAME" env-default:"go_db" yaml:"db-name"`
	} `yaml:"database"`

	DefaultStatus string `env:"DEFAULT_STATUS" yaml:"defaultStatus"`
}

func (c *Config) LoadEnv() {
	err := cleanenv.ReadConfig("config.yml", c)

	if err != nil {
		panic(err)
	}
}
