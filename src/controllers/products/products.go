package controllers

import (
	"final-project/src/bussiness"
	response "final-project/src/commons/responses"
	"final-project/src/middlewares"
	"final-project/src/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductsController struct {
	productService *bussiness.ProductService
	JWTMiddleware  middlewares.IAuthenticator
}

func NewProductsController(productService *bussiness.ProductService, JWTMiddleware middlewares.IAuthenticator) *ProductsController {
	return &ProductsController{productService, JWTMiddleware}
}

func (c *ProductsController) GetAllProducts(ctx *gin.Context) {
	datas, err := c.productService.GetAll()
	if err != nil {
		response.JSONErrorResponse(ctx, err)
		return
	}

	response.JSONBasicData(ctx, http.StatusOK, "Get All Products", responses.ToProductsResponse(datas))
}

func (c *ProductsController) GetDetailProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	data, err := c.productService.DetailById(id)
	if err != nil {
		response.JSONErrorResponse(ctx, err)
		return
	}

	response.JSONBasicData(ctx, http.StatusOK, "Get Detail Product", responses.ToProductResponse(data))
}
