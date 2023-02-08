package response

import (
	"github.com/gin-gonic/gin"
)

func JSONErrorResponse(ctx *gin.Context, err error) {
	switch v := err.(type) {

	}
}
