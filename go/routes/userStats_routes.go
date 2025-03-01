package routes

import (
	"spellsbee/controllers"

	"github.com/gin-gonic/gin"
)

func SetupUserStatsRoutes(router *gin.Engine) {
	userStatsGroup := router.Group("/api/stats")
	{
		userStatsGroup.GET("/user/:username", controllers.GetUserStatsByUsername)
		userStatsGroup.PUT("/user/:username", controllers.UpdateUserStats)
	}
}
