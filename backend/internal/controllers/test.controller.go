package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/vo"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/pkg/response"
)

var Test = new(CTest)

type CTest struct {
}

func (c *CTest) CreateAccommodation(ctx *gin.Context) {
	var params vo.CreateAccommodationInput
	if err := ctx.ShouldBind(&params); err != nil {
		fmt.Println("error: ", err.Error())
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid)
		return
	}

	fmt.Println("params: ", params)

	// github.com/go-playground/validator/v10 v10.26.0 // indirect

	// codeStatus, data, err := services.Test().CreateAccommodation(ctx, &params)
	// if err != nil {
	// 	global.Logger.Error("CreateAccommodation: ", zap.String("error", err.Error()))
	// 	response.ErrorResponse(ctx, codeStatus)
	// 	return
	// }
	// global.Logger.Info("CreateAccommodation: ", zap.String("info", data.ManagerId))
	response.SuccessResponse(ctx, 200, nil)
}
