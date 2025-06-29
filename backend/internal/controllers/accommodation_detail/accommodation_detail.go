package accommodation_detail

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/global"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/consts"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/response"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils"
	"go.uber.org/zap"
)

func (c *Controller) CreateAccommodationDetail(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("CreateAccommodationDetail validation not found\n")
		global.Logger.Error("CreateAccommodationDetail validation not found")
		response.ErrorResponse(ctx, response.ErrCodeInternalServerError, nil)
		return
	}

	var params vo.CreateAccommodationDetailInput
	if err := ctx.ShouldBind(&params); err != nil {
		fmt.Printf("CreateAccommodationDetail binding error: %s\n", err.Error())
		global.Logger.Error("CreateAccommodationDetail binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeValidator, err)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		validationErrors := response.FormatValidationErrorsToStruct(err, params)
		fmt.Printf("CreateAccommodationDetail validation error: %s\n", validationErrors)
		global.Logger.Error("CreateAccommodationDetail validation error: ", zap.Any("error", validationErrors))
		response.ErrorResponse(ctx, response.ErrCodeValidator, validationErrors)
		return
	}

	codeStatus, data, err := services.AccommodationDetail().CreateAccommodationDetail(ctx, &params)
	if err != nil {
		fmt.Printf("CreateAccommodationDetail error: %s\n", err.Error())
		global.Logger.Error("CreateAccommodationDetail error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, err)
		return
	}

	userID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		fmt.Printf("CreateAccommodationDetail cannot found userID")
		global.Logger.Error("CreateAccommodationDetail cannot found userID")
		userID = consts.UNKNOWN
	}

	fmt.Printf("CreateAccommodationDetail success: manager: %s, accommodation detail: %s\n", userID, data.ID)
	global.Logger.Info("CreateAccommodationDetail success",
		zap.String("managerId", userID),
		zap.String("accommodationDetailId", data.ID))
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) GetAccommodationDetails(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("GetAccommodationDetails validation not found\n")
		global.Logger.Error("GetAccommodationDetails validation not found")
		response.ErrorResponse(ctx, response.ErrCodeInternalServerError, nil)
		return
	}

	var params vo.GetAccommodationDetailsInput
	if err := ctx.ShouldBindUri(&params); err != nil {
		fmt.Printf("GetAccommodationDetails uri binding error: %s\n", err.Error())
		global.Logger.Error("GetAccommodationDetails uri binding error", zap.Error(err))
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}

	if err := ctx.ShouldBindQuery(&params); err != nil {
		fmt.Printf("GetAccommodationDetails query binding error: %s\n", err.Error())
		global.Logger.Error("GetAccommodationDetails query binding error", zap.Error(err))
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}
	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		validationErrors := response.FormatValidationErrorsToStruct(err, params)
		fmt.Printf("GetAccommodationDetails validation error: %s\n", validationErrors)
		global.Logger.Error("GetAccommodationDetails validation error: ", zap.Any("error", validationErrors))
		response.ErrorResponse(ctx, response.ErrCodeValidator, validationErrors)
		return
	}

	codeStatus, data, err := services.AccommodationDetail().GetAccommodationDetails(ctx, &params)
	if err != nil {
		fmt.Printf("GetAccommodationDetails error: %s\n", err.Error())
		global.Logger.Error("GetAccommodationDetails error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("GetAccommodationDetails success: %s", params.AccommodationID)
	global.Logger.Info("GetAccommodationDetails success", zap.String("accommodationDetailId", params.AccommodationID))
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) GetAccommodationDetailsByManager(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("GetAccommodationDetailsByManager validation not found\n")
		global.Logger.Error("GetAccommodationDetailsByManager validation not found")
		response.ErrorResponse(ctx, response.ErrCodeInternalServerError, nil)
		return
	}

	var params vo.GetAccommodationDetailsByManagerInput
	if err := ctx.ShouldBindUri(&params); err != nil {
		fmt.Printf("GetAccommodationDetailsByManager uri binding error: %s\n", err.Error())
		global.Logger.Error("GetAccommodationDetailsByManager uri binding error", zap.Error(err))
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}

	if err := ctx.ShouldBindQuery(&params); err != nil {
		fmt.Printf("GetAccommodationDetailsByManager query binding error: %s\n", err.Error())
		global.Logger.Error("GetAccommodationDetailsByManager query binding error", zap.Error(err))
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}
	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		validationErrors := response.FormatValidationErrorsToStruct(err, params)
		fmt.Printf("GetAccommodationDetailsByManager validation error: %s\n", validationErrors)
		global.Logger.Error("GetAccommodationDetailsByManager validation error: ", zap.Any("error", validationErrors))
		response.ErrorResponse(ctx, response.ErrCodeValidator, validationErrors)
		return
	}

	codeStatus, data, err := services.AccommodationDetail().GetAccommodationDetailsByManager(ctx, &params)
	if err != nil {
		fmt.Printf("GetAccommodationDetailsByManager error: %s\n", err.Error())
		global.Logger.Error("GetAccommodationDetailsByManager error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("GetAccommodationDetailsByManager success: %s", params.AccommodationID)
	global.Logger.Info("GetAccommodationDetailsByManager success", zap.String("accommodationDetailId", params.AccommodationID))
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) UpdateAccommodationDetail(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("UpdateAccommodationDetail validation not found \n")
		global.Logger.Error("UpdateAccommodationDetail validation not found")
		response.ErrorResponse(ctx, response.ErrCodeInternalServerError, nil)
		return
	}

	var params vo.UpdateAccommodationDetailInput
	if err := ctx.ShouldBind(&params); err != nil {
		fmt.Printf("UpdateAccommodationDetail binding error: %s\n", err.Error())
		global.Logger.Error("UpdateAccommodationDetail binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		validationErrors := response.FormatValidationErrorsToStruct(err, params)
		fmt.Printf("UpdateAccommodationDetailInput validation error: %s\n", validationErrors)
		global.Logger.Error("UpdateAccommodationDetailInput validation error: ", zap.Any("error", validationErrors))
		response.ErrorResponse(ctx, response.ErrCodeValidator, validationErrors)
		return
	}

	codeStatus, data, err := services.AccommodationDetail().UpdateAccommodationDetail(ctx, &params)
	if err != nil {
		fmt.Printf("UpdateAccommodationDetail error: %s\n", err.Error())
		global.Logger.Error("UpdateAccommodationDetail error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	userID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		fmt.Printf("UpdateAccommodationDetail cannot found userID")
		global.Logger.Error("UpdateAccommodationDetail cannot found userID")
		userID = "unknown"
	}

	fmt.Printf("UpdateAccommodationDetail success: manager: %s, accommodation detail: %s\n", userID, data.ID)
	global.Logger.Info("UpdateAccommodationDetail success",
		zap.String("managerId", userID),
		zap.String("accommodationDetailId", data.ID))

	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) DeleteAccommodationDetail(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("DeleteAccommodationDetail validation not found\n")
		global.Logger.Error("DeleteAccommodationDetail validation not found")
		response.ErrorResponse(ctx, response.ErrCodeInternalServerError, nil)
		return
	}

	var params vo.DeleteAccommodationDetailInput
	if err := ctx.ShouldBind(&params); err != nil {
		fmt.Printf("DeleteAccommodationDetail binding error: %s\n", err.Error())
		global.Logger.Error("DeleteAccommodationDetail binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		validationErrors := response.FormatValidationErrorsToStruct(err, params)
		fmt.Printf("DeleteAccommodationDetail validation error: %s\n", validationErrors)
		global.Logger.Error("DeleteAccommodationDetail validation error: ", zap.Any("error", validationErrors))
		response.ErrorResponse(ctx, response.ErrCodeValidator, validationErrors)
		return
	}

	codeStatus, err := services.AccommodationDetail().DeleteAccommodationDetail(ctx, &params)
	if err != nil {
		fmt.Printf("DeleteAccommodationDetail error: %s\n", err.Error())
		global.Logger.Error("DeleteAccommodationDetail error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	userID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		fmt.Printf("CreateAccommodationDetail cannot found userID")
		global.Logger.Error("CreateAccommodationDetail cannot found userID")
		userID = "unknown"
	}

	fmt.Printf("DeleteAccommodationDetail success: manager: %s, accommodation detail: %s\n", userID, params.ID)
	global.Logger.Info("DeleteAccommodationDetail success",
		zap.String("managerId", userID),
		zap.String("accommodationDetailId", params.ID))
	response.SuccessResponse(ctx, codeStatus, nil)
}
