package handlers

import (
	"database/sql"
	"fmt"
	"os"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

func ConfigFromEnv() Config {
	return Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
	}
}

func DatabaseConnection() (*sql.DB, error) {
	cfg := ConfigFromEnv()

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DbName)

	return sql.Open("postgres", psqlInfo)
}
