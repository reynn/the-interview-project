package config

import (
	"github.com/caarlos0/env/v11"
)

type Application struct{
	ServerAddr string `env:"SERVER_ADDR" envDefault:"127.0.0.1:8080"`
}

func NewFromEnv() (*Application, error) {
	cfg := Application{}
	return &cfg, env.Parse(&cfg)
}
