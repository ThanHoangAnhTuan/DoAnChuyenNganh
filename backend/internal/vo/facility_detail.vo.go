package vo

type CreateFacilityDetailInput struct {
	Name string `form:"name" validate:"required"`
}
type CreateFacilityDetailOutput struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
type GetFacilityDetailInput struct {
}
type GetFacilityDetailOutput struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
