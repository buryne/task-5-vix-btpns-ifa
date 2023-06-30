package main

import (
	"crud/database"
	"crud/router"
	"github.com/gin-gonic/gin"
)

func init() {
	initial.LoadEnvVariabels()
	initial.ConnectToDB()
}

func main() {
	r := gin.Default()

	router.SetupUserRoutes(r)
	router.SetupPhotoRoutes(r)

	r.Run()
}
