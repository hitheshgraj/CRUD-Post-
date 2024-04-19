package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Server      server
	ShippingUrl string
}

type server struct {
	Port string
}

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// return &AppConfig{
	// 	Server: server{
	// 		Port: "8080",
	// 	},

	//	ShippingUrl: "http://localhost:8080",
}
