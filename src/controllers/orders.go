package controllers

import (
	"final-project/src/bussiness"
	response "final-project/src/commons/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrdersController struct {
	orderService *bussiness.OrderService
}

func NewOrdersController(orderService *bussiness.OrderService) *OrdersController {
	return &OrdersController{orderService}
}

func (c *OrdersController) GetAllOrders(ctx *gin.Context) {
	response.JSONBasicData(ctx, http.StatusOK, "Get All Order", "orders")
}

func (c *OrdersController) GetDetailOrder(ctx *gin.Context) {
	id := ctx.Param("id")

	response.JSONBasicData(ctx, http.StatusOK, "Get Detail Order", id)
}

func (c *OrdersController) CreateOrder(ctx *gin.Context) {
	id := ctx.Param("id")

	response.JSONBasicData(ctx, http.StatusOK, "Create Order", id)
}

func (c *OrdersController) CancelOrder(ctx *gin.Context) {
	id := ctx.Param("id")

	response.JSONBasicData(ctx, http.StatusOK, "Cancel Order", id)
}

func (c *OrdersController) ShipOrder(ctx *gin.Context) {
	id := ctx.Param("id")

	response.JSONBasicData(ctx, http.StatusOK, "Ship Order", id)
}

func (c *OrdersController) RejectOrder(ctx *gin.Context) {
	id := ctx.Param("id")

	response.JSONBasicData(ctx, http.StatusOK, "Reject Order", id)
}
