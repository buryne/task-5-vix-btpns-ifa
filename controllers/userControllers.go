package controllers

import (
	"crud/database"
	"crud/models"
	"net/http"
	"os"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	// get the email/pass of req.body

	var body struct {
		UserName string
		Email    string
		Password string
		Title    string
		Caption  string
		PhotoUrl string
	}

	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{
			"error": "bad request",
		})

		return

	}

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "failed has passoword",
		})

		return
	}
	// create user
	user := models.User{Username: body.UserName, Email: body.Email, Password: string(hash), Photo: []models.Photo{models.Photo{Title: body.Title, Caption: body.Caption, PhotoUrl: body.PhotoUrl}}}

	result := initial.DB.Create(&user) 

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed create user",
		})
	}

	c.JSON(200, gin.H{
		"data": result,
	})
}

func Login(c *gin.Context) {
	// get user email pas body req
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{
			"error": "bad request",
		})

		return
	}

	// lool req user
	var user models.User
	initial.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})

		return
	}
	// compare
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid password ",
		})

		return
	}

	// generate jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "fail create token",
		})

		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(200, gin.H{})
}

func UserUpdate(c *gin.Context) {
	// get id from url
	id := c.Param("id")

	// get data req body
	var body struct {
		UserName string
		Email    string
		Password string
	}

	c.Bind(&body)

	// get find post
	var user models.User
	initial.DB.First(&user, id)

	// update it
	initial.DB.Model(&user).Updates(models.User{Username: body.UserName}) // pass pointer of data to Create

	c.JSON(200, gin.H{
		"user": user,
	})
}

func UserGetAll(c *gin.Context) {

	var users []models.User
	initial.DB.Find(&users)

	c.JSON(200, gin.H{
		"users": users,
	})
}

func UserDelete(c *gin.Context) {
	// get id from url
	id := c.Param("id")
	// delete post
	initial.DB.Unscoped().Delete(&models.User{}, id)

	// response
	c.JSON(200, gin.H{
		"message": `success delete post`,
	})
}
