package confi

import (
	//"CRUDAPI/internal/app/initializers"
	// "database/sql"
	"log"
	// "net/http"
	// "os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	//"gorm.io/gorm"
)

type AppConfig struct {
	Server server

	//ShippingUrl string
}

var DB *gorm.DB

type server struct {
	Port string
}

func NewConfig() *AppConfig {
	//initializers.ConnectToDB()
	var err error
	dsn := "root:root@tcp(127.0.0.1:3306)/crudapi?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	return &AppConfig{

		Server: server{
			Port: "8080",
		},

		//ShippingUrl: "http://localhost:8080",
	}
}
