package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	API_PORT string
	DB_TYPE  string
	DB_PORT  string
	DB_HOST  string
	DB_USER  string
	DB_PASS  string
	DB_NAME  string
}

func Load() *Config {

	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	ConfigEnv := Config{
		API_PORT: os.Getenv("API_PORT"),
		DB_TYPE:  os.Getenv("DB_TYPE"),
		DB_PORT:  os.Getenv("DB_PORT"),
		DB_HOST:  os.Getenv("DB_HOST"),
		DB_USER:  os.Getenv("DB_USER"),
		DB_PASS:  os.Getenv("DB_PASS"),
		DB_NAME:  os.Getenv("DB_NAME"),
	}

	return &ConfigEnv

}
