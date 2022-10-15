package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jintoples/restaurant-management/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
	incomingRoutes.POST("/users/signup", controllers.GetUsers())
	incomingRoutes.POST("users/login", controllers.Login())
}
