package controller

import (
	"bcg/ecommerce/api/response"
	"bcg/ecommerce/api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ItemsController interface {
	Checkout(ctx *gin.Context)
}

type itemsController struct {
	itemsService service.ItemsService
}

func NewItemsController(itemsService service.ItemsService) ItemsController {
	return itemsController{itemsService: itemsService}
}

func (ic itemsController) Checkout(ctx *gin.Context) {
	var ids []string
	if err := ctx.ShouldBindJSON(&ids); err != nil {
		wrapBadRequestError(ctx, err)
		return
	}
	if err := ic.itemsService.Validate(ids); err != nil {
		wrapBadRequestError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, ic.itemsService.Checkout(ids))
}

func wrapBadRequestError(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, response.Error{Message: err.Error()})
}
