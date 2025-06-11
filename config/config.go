package config

import (
	"os"
	"strconv"
)

type PostgresConfig struct {
    Host     string
    Port     int
    User     string
    Password string
    DBName   string
    SSLMode  string
}

func LoadPostgresConfig() (*PostgresConfig, error) {
    port, err := strconv.Atoi(os.Getenv("DV_PORT"))
    if err != nil {
        port = 5432
    }

    return &PostgresConfig{
        User:     os.Getenv("DB_USER"),
        Password: os.Getenv("DB_PASSWORD"),
        Host:     os.Getenv("DB_HOST"),
        Port:     port,
        DBName:   os.Getenv("DB_NAME"),
        SSLMode:  os.Getenv("DB_SSLMODE"),
    }, nil
}
