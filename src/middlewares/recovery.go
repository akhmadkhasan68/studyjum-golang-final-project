package middlewares

import (
	response "final-project/src/commons/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(ctx *gin.Context, err any) {
	if val, ok := err.(error); ok {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			response.ErrResponse{
				Message:     "Internal Server Error",
				Status:      http.StatusInternalServerError,
				Description: val.Error(),
			},
		)
	}
}
