package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/global"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/services"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/vo"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/pkg/response"
	"go.uber.org/zap"
)

var UploadImages = new(CUploadImages)

type CUploadImages struct {
}

func (c *CUploadImages) UploadImages(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("Validation not found")
		global.Logger.Error("Validation not found")
		response.ErrorResponse(ctx, response.ErrCodeValidatorNotFound, nil)
		return
	}

	var params vo.UploadImages
	if err := ctx.ShouldBind(&params); err != nil {
		fmt.Printf("UploadImages binding error: %s\n", err.Error())
		global.Logger.Error("UploadImages binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		fmt.Printf("UploadImages validation error: %s\n", err.Error())
		global.Logger.Error("UploadImages validation error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeValidator, err.Error())
		return
	}

	codeStatus, data, err := services.UploadImage().UploadImages(ctx, &params)
	if err != nil {
		fmt.Printf("UploadImages error: %s\n", err.Error())
		global.Logger.Error("UploadImages error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("UploadImages success")
	global.Logger.Info("UploadImages success")
	response.SuccessResponse(ctx, codeStatus, data)
}
