package config

import (
	"os"
	"github.com/joho/godotenv" 
	"log"
)  


type Config struct {
	Host string
	Port string
	DB_name string
	Username string
	Password string
	Secret_key string
}


func Get_config() Config {
	err := godotenv.Load("./.env")

    if err != nil {
        log.Fatalf("Error loading .env file")
    }
	config := Config{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		DB_name: os.Getenv("DB_NAME"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Secret_key: os.Getenv("SECRET_KEY"),
	}
	return config
}