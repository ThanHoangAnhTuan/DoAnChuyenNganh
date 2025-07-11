package accommodation_room

import (
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

func (c *Controller) CreateAccommodationRoom(ctx *gin.Context) {
	start := utils.NowMs()
	var params vo.CreateAccommodationRoomInput

	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.CreateAccommodationRoomInput) error {
		return ctx.ShouldBind(p)
	}); err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", response.ErrCodeValidator, utils.NowMs()-start, err)
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}

	codeStatus, data, err := services.AccommodationRoom().CreateAccommodationRoom(ctx, &params)
	duration := utils.NowMs() - start

	if err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", codeStatus, duration, err)
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	userID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		global.Logger.Error("CreateAccommodationRoom cannot found userID")
		userID = consts.UNKNOWN
	}
	global.Logger.Info("CreateAccommodationRoom success",
		zap.String("managerId", userID),
		zap.String("AccommodationTypeID", params.AccommodationTypeID))
	controllerutil.HandleStructuredLog(ctx, params, "success", codeStatus, duration, nil)
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) GetAccommodationRooms(ctx *gin.Context) {
	start := utils.NowMs()
	var params vo.GetAccommodationRoomsInput

	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.GetAccommodationRoomsInput) error {
		return ctx.ShouldBindUri(p)
	}); err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", response.ErrCodeValidator, utils.NowMs()-start, err)
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}

	codeStatus, data, err := services.AccommodationRoom().GetAccommodationRooms(ctx, &params)
	duration := utils.NowMs() - start

	if err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", codeStatus, duration, err)
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	global.Logger.Info("GetAccommodationRooms success", zap.String("AccommodationRoomID", params.AccommodationTypeID))
	controllerutil.HandleStructuredLog(ctx, params, "success", codeStatus, duration, nil)
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) UpdateAccommodationRoom(ctx *gin.Context) {
	start := utils.NowMs()
	var params vo.UpdateAccommodationRoomInput

	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.UpdateAccommodationRoomInput) error {
		return ctx.ShouldBind(p)
	}); err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", response.ErrCodeValidator, utils.NowMs()-start, err)
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}

	codeStatus, data, err := services.AccommodationRoom().UpdateAccommodationRoom(ctx, &params)
	duration := utils.NowMs() - start

	if err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", codeStatus, duration, err)
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	userID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		global.Logger.Error("UpdateAccommodationRoom cannot found userID")
		userID = consts.UNKNOWN
	}
	global.Logger.Info("UpdateAccommodationRoom success",
		zap.String("managerId", userID),
		zap.String("AccommodationRoomID", data.ID))
	controllerutil.HandleStructuredLog(ctx, params, "success", codeStatus, duration, nil)
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) DeleteAccommodationRoom(ctx *gin.Context) {
	start := utils.NowMs()
	var params vo.DeleteAccommodationRoomInput

	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.DeleteAccommodationRoomInput) error {
		return ctx.ShouldBindUri(p)
	}); err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", response.ErrCodeValidator, utils.NowMs()-start, err)
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}

	codeStatus, err := services.AccommodationRoom().DeleteAccommodationRoom(ctx, &params)
	duration := utils.NowMs() - start

	if err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", codeStatus, duration, err)
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	userID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		global.Logger.Error("DeleteAccommodationRoom cannot found userID")
		userID = consts.UNKNOWN
	}
	global.Logger.Info("DeleteAccommodationRoom success",
		zap.String("managerId", userID),
		zap.String("AccommodationRoomID", params.ID))
	controllerutil.HandleStructuredLog(ctx, params, "success", codeStatus, duration, nil)
	response.SuccessResponse(ctx, codeStatus, nil)
}
