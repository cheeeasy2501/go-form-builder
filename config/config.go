package config

import "github.com/cheeeasy2501/go-form-builder/pkg/database"

type Config struct {
	Database      database.IConfig
}

func NewConfig() *Config {
	return &Config{
		Database:      database.NewPostgresConfig(),
	}
}
