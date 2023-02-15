package controllers

import (
	"final-project/src/bussiness"
	response "final-project/src/commons/responses"
	"final-project/src/middlewares"
	"final-project/src/requests"
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
	datas, err := c.productService.GetAllProducts()
	if err != nil {
		response.JSONErrorResponse(ctx, err)
		return
	}

	response.JSONBasicData(ctx, http.StatusOK, "Get All Products", responses.ToProductsResponse(datas))
}

func (c *ProductsController) GetDetailProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	data, err := c.productService.DetailProductById(id)
	if err != nil {
		response.JSONErrorResponse(ctx, err)
		return
	}

	response.JSONBasicData(ctx, http.StatusOK, "Get Detail Product", responses.ToProductResponse(data))
}

func (c *ProductsController) CreateProduct(ctx *gin.Context) {
	user, jwterr := c.JWTMiddleware.ExtractJWTUser(ctx)
	if jwterr != nil {
		response.JSONErrorResponse(ctx, jwterr)
		return
	}

	var createProductRequest requests.CreateProductRequest
	if err := ctx.ShouldBind(&createProductRequest); err != nil {
		response.JSONErrorResponse(ctx, err)
		return
	}

	err := c.productService.CreateProduct(createProductRequest, user.ID)
	if err != nil {
		response.JSONErrorResponse(ctx, err)
		return
	}

	response.JSONBasicResponse(ctx, http.StatusCreated, "Success create product!")
}

func (c *ProductsController) UpdateProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	var updateProductRequest requests.CreateProductRequest
	if err := ctx.ShouldBind(&updateProductRequest); err != nil {
		response.JSONErrorResponse(ctx, err)
		return
	}

	err := c.productService.UpdateProduct(id, updateProductRequest)
	if err != nil {
		response.JSONErrorResponse(ctx, err)
		return
	}

	response.JSONBasicResponse(ctx, http.StatusOK, "Success update product!")
}

func (c *ProductsController) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	err := c.productService.DeleteProduct(id)
	if err != nil {
		response.JSONErrorResponse(ctx, err)
		return
	}

	response.JSONBasicResponse(ctx, http.StatusOK, "Success delete product!")
}
