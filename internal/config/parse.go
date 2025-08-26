package config

import (
	"os"

	"github.com/BurntSushi/toml"

	"github.com/zaytcevcom/golang-ninja/internal/validator"
)

func ParseAndValidate(filename string) (Config, error) {
	var cfg Config

	data, err := os.ReadFile(filename)
	if err != nil {
		return Config{}, err
	}

	if err := toml.Unmarshal(data, &cfg); err != nil {
		return Config{}, err
	}

	if err := validator.Validator.Struct(cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
