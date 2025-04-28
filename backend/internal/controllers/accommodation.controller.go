package controllers

import (
	"encoding/json"
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
		fmt.Println("error: ", err.Error())
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid)
		return
	}

	// process json data
	facilitiesString := ctx.PostForm("facilities")
	fmt.Println("facilitiesString: ", facilitiesString)
	if facilitiesString != "" {
		if err := json.Unmarshal([]byte(facilitiesString), &params.Facilities); err != nil {
			response.ErrorResponse(ctx, response.ErrCodeParamsInvalid)
			return
		}
	}

	fmt.Println("Facilities: ", params.Facilities)

	codeStatus, data, err := services.Accommodation().CreateAccommodation(ctx, &params)
	if err != nil {
		global.Logger.Error("CreateAccommodation: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus)
		return
	}
	global.Logger.Info("CreateAccommodation: ", zap.String("info", data.ManagerId))
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *CAccommodation) GetAccommodations(ctx *gin.Context) {
	codeStatus, data, err := services.Accommodation().GetAccommodations(ctx)
	if err != nil {
		global.Logger.Error("GetAccommodations: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus)
		return
	}
	global.Logger.Info("GetAccommodations: ", zap.String("info", "Get accommodations success"))
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *CAccommodation) UpdateAccommodation(ctx *gin.Context) {
	var params vo.UpdateAccommodationInput
	if err := ctx.ShouldBind(&params); err != nil {
		fmt.Println("error: ", err.Error())
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid)
		return
	}

	// process json data
	facilitiesString := ctx.PostForm("facilities")
	if facilitiesString != "" {
		if err := json.Unmarshal([]byte(facilitiesString), &params.Facilities); err != nil {
			response.ErrorResponse(ctx, response.ErrCodeParamsInvalid)
			return
		}
	}

	fmt.Println("params: ", params)

	codeStatus, data, err := services.Accommodation().UpdateAccommodation(ctx, &params)
	if err != nil {
		global.Logger.Error("UpdateAccommodation: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus)
		return
	}
	global.Logger.Info("UpdateAccommodation: ", zap.String("info", data.ManagerId))
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *CAccommodation) DeleteAccommodation(ctx *gin.Context) {
	var params vo.DeleteAccommodationInput

	id := ctx.Param("id")
	fmt.Println("id: ", id)
	if id == "" {
		fmt.Println("error: ", "Id is empty")
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid)
		return
	}
	params.Id = id

	codeStatus, err := services.Accommodation().DeleteAccommodation(ctx, &params)
	if err != nil {
		global.Logger.Error("DeleteAccommodation: ", zap.String("error", err.Error()))
		fmt.Printf("DeleteAccommodation: %s\n", err.Error())
		response.ErrorResponse(ctx, codeStatus)
		return
	}
	// global.Logger.Info("DeleteAccommodation: ", zap.String("info", ctx.Value("userId").(string)))
	// fmt.Printf("DeleteAccommodation: userId:%s\taccommodationId:%s\n", ctx.Value("userId").(string), params.Id)
	response.SuccessResponse(ctx, codeStatus, nil)
}
