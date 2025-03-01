package routes

import (
	"github.com/gin-gonic/gin"
	"spellsbee/controllers"
)

// SetupUserRouter attaches user-related routes to the given router
func SetupUserRouter(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/register", controllers.RegisterUser)
		api.POST("/login", controllers.LoginUser)
		api.PUT("/verify-email", controllers.VerifyEmail)
		api.POST("/forgot-password", controllers.ForgotPassword)
		api.POST("/reset-password", controllers.ResetPassword)

		api.GET("/users", controllers.GetUsers)
		api.POST("/users", controllers.CreateUser)
		api.GET("/users/:id", controllers.GetUserByID)
		api.GET("/users/check-user/:username", controllers.CheckUser)
	}
}
 