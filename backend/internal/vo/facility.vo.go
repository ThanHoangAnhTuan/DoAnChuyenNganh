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

type UpdateFacilityInput struct {
	ID    string                `form:"id" validate:"required"`
	Name  string                `form:"name" validate:"required"`
	Image *multipart.FileHeader `form:"image"`
}
type UpdateFacilityOutput struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

type DeleteFacilityInput struct {
	ID string `uri:"id"`
}
