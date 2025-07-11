package stats

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/global"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/consts"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/controllerutil"
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
	var params vo.GetDailyEarningsByMonthInput
	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.GetDailyEarningsByMonthInput) error {
		return ctx.ShouldBindUri(p)
	}); err != nil {
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
	var params vo.GetMonthlyEarningsByYearInput
	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.GetMonthlyEarningsByYearInput) error {
		return ctx.ShouldBindUri(p)
	}); err != nil {
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
	var params vo.GetDailyEarningsByMonthInput
	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.GetDailyEarningsByMonthInput) error {
		return ctx.ShouldBindUri(p)
	}); err != nil {
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

	// Set headers for CSV download
	filename := fmt.Sprintf("daily_earnings_%d_%02d.csv", params.Year, params.Month)
	ctx.Header("Content-Type", "text/csv")
	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	ctx.Header("Content-Length", strconv.Itoa(len(data)))

	ctx.Data(http.StatusOK, "text/csv", data)
}

func (c *Controller) ExportMonthlyEarningsCSV(ctx *gin.Context) {
	var params vo.GetMonthlyEarningsByYearInput
	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.GetMonthlyEarningsByYearInput) error {
		return ctx.ShouldBindUri(p)
	}); err != nil {
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
