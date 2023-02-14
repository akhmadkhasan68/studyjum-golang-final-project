package controllers

import (
	response "final-project/src/commons/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductsController struct {
}

func NewProductsController() *ProductsController {
	return &ProductsController{}
}

func (c *ProductsController) GetAllProducts(ctx *gin.Context) {
	response.JSONBasicData(ctx, http.StatusOK, "Get All Products", "products")
}

func (c *ProductsController) GetDetailProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	response.JSONBasicData(ctx, http.StatusOK, "Get Detail Product", id)
}
