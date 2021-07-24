package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"os"
)

func New(target interface{}) error {
	var (
		filename = os.Getenv("CONFIG_FILE")
	)

	if filename == "" {
		filename = ".env"
	}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		if err := envconfig.Process("", target); err != nil {
			return err
		}
		return nil
	}

	if err := godotenv.Load(filename); err != nil {
		return err
	}

	if err := envconfig.Process("", target); err != nil {
		return err
	}

	return nil
}
