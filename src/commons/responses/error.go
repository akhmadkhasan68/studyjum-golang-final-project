package response

import (
	"final-project/src/utils/validator"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ErrDuplicateEntry struct{}

func (e *ErrDuplicateEntry) Error() string {
	return "data already exists"
}

type ErrDuplicateUniqueColumn struct {
	Column string
}

func (e *ErrDuplicateUniqueColumn) Error() string {
	if e.Column == "" {
		e.Column = "Data"
	}
	return fmt.Sprintf("%s already exist", e.Column)
}

func NewErrDuplicateUniqueColumn(column string) *ErrDuplicateUniqueColumn {
	return &ErrDuplicateUniqueColumn{column}
}

type ErrNotFound struct{}

func (e *ErrNotFound) Error() string {
	return "data not found"
}

type ErrUnauthorized struct {
	Message string
}

type ErrResponse struct {
	Message     string
	Status      uint64
	Description string
}

func (e *ErrUnauthorized) Error() string {
	if e.Message == "" {
		e.Message = "unauthorized"
	}
	return e.Message
}

func NewErrUnauthorized(msg string) *ErrUnauthorized {
	return &ErrUnauthorized{msg}
}

func JSONErrorResponse(ctx *gin.Context, err error) {
	switch v := err.(type) {
	case *ErrDuplicateEntry:
		ctx.JSON(
			http.StatusBadRequest,
			BasicResponse{
				Status:    http.StatusBadRequest,
				Message:   v.Error(),
				Timestamp: time.Now().UnixNano(),
			},
		)

	case *ErrDuplicateUniqueColumn:
		ctx.JSON(
			http.StatusBadRequest,
			BasicResponse{
				Status:    http.StatusBadRequest,
				Message:   v.Error(),
				Timestamp: time.Now().UnixNano(),
			},
		)

	case validator.ValidationErrors:
		ctx.JSON(
			http.StatusBadRequest,
			BasicResponse{
				Status:    http.StatusBadRequest,
				Message:   validator.TranslateError(v).Error(),
				Timestamp: time.Now().UnixNano(),
			},
		)

	case *ErrUnauthorized:
		ctx.JSON(
			http.StatusUnauthorized,
			BasicResponse{
				Status:    http.StatusUnauthorized,
				Message:   v.Error(),
				Timestamp: time.Now().UnixNano(),
			},
		)

	default:
		ctx.JSON(
			http.StatusInternalServerError,
			BasicResponse{
				Status:    http.StatusInternalServerError,
				Message:   v.Error(),
				Timestamp: time.Now().UnixNano(),
			},
		)
	}
}
