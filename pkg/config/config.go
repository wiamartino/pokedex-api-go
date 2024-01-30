package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
}

func LoadConfig() (Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return Config{}, err
	}

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return Config{}, err
	}

	return Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     port,
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
	}, nil
}
