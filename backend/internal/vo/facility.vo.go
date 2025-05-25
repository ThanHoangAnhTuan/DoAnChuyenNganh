package vo

import "mime/multipart"

type CreateFacilityInput struct {
	Name  string                `form:"name" validate:"required"`
	Image *multipart.FileHeader `form:"image" validate:"required"`
}
type CreateFacilityOutput struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}
type GetFacilitiesInput struct {
}
type GetFacilitiesOutput struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}
