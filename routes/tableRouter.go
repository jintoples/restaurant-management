package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jintoples/restaurant-management/controllers"
)

func TableRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/tables", controllers.GetTables())
	incomingRoutes.GET("/tables/:table_id", controllers.GetTables())
	incomingRoutes.POST("/tables", controllers.CreateTable())
	incomingRoutes.PATCH("tables/:table_id", controllers.UpdateTable())
}
