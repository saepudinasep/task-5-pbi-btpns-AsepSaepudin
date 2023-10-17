package router

import (
	"github.com/gin-gonic/gin"
	"github.com/saepudinasep/task-5-pbi-btpns-AsepSaepudin/controllers"
	middlewares "github.com/saepudinasep/task-5-pbi-btpns-AsepSaepudin/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Public routes (tanpa otentikasi)
	r.POST("/users/register", controllers.RegisterUser)
	r.POST("/users/login", controllers.LoginUser)

	// Middleware otentikasi JWT (untuk rute yang memerlukan otentikasi)
	auth := r.Group("/auth")
	auth.Use(middlewares.AuthMiddleware())
	{
		auth.PUT("/users/:userId", controllers.UpdateUser)
		auth.DELETE("/users/:userId", controllers.DeleteUser)
		auth.POST("/photos", controllers.CreatePhoto)
		auth.GET("/photos", controllers.GetPhotos)
		auth.PUT("/photos/:photoId", controllers.UpdatePhoto)
		auth.DELETE("/photos/:photoId", controllers.DeletePhoto)
	}

	return r
}
