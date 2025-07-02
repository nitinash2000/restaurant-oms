package routes

import (
	"restaurant-oms/handlers"
	"restaurant-oms/repository"
	"restaurant-oms/services/tables"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func tableRoutes(r *gin.RouterGroup, client *mongo.Client) {
	tableGroup := r.Group("/tables")

	tableRepo := repository.NewTableRepo(client, "restaurant-oms", "tables")
	tableService := tables.NewTableService(tableRepo)
	tableHandler := handlers.NewTableHandler(tableService)

	tableGroup.GET("/:id", tableHandler.GetTable)
	tableGroup.POST("", tableHandler.CreateTable)
	tableGroup.PUT("/:id", tableHandler.UpdateTable)
	tableGroup.DELETE("/:id", tableHandler.DeleteTable)
}
