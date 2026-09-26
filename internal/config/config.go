package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env    string
	Port   string
	Db_URl string
}

func MustLoad() *Config {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		panic("PORT is not set")
	}
	env := os.Getenv("ENV")
	if env == "" {
		panic("ENV is not set")
	}
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		panic("DB URL is not set")
	}
	return &Config{
		Env:    env,
		Port:   port,
		Db_URl: dbUrl,
	}
}
