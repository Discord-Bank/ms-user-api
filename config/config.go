package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Environments define the environment variables
type Environments struct {
	APIPort string `envconfig:"API_PORT"`
	DSN     string `envconfig:"DSN"`
}

// LoadEnvVars load the environment variables
func LoadEnvVars() (*Environments, error) {
	godotenv.Load()
	c := &Environments{}
	if err := envconfig.Process("", c); err != nil {
		return nil, err
	}
	return c, nil
}
