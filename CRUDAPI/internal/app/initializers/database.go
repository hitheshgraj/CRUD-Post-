package initializers

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DB_URL")
	Db,err= gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err !=nil{
		log.Fatal("Failed to connect to DB")
	}
}