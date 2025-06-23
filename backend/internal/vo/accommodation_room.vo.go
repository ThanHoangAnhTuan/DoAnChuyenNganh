package vo

type CreateAccommodationRoomInput struct {
	AccommodationTypeID string `json:"accommodation_type_id"`
	Prefix              string `json:"prefix"`
	Quantity            int    `json:"quantity"`
}

type CreateAccommodationRoomOutput struct {
	ID                  string `json:"id"`
	AccommodationTypeID string `json:"accommodation_type_id"`
	Name                string `json:"name"`
	Status              string `json:"status"`
}

type GetAccommodationRoomsInput struct {
	AccommodationTypeID string `uri:"accommodation_type_id"`
}

type GetAccommodationRoomsOutput struct {
	ID                  string `json:"id"`
	AccommodationTypeID string `json:"accommodation_type_id"`
	Name                string `json:"name"`
	Status              string `json:"status"`
}

type UpdateAccommodationRoomInput struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type UpdateAccommodationRoomOutput struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type DeleteAccommodationRoomInput struct {
	ID string `uri:"id" validate:"required"`
}
