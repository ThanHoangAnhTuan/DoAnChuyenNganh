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

var AccommodationDetail = new(CAccommodationDetail)

type CAccommodationDetail struct {
}

func (c *CAccommodationDetail) CreateAccommodationDetail(ctx *gin.Context) {
	var params vo.CreateAccommodationDetailInput
	if err := ctx.ShouldBind(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid)
		return
	}

	codeStatus, data, err := services.AccommodationDetail().CreateAccommodationDetail(ctx, &params)
	if err != nil {
		fmt.Printf("CreateAccommodationDetail error: %s\n", err.Error())
		global.Logger.Error("CreateAccommodationDetail error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus)
		return
	}

	fmt.Printf("CreateAccommodationDetail success: manager: %s\taccommodation detail: %s\n", data.ManagerId, data.Id)
	global.Logger.Info("CreateAccommodationDetail success: ",
		zap.String("info", fmt.Sprintf("manager:%s\taccommodation detail:%s", data.ManagerId, data.Id)))
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *CAccommodationDetail) GetAccommodationDetails(ctx *gin.Context) {
	codeStatus, data, err := services.AccommodationDetail().GetAccommodationDetails(ctx)
	if err != nil {
		fmt.Printf("GetAccommodationDetails error: %s\n", err.Error())
		global.Logger.Error("GetAccommodationDetails error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus)
		return
	}

	fmt.Printf("GetAccommodationDetails success: %s\n", "Get accommodations details success")
	global.Logger.Info("GetAccommodationDetails success: ", zap.String("info", "Get accommodations details success"))
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *CAccommodationDetail) UpdateAccommodationDetail(ctx *gin.Context) {
	var params vo.UpdateAccommodationDetailInput
	if err := ctx.ShouldBind(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid)
		return
	}

	codeStatus, data, err := services.AccommodationDetail().UpdateAccommodationDetail(ctx, &params)
	if err != nil {
		fmt.Printf("UpdateAccommodationDetail error: %s\n", err.Error())
		global.Logger.Error("UpdateAccommodationDetail error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus)
		return
	}

	fmt.Printf("UpdateAccommodationDetail success: manager%s\taccommodation detail:%s\n", data.ManagerId, data.Id)
	global.Logger.Info("UpdateAccommodationDetail success: ",
		zap.String("info", fmt.Sprintf("manager%s\taccommodation detail:%s", data.ManagerId, data.Id)))
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *CAccommodationDetail) DeleteAccommodationDetail(ctx *gin.Context) {
	var params vo.DeleteAccommodationDetailInput

	id := ctx.Param("id")
	if id == "" {
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid)
		return
	}
	params.Id = id

	codeStatus, err := services.AccommodationDetail().DeleteAccommodationDetail(ctx, &params)
	if err != nil {
		fmt.Printf("DeleteAccommodationDetail error: %s\n", err.Error())
		global.Logger.Error("DeleteAccommodationDetail error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus)
		return
	}

	fmt.Printf("DeleteAccommodationDetail success: userId:%s\taccommodation detail id:%s\n", "1", params.Id)
	global.Logger.Info("DeleteAccommodationDetail success: ",
		zap.String("info", fmt.Sprintf("userId:%s\taccommodation detail id:%s", "1", params.Id)))
	response.SuccessResponse(ctx, codeStatus, nil)
}

func (c *CAccommodationDetail) GetAccommodationDetailsByManager(ctx *gin.Context) {
	codeStatus, data, err := services.AccommodationDetail().GetAccommodationDetailsByManager(ctx)
	if err != nil {
		fmt.Printf("GetAccommodationDetailsByManager error: %s\n", err.Error())
		global.Logger.Error("GetAccommodationDetailsByManager error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus)
		return
	}

	fmt.Printf("GetAccommodationDetailsByManager success: %s\n", "Get accommodation details by manager success")
	global.Logger.Info("GetAccommodationDetailsByManager success: ",
		zap.String("info", "Get accommodation details by manager success"))
	response.SuccessResponse(ctx, codeStatus, data)
}
