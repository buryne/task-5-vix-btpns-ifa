package router

import (
	"crud/controllers"
	"crud/middleware"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine) {
	userRouter := r.Group("/users")
	{
		userRouter.GET("", controllers.UserGetAll)
		userRouter.POST("/login", controllers.Login)
		userRouter.POST("/register", controllers.Register)
		userRouter.PUT("/:id", middleware.RequiredAuth, controllers.UserUpdate)
		userRouter.DELETE("/:id", middleware.RequiredAuth, controllers.UserDelete)
	}
}
