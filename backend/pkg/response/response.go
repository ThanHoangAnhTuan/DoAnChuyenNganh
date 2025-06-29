package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
)

type Response struct {
	Code       int                      `json:"code"`
	Message    string                   `json:"message"`
	Data       interface{}              `json:"data,omitempty"`
	Error      interface{}              `json:"error,omitempty"`
	Pagination *vo.BasePaginationOutput `json:"pagination,omitempty"`
}

func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message[code],
		Data:    data,
	})
}

func SuccessResponseWithPagination(c *gin.Context, code int, data interface{}, pagination *vo.BasePaginationOutput) {
	c.JSON(http.StatusOK, Response{
		Code:       code,
		Message:    message[code],
		Data:       data,
		Pagination: pagination,
	})
}

func ErrorResponse(c *gin.Context, code int, err interface{}) {
	c.JSON(http.StatusBadRequest, Response{
		Code:    code,
		Message: message[code],
		Error:   err,
	})
}
