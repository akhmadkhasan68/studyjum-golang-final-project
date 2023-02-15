package controllers

import (
	"final-project/src/bussiness"
	"final-project/src/commons/enums"
	response "final-project/src/commons/responses"
	"final-project/src/database/models"
	"final-project/src/middlewares"
	"final-project/src/requests"
	"final-project/src/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrdersController struct {
	orderService  *bussiness.OrderService
	JWTMiddleware middlewares.IAuthenticator
}

func NewOrdersController(orderService *bussiness.OrderService, JWTMiddleware middlewares.IAuthenticator) *OrdersController {
	return &OrdersController{orderService, JWTMiddleware}
}

func (c *OrdersController) GetAllOrders(ctx *gin.Context) {
	user, jwterr := c.JWTMiddleware.ExtractJWTUser(ctx)
	if jwterr != nil {
		response.JSONErrorResponse(ctx, jwterr)
		return
	}

	var orders *[]models.Order
	if user.Role == enums.OUTLET {
		data, err := c.orderService.GetAllOrdersOutlet(user.ID)
		if err != nil {
			response.JSONErrorResponse(ctx, err)
			return
		}
		orders = data
	} else {
		data, err := c.orderService.GetAllOrdersMember(user.ID)
		if err != nil {
			response.JSONErrorResponse(ctx, err)
			return
		}
		orders = data
	}

	response.JSONBasicData(ctx, http.StatusOK, "Get All Order", responses.ToOrdersResponse(orders))
}

func (c *OrdersController) GetDetailOrder(ctx *gin.Context) {
	id := ctx.Param("id")
	user, jwterr := c.JWTMiddleware.ExtractJWTUser(ctx)
	if jwterr != nil {
		response.JSONErrorResponse(ctx, jwterr)
		return
	}

	var order *models.Order
	if user.Role == enums.OUTLET {
		data, err := c.orderService.GetDetailOrderOutlet(user.ID, id)
		if err != nil {
			response.JSONErrorResponse(ctx, err)
			return
		}
		order = data
	} else {
		data, err := c.orderService.GetDetailOrderMember(user.ID, id)
		if err != nil {
			response.JSONErrorResponse(ctx, err)
			return
		}
		order = data
	}

	response.JSONBasicData(ctx, http.StatusOK, "Success get detail order", responses.ToOrderResponse(order))
}

func (c *OrdersController) CreateOrder(ctx *gin.Context) {
	user, jwterr := c.JWTMiddleware.ExtractJWTUser(ctx)
	if jwterr != nil {
		response.JSONErrorResponse(ctx, jwterr)
		return
	}

	var createOrderRequest requests.CreateOrderRequest
	if err := ctx.ShouldBindJSON(&createOrderRequest); err != nil {
		response.JSONErrorResponse(ctx, err)
		return
	}

	err := c.orderService.CreateOrder(ctx, user.ID, createOrderRequest)
	if err != nil {
		response.JSONErrorResponse(ctx, err)
		return
	}

	response.JSONBasicResponse(ctx, http.StatusCreated, "Success create order!")
}

func (c *OrdersController) CancelOrder(ctx *gin.Context) {
	id := ctx.Param("id")
	user, jwterr := c.JWTMiddleware.ExtractJWTUser(ctx)
	if jwterr != nil {
		response.JSONErrorResponse(ctx, jwterr)
		return
	}

	err := c.orderService.CancelOrder(ctx, user.ID, id)
	if err != nil {
		response.JSONErrorResponse(ctx, err)
		return
	}

	response.JSONBasicResponse(ctx, http.StatusOK, "Success cancel order!")
}

func (c *OrdersController) ShipOrder(ctx *gin.Context) {

	response.JSONBasicResponse(ctx, http.StatusOK, "Ship Order")
}

func (c *OrdersController) RejectOrder(ctx *gin.Context) {
	id := ctx.Param("id")
	user, jwterr := c.JWTMiddleware.ExtractJWTUser(ctx)
	if jwterr != nil {
		response.JSONErrorResponse(ctx, jwterr)
		return
	}

	err := c.orderService.RejectOrder(ctx, user.ID, id)
	if err != nil {
		response.JSONErrorResponse(ctx, err)
		return
	}

	response.JSONBasicData(ctx, http.StatusOK, "Reject Order", id)
}
