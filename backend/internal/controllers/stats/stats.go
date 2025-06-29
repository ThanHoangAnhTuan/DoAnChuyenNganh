package stats

import (
	"fmt"
	"net/http"
	"strconv"

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

func (c *Controller) GetMonthlyEarnings(ctx *gin.Context) {
	codeStatus, data, err := services.Stats().GetMonthlyEarnings(ctx)
	if err != nil {
		fmt.Printf("GetMonthlyEarnings error: %s\n", err.Error())
		global.Logger.Error("GetMonthlyEarnings error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	userID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		fmt.Printf("GetMonthlyEarnings cannot found userID")
		global.Logger.Error("GetMonthlyEarnings cannot found userID")
		userID = "unknown"
	}

	fmt.Printf("GetMonthlyEarnings success: %s\n", userID)
	global.Logger.Info("GetMonthlyEarnings success", zap.String("info", userID))
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) GetDailyEarnings(ctx *gin.Context) {

	codeStatus, data, err := services.Stats().GetDailyEarnings(ctx)
	if err != nil {
		fmt.Printf("GetDailyEarnings error: %s\n", err.Error())
		global.Logger.Error("GetDailyEarnings error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	userID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		fmt.Printf("GetDailyEarnings cannot found userID")
		global.Logger.Error("GetDailyEarnings cannot found userID")
		userID = consts.UNKNOWN
	}

	fmt.Printf("GetDailyEarnings success: %s\n", userID)
	global.Logger.Info("GetDailyEarnings success", zap.String("info", userID))
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) GetDailyEarningsByMonth(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("GetDailyEarningsByMonth validation not found\n")
		global.Logger.Error("GetDailyEarningsByMonth validation not found")
		response.ErrorResponse(ctx, response.ErrCodeInternalServerError, nil)
		return
	}

	var params vo.GetDailyEarningsByMonthInput
	if err := ctx.ShouldBindUri(&params); err != nil {
		fmt.Printf("GetDailyEarningsByMonth binding error")
		global.Logger.Error("GetDailyEarningsByMonth binding error")
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		validationErrors := response.FormatValidationErrorsToStruct(err, params)
		fmt.Printf("GetDailyEarningsByMonth validation error: %s\n", validationErrors)
		global.Logger.Error("GetDailyEarningsByMonth validation error: ", zap.Any("error", validationErrors))
		response.ErrorResponse(ctx, response.ErrCodeValidator, validationErrors)
		return
	}

	codeStatus, data, err := services.Stats().GetDailyEarningsByMonth(ctx, &params)
	if err != nil {
		fmt.Printf("GetDailyEarningsByMonth error: %s\n", err.Error())
		global.Logger.Error("GetDailyEarningsByMonth error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("GetDailyEarningsByMonth success")
	global.Logger.Info("GetDailyEarningsByMonth success")
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) GetMonthlyEarningsByYear(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("GetMonthlyEarningsByYear validation not found\n")
		global.Logger.Error("GetMonthlyEarningsByYear validation not found")
		response.ErrorResponse(ctx, response.ErrCodeInternalServerError, nil)
		return
	}

	var params vo.GetMonthlyEarningsByYearInput
	if err := ctx.ShouldBindUri(&params); err != nil {
		fmt.Printf("GetMonthlyEarningsByYear binding error")
		global.Logger.Error("GetMonthlyEarningsByYear binding error")
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		validationErrors := response.FormatValidationErrorsToStruct(err, params)
		fmt.Printf("GetMonthlyEarningsByYear validation error: %s\n", validationErrors)
		global.Logger.Error("GetMonthlyEarningsByYear validation error: ", zap.Any("error", validationErrors))
		response.ErrorResponse(ctx, response.ErrCodeValidator, validationErrors)
		return
	}

	codeStatus, data, err := services.Stats().GetMonthlyEarningsByYear(ctx, &params)
	if err != nil {
		fmt.Printf("GetMonthlyEarningsByYear error: %s\n", err.Error())
		global.Logger.Error("GetMonthlyEarningsByYear error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("GetMonthlyEarningsByYear success")
	global.Logger.Info("GetMonthlyEarningsByYear success")
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) ExportDailyEarningsCSV(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("ExportDailyEarningsCSV validation not found\n")
		global.Logger.Error("ExportDailyEarningsCSV validation not found")
		response.ErrorResponse(ctx, response.ErrCodeInternalServerError, nil)
		return
	}

	var params vo.GetDailyEarningsByMonthInput
	if err := ctx.ShouldBindUri(&params); err != nil {
		fmt.Printf("ExportDailyEarningsCSV binding error")
		global.Logger.Error("ExportDailyEarningsCSV binding error")
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		validationErrors := response.FormatValidationErrorsToStruct(err, params)
		fmt.Printf("ExportDailyEarningsCSV validation error: %s\n", validationErrors)
		global.Logger.Error("ExportDailyEarningsCSV validation error: ", zap.Any("error", validationErrors))
		response.ErrorResponse(ctx, response.ErrCodeValidator, validationErrors)
		return
	}

	codeStatus, data, err := services.Stats().ExportDailyEarningsCSV(ctx, &params)
	if err != nil {
		fmt.Printf("ExportDailyEarningsCSV error: %s\n", err.Error())
		global.Logger.Error("ExportDailyEarningsCSV error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("ExportDailyEarningsCSV success")
	global.Logger.Info("ExportDailyEarningsCSV success")
	response.SuccessResponse(ctx, codeStatus, data)

	// Set headers for CSV download
	filename := fmt.Sprintf("daily_earnings_%d_%02d.csv", params.Year, params.Month)
	ctx.Header("Content-Type", "text/csv")
	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	ctx.Header("Content-Length", strconv.Itoa(len(data)))

	ctx.Data(http.StatusOK, "text/csv", data)
}

func (c *Controller) ExportMonthlyEarningsCSV(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("ExportMonthlyEarningsCSV validation not found\n")
		global.Logger.Error("ExportMonthlyEarningsCSV validation not found")
		response.ErrorResponse(ctx, response.ErrCodeInternalServerError, nil)
		return
	}

	var params vo.GetMonthlyEarningsByYearInput
	if err := ctx.ShouldBindUri(&params); err != nil {
		fmt.Printf("ExportMonthlyEarningsCSV binding error")
		global.Logger.Error("ExportMonthlyEarningsCSV binding error")
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		validationErrors := response.FormatValidationErrorsToStruct(err, params)
		fmt.Printf("ExportMonthlyEarningsCSV validation error: %s\n", validationErrors)
		global.Logger.Error("ExportMonthlyEarningsCSV validation error: ", zap.Any("error", validationErrors))
		response.ErrorResponse(ctx, response.ErrCodeValidator, validationErrors)
		return
	}

	codeStatus, data, err := services.Stats().ExportMonthlyEarningsCSV(ctx, &params)
	if err != nil {
		fmt.Printf("ExportMonthlyEarningsCSV error: %s\n", err.Error())
		global.Logger.Error("ExportMonthlyEarningsCSV error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("ExportMonthlyEarningsCSV success")
	global.Logger.Info("ExportMonthlyEarningsCSV success")
	filename := fmt.Sprintf("monthly_earnings_%d.csv", params.Year)
	ctx.Header("Content-Type", "text/csv")
	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	ctx.Header("Content-Length", strconv.Itoa(len(data)))

	ctx.Data(http.StatusOK, "text/csv", data)
}
