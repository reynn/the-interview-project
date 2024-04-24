package config

import (
	"github.com/caarlos0/env/v11"
)

type Application struct{
	ServerAddr string `env:"SERVER_ADDR" envDefault:"127.0.0.1:8080"`
	Debug bool `env:"DEBUG"`
}

func NewFromEnv() (*Application, error) {
	cfg := Application{}
	parseErr := env.Parse(&cfg)
	return &cfg, parseErr
}
