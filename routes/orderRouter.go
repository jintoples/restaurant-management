package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jintoples/restaurant-management/controllers"
)

func OrderRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orders", controllers.GetOrders())
	incomingRoutes.GET("/orders/:order_id", controllers.GetOrders())
	incomingRoutes.POST("/orders", controllers.CreateOrder())
	incomingRoutes.PATCH("orders/:order_id", controllers.UpdateOrder())
}
