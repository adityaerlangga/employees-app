package config

import (
    "os"
)

type Config struct {
    DBHost     string
    DBUser     string
    DBPassword string
    DBName     string
    DBPort     string
}

func GetConfig() *Config {
    return &Config{
        DBHost:     os.Getenv("DB_HOST"),
        DBUser:     os.Getenv("DB_USER"),
        DBPassword: os.Getenv("DB_PASSWORD"),
        DBName:     os.Getenv("DB_NAME"),
        DBPort:     os.Getenv("DB_PORT"),
    }
}
