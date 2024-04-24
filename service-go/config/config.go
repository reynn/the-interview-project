package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type (
	Application struct {
		Debug     bool   `env:"DEBUG"`
		JWTSecret string `env:"JWT_SECRET"`
		GRPC      GRPC
	}

	GRPC struct {
		ServerHost   string `env:"server_host" envDefault:"localhost"`
		UnsecurePort string `env:"unsecure_port" envDefault:"8080"`
	}
)

func Load() (*Application, error) {
	cfg := Application{}

	// keeping this outside of the return in the case we want to do additional checks
	// after parsing before we return
	err := env.Parse(&cfg)

	if cfg.JWTSecret == "" {
		return nil, fmt.Errorf("error loading JWT secret from environment")
	}

	return &cfg, err
}
