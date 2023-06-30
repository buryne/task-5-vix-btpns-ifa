package controllers

import (
	"crud/database"
	"crud/models"

	"github.com/gin-gonic/gin"
)

func PhotoGetAll(c *gin.Context) {

	var photos []models.Photo
	initial.DB.Find(&photos)

	c.JSON(200, gin.H{
		"photos": photos,
	})
}

func PhotoUpdate(c *gin.Context) {
	// get id from url
	id := c.Param("photoId")

	// get data req body
	var body struct {
		Title    string
		Caption  string
		PhotoUrl string
	}

	c.Bind(&body)

	var photo models.Photo
	initial.DB.First(&photo, id)

	// update it
	initial.DB.Model(&photo).Updates(models.Photo{PhotoUrl: body.PhotoUrl}) // pass pointer of data to Create

	c.JSON(200, gin.H{
		"photo": photo,
	})
}

func PhotoDelete(c *gin.Context) {
	// get id from url
	id := c.Param("photoId")

	// delete post
	initial.DB.Unscoped().Delete(&models.Photo{}, id)

	// response
	c.JSON(200, gin.H{
		"message": `success delete post`,
	})
}
