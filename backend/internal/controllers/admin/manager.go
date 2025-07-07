package admin

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/global"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/response"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils"
	"go.uber.org/zap"
)

type CAdminManager struct{}

func (c *CAdminManager) GetManagers(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("GetManagers validation not found\n")
		global.Logger.Error("GetManagers validation not found")
		response.ErrorResponse(ctx, response.ErrCodeInternalServerError, nil)
		return
	}

	var params vo.GetManagerInput
	if err := ctx.ShouldBindQuery(&params); err != nil {
		fmt.Printf("GetManagers binding error: %s\n", err.Error())
		global.Logger.Error("GetManagers binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		validationErrors := response.FormatValidationErrorsToStruct(err, params)
		fmt.Printf("GetManagers validation error: %s\n", validationErrors)
		global.Logger.Error("GetManagers validation error: ", zap.Any("error", validationErrors))
		response.ErrorResponse(ctx, response.ErrCodeValidator, validationErrors)
		return
	}

	codeStatus, data, pagination, err := services.AdminManager().GetManagers(ctx, &params)
	if err != nil {
		fmt.Printf("GetManagers error: %s\n", err.Error())
		global.Logger.Error("GetManagers error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	userId, _ := utils.GetUserIDFromGin(ctx)

	fmt.Printf("GetManagers success: %s\n", userId)
	global.Logger.Info("GetManagers success: ", zap.String("info", userId))
	response.SuccessResponseWithPagination(ctx, codeStatus, data, pagination)
}

func (c *CAdminManager) GetAccommodationsOfManager(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("GetAccommodationsOfManager validation not found\n")
		global.Logger.Error("GetAccommodationsOfManager validation not found")
		response.ErrorResponse(ctx, response.ErrCodeInternalServerError, nil)
		return
	}

	var params vo.GetAccommodationsOfManagerInput
	if err := ctx.ShouldBindUri(&params); err != nil {
		fmt.Printf("GetAccommodationsOfManager binding error: %s\n", err.Error())
		global.Logger.Error("GetAccommodationsOfManager binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		validationErrors := response.FormatValidationErrorsToStruct(err, params)
		fmt.Printf("GetAccommodationsOfManager validation error: %s\n", validationErrors)
		global.Logger.Error("GetAccommodationsOfManager validation error: ", zap.Any("error", validationErrors))
		response.ErrorResponse(ctx, response.ErrCodeValidator, validationErrors)
		return
	}

	codeStatus, data, pagination, err := services.AdminManager().GetAccommodationsOfManager(ctx, &params)
	if err != nil {
		fmt.Printf("GetAccommodationsOfManager error: %s\n", err.Error())
		global.Logger.Error("GetAccommodationsOfManager error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	userId, _ := utils.GetUserIDFromGin(ctx)

	fmt.Printf("GetAccommodationsOfManager success: %s\n", userId)
	global.Logger.Info("GetAccommodationsOfManager success: ", zap.String("info", userId))
	response.SuccessResponseWithPagination(ctx, codeStatus, data, pagination)
}

func (c *CAdminManager) VerifyAccommodation(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("VerifyAccommodation validation not found\n")
		global.Logger.Error("VerifyAccommodation validation not found")
		response.ErrorResponse(ctx, response.ErrCodeInternalServerError, nil)
		return
	}

	var params vo.VerifyAccommodationInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		fmt.Printf("VerifyAccommodation binding error: %s\n", err.Error())
		global.Logger.Error("VerifyAccommodation binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		validationErrors := response.FormatValidationErrorsToStruct(err, params)
		fmt.Printf("VerifyAccommodation validation error: %s\n", validationErrors)
		global.Logger.Error("VerifyAccommodation validation error: ", zap.Any("error", validationErrors))
		response.ErrorResponse(ctx, response.ErrCodeValidator, validationErrors)
		return
	}

	codeStatus, err := services.AdminManager().VerifyAccommodation(ctx, &params)
	if err != nil {
		fmt.Printf("VerifyAccommodation error: %s\n", err.Error())
		global.Logger.Error("VerifyAccommodation error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	userId, _ := utils.GetUserIDFromGin(ctx)

	fmt.Printf("VerifyAccommodation success: %s\n", userId)
	global.Logger.Info("VerifyAccommodation success: ", zap.String("info", userId))
	response.SuccessResponse(ctx, codeStatus, nil)
}

func (c *CAdminManager) SetDeletedAccommodation(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("SetDeletedAccommodation validation not found\n")
		global.Logger.Error("SetDeletedAccommodation validation not found")
		response.ErrorResponse(ctx, response.ErrCodeInternalServerError, nil)
		return
	}
	var params vo.SetDeletedAccommodationInput
	if err := ctx.ShouldBind(&params); err != nil {
		fmt.Printf("SetDeletedAccommodation binding error")
		global.Logger.Error("SetDeletedAccommodation binding error")
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		validationErrors := response.FormatValidationErrorsToStruct(err, params)
		fmt.Printf("SetDeletedAccommodation validation error: %s\n", validationErrors)
		global.Logger.Error("SetDeletedAccommodation validation error: ", zap.Any("error", validationErrors))
		response.ErrorResponse(ctx, response.ErrCodeValidator, validationErrors)
		return
	}

	codeStatus, err := services.AdminManager().SetDeletedAccommodation(ctx, &params)
	if err != nil {
		fmt.Printf("SetDeletedAccommodation error: %s\n", err.Error())
		global.Logger.Error("SetDeletedAccommodation error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	userId, _ := utils.GetUserIDFromGin(ctx)

	fmt.Printf("SetDeletedAccommodation success: userId:%s\naccommodationId:%s\n", userId, params.AccommodationID)
	global.Logger.Info("SetDeletedAccommodation success: ",
		zap.String("info", fmt.Sprintf("userId:%s\naccommodationId:%s", userId, params.AccommodationID)))
	response.SuccessResponse(ctx, codeStatus, nil)
}
