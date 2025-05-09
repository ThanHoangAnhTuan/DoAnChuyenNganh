package vo

import "mime/multipart"

type UploadImages struct {
	Images              []*multipart.FileHeader `form:"images"`
	DeleteImages        []string                `form:"delete_images"`
	Accommodation       string                  `form:"accommodation_id"`
	AccommodationDetail string                  `form:"accommodation_detail_id"`
}
