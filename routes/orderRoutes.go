package routes

import (
	"restaurant-oms/handlers"
	"restaurant-oms/repository"
	"restaurant-oms/services/orders"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func orderRoutes(r *gin.RouterGroup, client *mongo.Client) {
	orderGroup := r.Group("/orders")

	orderRepo := repository.NewOrderRepo(client, "restaurant-oms", "orders")
	orderService := orders.NewOrderService(orderRepo)
	orderHandler := handlers.NewOrderHandler(orderService)

	orderGroup.GET("/:id", orderHandler.GetOrder)
	orderGroup.POST("", orderHandler.CreateOrder)
	orderGroup.PUT("/:id", orderHandler.UpdateOrder)
	orderGroup.DELETE("/:id", orderHandler.DeleteOrder)
}
