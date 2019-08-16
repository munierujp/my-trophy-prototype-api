package config

import (
	"os"
	"strings"
)

// Config is config
type Config struct {
	DBHost       string
	DBPort       string
	DBUserName   string
	DBPassword   string
	PORT         string
	AllowOrigins []string
}

// NewConfig is constructor for Config
func NewConfig() Config {
	conf := Config{
		DBHost:       os.Getenv("DB_HOST"),
		DBPort:       os.Getenv("DB_PORT"),
		DBUserName:   os.Getenv("DB_USER_NAME"),
		DBPassword:   os.Getenv("DB_PASSWORD"),
		PORT:         os.Getenv("PORT"),
		AllowOrigins: strings.Split(os.Getenv("ALLOW_ORIGINS"), ","),
	}
	return conf
}
