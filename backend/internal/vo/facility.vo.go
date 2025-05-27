package vo

import "mime/multipart"

type CreateFacilityInput struct {
	Name  string                `form:"name" validate:"required"`
	Image *multipart.FileHeader `form:"image" validate:"required"`
}
type CreateFacilityOutput struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}
type GetFacilitiesInput struct {
}
type GetFacilitiesOutput struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}
