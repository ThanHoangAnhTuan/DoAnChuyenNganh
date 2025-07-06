package vo

type CreateAccommodationRoomInput struct {
	AccommodationTypeID string `json:"accommodation_type_id"`
	Prefix              string `json:"prefix"`
	Quantity            int    `json:"quantity"`
}

type CreateAccommodationRoomOutput struct {
	ID string `json:"id"`
	// AccommodationTypeID string `json:"accommodation_type_id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type GetAccommodationRoomsInput struct {
	AccommodationTypeID string `uri:"accommodation_type_id"`
}

type GetAccommodationRoomsOutput struct {
	ID string `json:"id"`
	// AccommodationTypeID string `json:"accommodation_type_id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type RoomStatus string

const (
	StatusAvailable   RoomStatus = "available"
	StatusUnavailable RoomStatus = "unavailable"
	StatusOccupied    RoomStatus = "occupied"
)

type UpdateAccommodationRoomInput struct {
	ID     string     `json:"id"`
	Name   string     `json:"name"`
	Status RoomStatus `json:"status"`
}

type UpdateAccommodationRoomOutput struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type DeleteAccommodationRoomInput struct {
	ID string `uri:"id" validate:"required"`
}
