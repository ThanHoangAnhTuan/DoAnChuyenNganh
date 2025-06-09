package vo

type Beds struct {
	SingleBed           uint `json:"single_bed"`
	DoubleBed           uint `json:"double_bed"`
	LargeDoubleBed      uint `json:"large_double_bed"`
	ExtraLargeDoubleBed uint `json:"extra_large_double_bed"`
}
type GetAccommodationDetailsInput struct {
	AccommodationID string `json:"accommodation_id"`
}

type CreateAccommodationDetailInput struct {
	AccommodationID string   `json:"accommodation_id" validate:"required"`
	Name            string   `json:"name" validate:"required"`
	Guests          uint8    `json:"guests" validate:"gte=1"`
	Beds            Beds     `json:"beds" validate:"required"`
	Facilities      []string `json:"facilities"`
	AvailableRooms  uint8    `json:"available_rooms" validate:"gte=0"`
	Price           string   `json:"price" validate:"gte=1"`
	DiscountID      string   `json:"discount_id"`
}

type FacilityDetailOutput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CreateAccommodationDetailOutput struct {
	ID              string                 `json:"id"`
	AccommodationID string                 `json:"accommodation_id"`
	Name            string                 `json:"name"`
	Guests          uint8                  `json:"guests"`
	Beds            Beds                   `json:"beds"`
	Facilities      []FacilityDetailOutput `json:"facilities"`
	AvailableRooms  uint8                  `json:"available_rooms"`
	Price           string                 `json:"price"`
	DiscountID      string                 `json:"discount_id"`
	Images          []string               `json:"images"`
}

type GetAccommodationDetailsOutput struct {
	ID              string                 `json:"id"`
	AccommodationID string                 `json:"accommodation_id"`
	Name            string                 `json:"name"`
	Guests          uint8                  `json:"guests"`
	Beds            Beds                   `json:"beds"`
	Facilities      []FacilityDetailOutput `json:"facilities"`
	AvailableRooms  uint8                  `json:"available_rooms"`
	Price           string                 `json:"price"`
	DiscountID      string                 `json:"discount_id"`
	Images          []string               `json:"images"`
}

type UpdateAccommodationDetailInput struct {
	ID              string   `json:"id" validate:"required"`
	AccommodationID string   `json:"accommodation_id"`
	Name            string   `json:"name"`
	Guests          uint8    `json:"guests"`
	Beds            Beds     `json:"beds"`
	Facilities      []string `json:"facilities"`
	AvailableRooms  uint8    `json:"available_rooms"`
	Price           string   `json:"price"`
	DiscountID      string   `json:"discount_id"`
}

type UpdateAccommodationDetailOutput struct {
	ID              string                 `json:"id"`
	AccommodationID string                 `json:"accommodation_id"`
	Name            string                 `json:"name"`
	Guests          uint8                  `json:"guests"`
	Beds            Beds                   `json:"beds"`
	Facilities      []FacilityDetailOutput `json:"facilities"`
	AvailableRooms  uint8                  `json:"available_rooms"`
	Price           string                 `json:"price"`
	DiscountID      string                 `json:"discount_id"`
	Images          []string               `json:"images"`
}

type DeleteAccommodationDetailInput struct {
	ID string `json:"id" validate:"required"`
}
