package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/global"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/response"
	"go.uber.org/zap"
)

var Facility = new(CFacility)

type CFacility struct {
}

func (c *CFacility) CreateFacility(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("CreateFacility validation not found\n")
		global.Logger.Error("CreateFacility validation not found")
		response.ErrorResponse(ctx, response.ErrCodeValidatorNotFound, nil)
		return
	}

	var params vo.CreateFacilityInput
	if err := ctx.ShouldBind(&params); err != nil {
		fmt.Printf("CreateFacility binding error: %s\n", err.Error())
		global.Logger.Error("CreateFacility binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		validationErrors := response.FormatValidationErrorsToStruct(err)
		fmt.Printf("CreateFacility validation error: %s\n", validationErrors)
		global.Logger.Error("CreateFacility validation error: ", zap.Any("error", validationErrors))
		response.ErrorResponse(ctx, response.ErrCodeValidator, validationErrors)
		return
	}

	codeStatus, data, err := services.Facility().CreateFacility(ctx, &params)
	if err != nil {
		fmt.Printf("CreateFacility error: %s\n", err.Error())
		global.Logger.Error("CreateFacility error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("CreateFacility success: %s\n", data)
	global.Logger.Info("CreateFacility success: ", zap.String("info", fmt.Sprintf("create facility success: %s", data.ID)))
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *CFacility) GetFacilities(ctx *gin.Context) {
	codeResult, data, err := services.Facility().GetFacilities(ctx)
	if err != nil {
		fmt.Printf("GetFacilities error: %s\n", err.Error())
		global.Logger.Error("GetFacilities error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeResult, nil)
		return
	}

	fmt.Printf("GetFacilities success\n")
	global.Logger.Info("GetFacilities success")
	response.SuccessResponse(ctx, codeResult, data)
}
