package vo

type CreateReviewInput struct {
	AccommodationID string `json:"accommodation_id" validate:"required"`
	Title           string `json:"title" validate:"required"`
	Comment         string `json:"comment" validate:"required"`
	Rating          uint8  `json:"rating" validate:"required"`
	OrderIDExternal string `json:"order_id" validate:"required"`
}
type CreateReviewOutput struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Image   string `json:"image"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Rating  uint8  `json:"rating"`
}

type GetReviewsInput struct {
	AccommodationID string `form:"accommodation_id" validate:"required"`
	BasePaginationInput
}

type GetReviewOutput struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Image           string `json:"image"`
	Title           string `json:"title"`
	Comment         string `json:"comment"`
	ManagerResponse string `json:"manager_response"`
	Rating          uint8  `json:"rating"`
}
