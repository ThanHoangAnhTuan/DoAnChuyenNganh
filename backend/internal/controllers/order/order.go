package order

import (
	"fmt"

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

func (c *Controller) CancelOrder(ctx *gin.Context) {
	var params vo.CancelOrderInput
	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.CancelOrderInput) error {
		return ctx.ShouldBind(p)
	}); err != nil {
		return
	}

	codeStatus, err := services.Order().CancelOrder(ctx, &params)
	if err != nil {
		fmt.Printf("CancelOrder error: %s\n", err.Error())
		global.Logger.Error("CancelOrder error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("CancelOrder success: orderID: %s\n", params.OrderID)
	global.Logger.Info("CancelOrder success", zap.String("orderID", params.OrderID))

	response.SuccessResponse(ctx, codeStatus, nil)
}

func (c *Controller) CheckIn(ctx *gin.Context) {
	var params vo.CheckInInput
	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.CheckInInput) error {
		return ctx.ShouldBind(p)
	}); err != nil {
		return
	}

	codeStatus, err := services.Order().CheckIn(ctx, &params)
	if err != nil {
		fmt.Printf("CheckIn error: %s\n", err.Error())
		global.Logger.Error("CheckIn error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("CheckIn success: orderID: %s\n", params.OrderID)
	global.Logger.Info("CheckIn success", zap.String("orderID", params.OrderID))

	response.SuccessResponse(ctx, codeStatus, nil)
}

func (c *Controller) CheckOut(ctx *gin.Context) {
	var params vo.CheckOutInput
	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.CheckOutInput) error {
		return ctx.ShouldBind(p)
	}); err != nil {
		return
	}

	codeStatus, err := services.Order().CheckOut(ctx, &params)
	if err != nil {
		fmt.Printf("CheckOut error: %s\n", err.Error())
		global.Logger.Error("CheckOut error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("CheckOut success: orderID: %s\n", params.OrderID)
	global.Logger.Info("CheckOut success", zap.String("orderID", params.OrderID))

	response.SuccessResponse(ctx, codeStatus, nil)
}

func (c *Controller) GetOrderInfoAfterPayment(ctx *gin.Context) {
	var params vo.GetOrderInfoAfterPaymentInput
	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.GetOrderInfoAfterPaymentInput) error {
		return ctx.ShouldBind(p)
	}); err != nil {
		return
	}

	codeStatus, data, err := services.Order().GetOrderInfoAfterPayment(ctx, &params)
	if err != nil {
		fmt.Printf("GetOrderInfoAfterPayment error: %s\n", err.Error())
		global.Logger.Error("GetOrderInfoAfterPayment error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("GetOrderInfoAfterPayment success: orderID: %s\n", data.OrderIDExternal)
	global.Logger.Info("GetOrderInfoAfterPayment success",
		zap.String("orderID", data.OrderIDExternal))

	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) GetOrdersByManager(ctx *gin.Context) {
	codeStatus, data, err := services.Order().GetOrdersByManager(ctx)
	if err != nil {
		fmt.Printf("GetOrdersByManager error: %s\n", err.Error())
		global.Logger.Error("GetOrdersByManager error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	// TODO: get managerID
	managerID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		fmt.Printf("GetOrdersByManager cannot found managerID")
		global.Logger.Error("GetOrdersByManager cannot found managerID")
		managerID = consts.UNKNOWN
	}

	fmt.Printf("GetOrdersByManager success: %s\n", managerID)
	global.Logger.Info("GetOrdersByManager success", zap.String("info", managerID))
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) GetOrdersByUser(ctx *gin.Context) {
	codeStatus, data, err := services.Order().GetOrdersByUser(ctx)
	if err != nil {
		fmt.Printf("GetOrdersByUser error: %s\n", err.Error())
		global.Logger.Error("GetOrdersByUser error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	// TODO: get managerID
	managerID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		fmt.Printf("GetOrdersByUser cannot found managerID")
		global.Logger.Error("GetOrdersByUser cannot found managerID")
		managerID = consts.UNKNOWN
	}

	fmt.Printf("GetOrdersByUser success: %s\n", managerID)
	global.Logger.Info("GetOrdersByUser success", zap.String("info", managerID))
	response.SuccessResponse(ctx, codeStatus, data)
}
