package internal

import (
	"os"
)

var cfg *config

type config struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBName     string
}

func NewCfg() *config {
	if cfg != nil {
		return cfg
	}
	newCfg := &config{
		DBUsername: os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBName:     os.Getenv("DB_NAME"),
	}

	cfg = newCfg
	return cfg
}
