package controllers

import (
	response "final-project/src/commons/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrdersController struct {
}

func NewOrdersController() *OrdersController {
	return &OrdersController{}
}

func (c *OrdersController) GetAllOrders(ctx *gin.Context) {
	response.JSONBasicData(ctx, http.StatusOK, "Get All Order", "orders")
}

func (c *OrdersController) GetDetailOrder(ctx *gin.Context) {
	id := ctx.Param("id")

	response.JSONBasicData(ctx, http.StatusOK, "Get Detail Order", id)
}
