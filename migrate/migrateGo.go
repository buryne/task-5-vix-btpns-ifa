package main

import (
	"crud/database"
	"crud/models"
)

func init() {
	initial.LoadEnvVariabels()
	initial.ConnectToDB()
}

func main() {
	initial.DB.AutoMigrate(&models.User{})
	initial.DB.AutoMigrate(&models.Photo{})
}
