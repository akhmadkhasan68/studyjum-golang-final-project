package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const requestID = "X-Request-ID"

func Headers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		traceID := ctx.GetHeader(requestID)

		if strings.TrimSpace(traceID) == "" {
			ctx.Request.Header.Set(requestID, uuid.New().String())
		}

		ctx.Next()
	}
}
