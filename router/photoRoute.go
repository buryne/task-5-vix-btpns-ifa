package router

import (
	"crud/controllers"
	"crud/middleware"

	"github.com/gin-gonic/gin"
)

func SetupPhotoRoutes(r *gin.Engine) {
	photoRouter := r.Group("/photos")
	{
		photoRouter.GET("", middleware.RequiredAuth, controllers.PhotoGetAll)
		photoRouter.PUT("/:photoId", middleware.RequiredAuth, controllers.PhotoUpdate)
		photoRouter.DELETE("/:photoId", middleware.RequiredAuth, controllers.PhotoDelete)
	}
}
