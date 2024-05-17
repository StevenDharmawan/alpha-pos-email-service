package config

import (
	"os"
	"strconv"
)

type EmailConfig struct {
	Name     string
	Host     string
	Port     int
	Email    string
	Password string
}

func NewEmailConfig() *EmailConfig {
	name := os.Getenv("CONFIG_NAME")
	host := os.Getenv("CONFIG_HOST")
	port := os.Getenv("CONFIG_PORT")
	portInt, _ := strconv.Atoi(port)
	email := os.Getenv("CONFIG_EMAIL")
	password := os.Getenv("CONFIG_PASSWORD")
	return &EmailConfig{Name: name, Host: host, Port: portInt, Email: email, Password: password}
}
