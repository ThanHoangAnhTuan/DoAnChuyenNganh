package vo

type CreateFacilityDetailInput struct {
	Name string `json:"name" validate:"required"`
}
type CreateFacilityDetailOutput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type GetFacilityDetailInput struct {
}
type GetFacilityDetailOutput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
