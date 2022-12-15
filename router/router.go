package router

import (
	"bcg/ecommerce/api/controller"
	"bcg/ecommerce/api/service"
	entities "bcg/ecommerce/data/entities"

	"github.com/gin-gonic/gin"
)

func SetupRouter(itemsMap map[string]entities.ItemEntity) *gin.Engine {
	router := gin.Default()
	itemsService := service.NewItemsService(itemsMap)
	itemsController := controller.NewItemsController(itemsService)
	router.POST("/checkout", itemsController.Checkout)
	return router
}
