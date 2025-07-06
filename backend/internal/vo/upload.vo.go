package vo

import "mime/multipart"

type UploadImages struct {
	Images       []*multipart.FileHeader `form:"images"`
	DeleteImages []string                `form:"delete_images"`
	ID           string                  `form:"id"`
	IsDetail     bool                    `form:"is_detail"`
}

type GetImagesInput struct {
	ID       string `uri:"id" binding:"required"`
	IsDetail bool   `form:"is_detail" binding:"omitempty"`
}
