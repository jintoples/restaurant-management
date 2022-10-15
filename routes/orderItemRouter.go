package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jintoples/restaurant-management/controllers"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems", controllers.GetOrderItems())
	incomingRoutes.GET("/orderItems/:orderItem_id", controllers.GetOrderItems())
	incomingRoutes.GET("/orderItemsByOrder/:order_id", controllers.GetOrderItemsByOrder())
	incomingRoutes.POST("/orderItems", controllers.CreateOrderItem())
	incomingRoutes.PATCH("orderItems/:orderItem_id", controllers.UpdateOrderItem())
}
