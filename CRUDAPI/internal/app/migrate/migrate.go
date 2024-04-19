package main

import (
	"CRUDAPI/internal/app/confi"
	"CRUDAPI/internal/app/initializers"
	"CRUDAPI/internal/app/models"
)

func init() {
	//initializers.LoadEnvVariables()
	confi.NewConfig()
}
func main() {
	initializers.Db.AutoMigrate(&models.Post{})
}
