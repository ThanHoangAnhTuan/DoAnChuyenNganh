package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message[code],
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, code int) {
	c.JSON(http.StatusBadRequest, Response{
		Code:    code,
		Message: message[code],
		Data:    nil,
	})
}
