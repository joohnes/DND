package internal

import (
	"os"

	"github.com/joho/godotenv"
)

var cfg *config

type config struct {
	Production bool

	DBUsername string
	DBPassword string
	DBAdress   string
	DBName     string
}

func NewCfg() *config {
	if cfg != nil {
		return cfg
	}
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	newCfg := &config{
		Production: os.Getenv("PRODUCTION") == "true",

		DBUsername: os.Getenv("DBUSER"),
		DBPassword: os.Getenv("DBPASS"),
		DBAdress:   os.Getenv("DBADDR"),
		DBName:     os.Getenv("DBNAME"),
	}

	cfg = newCfg
	return cfg
}
