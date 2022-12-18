package config

import (
	"os"
	_ "github.com/joho/godotenv/autoload" 

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
	config := Config{
		Host: os.Getenv("HOST"),
		Port: os.Getenv("PORT"),
		DB_name: os.Getenv("DB_NAME"),
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
		Secret_key: os.Getenv("SECRET_KEY"),
	}

	return config
}