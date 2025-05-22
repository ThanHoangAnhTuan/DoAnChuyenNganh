package vo

import "mime/multipart"

type UploadImages struct {
	Images    []*multipart.FileHeader `form:"images"`
	OldImages []string                `form:"old_images"`
	Id        string                  `form:"id"`
	IsDetail  bool                    `form:"is_detail"`
}

type GetImagesInput struct {
	Id       string `uri:"id" binding:"required"`
	IsDetail bool   `form:"is_detail" binding:"omitempty"`
}
