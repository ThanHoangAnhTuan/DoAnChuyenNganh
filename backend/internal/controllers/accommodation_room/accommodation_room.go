package accommodation_room

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

func (c *Controller) CreateAccommodationRoom(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("CreateAccommodationRoom validation not found\n")
		global.Logger.Error("CreateAccommodationRoom validation not found")
		response.ErrorResponse(ctx, response.ErrCodeInternalServerError, nil)
		return
	}

	var params vo.CreateAccommodationRoomInput
	if err := ctx.ShouldBind(&params); err != nil {
		fmt.Printf("CreateAccommodationRoom binding error: %s\n", err.Error())
		global.Logger.Error("CreateAccommodationRoom binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeValidator, err)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		validationErrors := response.FormatValidationErrorsToStruct(err, params)
		fmt.Printf("CreateAccommodationRoom validation error: %s\n", validationErrors)
		global.Logger.Error("CreateAccommodationRoom validation error: ", zap.Any("error", validationErrors))
		response.ErrorResponse(ctx, response.ErrCodeValidator, validationErrors)
		return
	}

	codeStatus, data, err := services.AccommodationRoom().CreateAccommodationRoom(ctx, &params)
	if err != nil {
		fmt.Printf("CreateAccommodationRoom error: %s\n", err.Error())
		global.Logger.Error("CreateAccommodationRoom error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, err)
		return
	}

	userID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		fmt.Printf("CreateAccommodationRoom cannot found userID")
		global.Logger.Error("CreateAccommodationRoom cannot found userID")
		userID = consts.UNKNOWN
	}

	fmt.Printf("CreateAccommodationRoom success: manager: %s, AccommodationTypeID: %s\n", userID, params.AccommodationTypeID)
	global.Logger.Info("CreateAccommodationRoom success",
		zap.String("managerId", userID),
		zap.String("AccommodationTypeID", params.AccommodationTypeID))
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) GetAccommodationRooms(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("GetAccommodationRooms validation not found\n")
		global.Logger.Error("GetAccommodationRooms validation not found")
		response.ErrorResponse(ctx, response.ErrCodeInternalServerError, nil)
		return
	}

	var params vo.GetAccommodationRoomsInput
	if err := ctx.ShouldBindUri(&params); err != nil {
		fmt.Printf("GetAccommodationRooms binding error")
		global.Logger.Error("GetAccommodationRooms binding error")
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		validationErrors := response.FormatValidationErrorsToStruct(err, params)
		fmt.Printf("GetAccommodationRooms validation error: %s\n", validationErrors)
		global.Logger.Error("GetAccommodationRooms validation error: ", zap.Any("error", validationErrors))
		response.ErrorResponse(ctx, response.ErrCodeValidator, validationErrors)
		return
	}

	codeStatus, data, err := services.AccommodationRoom().GetAccommodationRooms(ctx, &params)
	if err != nil {
		fmt.Printf("GetAccommodationRooms error: %s\n", err.Error())
		global.Logger.Error("GetAccommodationRooms error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("GetAccommodationRooms success: %s", params.AccommodationTypeID)
	global.Logger.Info("GetAccommodationRooms success", zap.String("AccommodationRoomID", params.AccommodationTypeID))
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) UpdateAccommodationRoom(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("UpdateAccommodationRoom validation not found \n")
		global.Logger.Error("UpdateAccommodationRoom validation not found")
		response.ErrorResponse(ctx, response.ErrCodeInternalServerError, nil)
		return
	}

	var params vo.UpdateAccommodationRoomInput
	if err := ctx.ShouldBind(&params); err != nil {
		fmt.Printf("UpdateAccommodationRoom binding error: %s\n", err.Error())
		global.Logger.Error("UpdateAccommodationRoom binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		validationErrors := response.FormatValidationErrorsToStruct(err, params)
		fmt.Printf("UpdateAccommodationRoomInput validation error: %s\n", validationErrors)
		global.Logger.Error("UpdateAccommodationRoomInput validation error: ", zap.Any("error", validationErrors))
		response.ErrorResponse(ctx, response.ErrCodeValidator, validationErrors)
		return
	}

	codeStatus, data, err := services.AccommodationRoom().UpdateAccommodationRoom(ctx, &params)
	if err != nil {
		fmt.Printf("UpdateAccommodationRoom error: %s\n", err.Error())
		global.Logger.Error("UpdateAccommodationRoom error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	userID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		fmt.Printf("UpdateAccommodationRoom cannot found userID")
		global.Logger.Error("UpdateAccommodationRoom cannot found userID")
		userID = "unknown"
	}

	fmt.Printf("UpdateAccommodationRoom success: manager: %s, AccommodationRoomID: %s\n", userID, data.ID)
	global.Logger.Info("UpdateAccommodationRoom success",
		zap.String("managerId", userID),
		zap.String("AccommodationRoomID", data.ID))

	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) DeleteAccommodationRoom(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("DeleteAccommodationRoom validation not found\n")
		global.Logger.Error("DeleteAccommodationRoom validation not found")
		response.ErrorResponse(ctx, response.ErrCodeInternalServerError, nil)
		return
	}

	var params vo.DeleteAccommodationRoomInput
	if err := ctx.ShouldBindUri(&params); err != nil {
		fmt.Printf("DeleteAccommodationRoom binding error: %s\n", err.Error())
		global.Logger.Error("DeleteAccommodationRoom binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		validationErrors := response.FormatValidationErrorsToStruct(err, params)
		fmt.Printf("DeleteAccommodationRoom validation error: %s\n", validationErrors)
		global.Logger.Error("DeleteAccommodationRoom validation error: ", zap.Any("error", validationErrors))
		response.ErrorResponse(ctx, response.ErrCodeValidator, validationErrors)
		return
	}

	codeStatus, err := services.AccommodationRoom().DeleteAccommodationRoom(ctx, &params)
	if err != nil {
		fmt.Printf("DeleteAccommodationRoom error: %s\n", err.Error())
		global.Logger.Error("DeleteAccommodationRoom error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	userID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		fmt.Printf("CreateAccommodationRoom cannot found userID")
		global.Logger.Error("CreateAccommodationRoom cannot found userID")
		userID = "unknown"
	}

	fmt.Printf("DeleteAccommodationRoom success: manager: %s, AccommodationRoomID: %s\n", userID, params.ID)
	global.Logger.Info("DeleteAccommodationRoom success",
		zap.String("managerId", userID),
		zap.String("AccommodationRoomID", params.ID))
	response.SuccessResponse(ctx, codeStatus, nil)
}
