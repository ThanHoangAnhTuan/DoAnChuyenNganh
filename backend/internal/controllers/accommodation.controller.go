package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/global"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/services"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/vo"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/pkg/response"
	"go.uber.org/zap"
)

var Accommodation = new(CAccommodation)

type CAccommodation struct {
}

func (c *CAccommodation) CreateAccommodation(ctx *gin.Context) {
	var params vo.CreateAccommodationInput
	if err := ctx.ShouldBind(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid)
		return
	}

	codeStatus, data, err := services.Accommodation().CreateAccommodation(ctx, &params)
	if err != nil {
		fmt.Printf("CreateAccommodation error: %s\n", err.Error())
		global.Logger.Error("CreateAccommodation error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus)
		return
	}

	fmt.Printf("CreateAccommodation success: manager: %s\taccommodation: %s\n", data.ManagerId, data.Id)
	global.Logger.Info("CreateAccommodation success: ",
		zap.String("info", fmt.Sprintf("manager:%s\taccommodation:%s", data.ManagerId, data.Id)))
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *CAccommodation) GetAccommodations(ctx *gin.Context) {
	codeStatus, data, err := services.Accommodation().GetAccommodations(ctx)
	if err != nil {
		fmt.Printf("GetAccommodations error: %s\n", err.Error())
		global.Logger.Error("GetAccommodations error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus)
		return
	}

	fmt.Printf("GetAccommodations success: %s\n", "Get accommodations success")
	global.Logger.Info("GetAccommodations success: ", zap.String("info", "Get accommodations success"))
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *CAccommodation) UpdateAccommodation(ctx *gin.Context) {
	var params vo.UpdateAccommodationInput
	if err := ctx.ShouldBind(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid)
		return
	}

	codeStatus, data, err := services.Accommodation().UpdateAccommodation(ctx, &params)
	if err != nil {
		fmt.Printf("UpdateAccommodation error: %s\n", err.Error())
		global.Logger.Error("UpdateAccommodation error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus)
		return
	}

	fmt.Printf("UpdateAccommodation success: manager%s\taccommodation:%s\n", data.ManagerId, data.Id)
	global.Logger.Info("UpdateAccommodation success: ",
		zap.String("info", fmt.Sprintf("manager%s\taccommodation:%s", data.ManagerId, data.Id)))
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *CAccommodation) DeleteAccommodation(ctx *gin.Context) {
	var params vo.DeleteAccommodationInput

	id := ctx.Param("id")
	if id == "" {
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid)
		return
	}
	params.Id = id

	codeStatus, err := services.Accommodation().DeleteAccommodation(ctx, &params)
	if err != nil {
		fmt.Printf("DeleteAccommodation error: %s\n", err.Error())
		global.Logger.Error("DeleteAccommodation error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus)
		return
	}

	fmt.Printf("DeleteAccommodation success: userId:%s\taccommodationId:%s\n", "1", params.Id)
	global.Logger.Info("DeleteAccommodation success: ",
		zap.String("info", fmt.Sprintf("userId:%s\taccommodationId:%s", "1", params.Id)))
	response.SuccessResponse(ctx, codeStatus, nil)
}

func (c *CAccommodation) GetAccommodationsByManager(ctx *gin.Context) {
	codeStatus, data, err := services.Accommodation().GetAccommodationsByManager(ctx)
	if err != nil {
		fmt.Printf("GetAccommodationsByManager error: %s\n", err.Error())
		global.Logger.Error("GetAccommodationsByManager error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus)
		return
	}

	fmt.Printf("GetAccommodationsByManager success: %s\n", "Get accommodations by manager success")
	global.Logger.Info("GetAccommodationsByManager success: ",
		zap.String("info", "Get accommodations by manager success"))
	response.SuccessResponse(ctx, codeStatus, data)
}
