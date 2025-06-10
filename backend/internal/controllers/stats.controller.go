package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/global"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/response"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils"
	"go.uber.org/zap"
)

var Stats = new(CStats)

type CStats struct {
}

func (c *CStats) GetMonthlyEarnings(ctx *gin.Context) {
	codeStatus, data, err := services.Stats().GetMonthlyEarnings(ctx)
	if err != nil {
		fmt.Printf("GetMonthlyEarnings error: %s\n", err.Error())
		global.Logger.Error("GetMonthlyEarnings error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	userID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		fmt.Printf("CreateAccommodationDetail cannot found userID")
		global.Logger.Error("CreateAccommodationDetail cannot found userID")
		userID = "unknown"
	}

	fmt.Printf("GetMonthlyEarnings success: %s\n", userID)
	global.Logger.Info("GetMonthlyEarnings success", zap.String("info", userID))
	response.SuccessResponse(ctx, codeStatus, data)
}
