package vo

import "mime/multipart"

type Beds struct {
	SingleBed           uint `json:"single_bed"`
	DoubleBed           uint `json:"double_bed"`
	LargeDoubleBed      uint `json:"large_double_bed"`
	ExtraLargeDoubleBed uint `json:"extra_large_double_bed"`
}

type FacilitiesAccommodationDetail struct {
	WiFi         bool `json:"wifi"`
	AirCondition bool `json:"air_condition"`
	TV           bool `json:"tv"`
}

type GetAccommodationDetailsInput struct {
	AccommodationId string `json:"accommodation_id"`
}

type CreateAccommodationDetailInput struct {
	AccommodationId string                        `form:"accommodation_id" validate:"required"`
	Name            string                        `form:"name" validate:"required"`
	Guests          uint8                         `form:"guests" validate:"required"`
	Beds            Beds                          `form:"beds" validate:"required"`
	Facilities      FacilitiesAccommodationDetail `form:"facilities" validate:"required"`
	AvailableRooms  uint8                         `form:"available_rooms" validate:"required"`
	Price           string                        `form:"price" validate:"required"`
	DiscountId      string                        `form:"discount_id"`
	Images          []*multipart.FileHeader       `form:"images" validate:"required"`
}

type CreateAccommodationDetailOutput struct {
	Id              string                        `json:"id"`
	AccommodationId string                        `json:"accommodation_id"`
	Name            string                        `json:"name"`
	Guests          uint8                         `json:"guests"`
	Beds            Beds                          `json:"beds"`
	Facilities      FacilitiesAccommodationDetail `json:"facilities"`
	AvailableRooms  uint8                         `json:"available_rooms"`
	Price           string                        `json:"price"`
	DiscountId      string                        `json:"discount_id"`
	Images          []string                      `json:"images"`
}

type GetAccommodationDetailsOutput struct {
	Id              string                        `json:"id"`
	AccommodationId string                        `json:"accommodation_id"`
	Name            string                        `json:"name"`
	Guests          uint8                         `json:"guests"`
	Beds            Beds                          `json:"beds"`
	Facilities      FacilitiesAccommodationDetail `json:"facilities"`
	AvailableRooms  uint8                         `json:"available_rooms"`
	Price           string                        `json:"price"`
	DiscountId      string                        `json:"discount_id"`
	Images          []string                      `json:"images"`
}

// newvalue -> update
// "" -> remove
// null -> no action
type UpdateAccommodationDetailInput struct {
	Id              string                        `form:"id" validate:"required"`
	AccommodationId string                        `form:"accommodation_id"`
	Name            string                        `form:"name"`
	Guests          uint8                         `form:"guests"`
	Beds            Beds                          `form:"beds"`
	Facilities      FacilitiesAccommodationDetail `form:"facilities"`
	AvailableRooms  uint8                         `form:"available_rooms"`
	Price           string                        `form:"price"`
	DiscountId      string                        `form:"discount_id"`
	Images          []*multipart.FileHeader       `form:"images"`     //len(Images) == 0; len(Images) > 0
	OldImages       []string                      `form:"old_images"` // len(Images) == 0 || Images == nil; len(Images) > 0
}

type UpdateAccommodationDetailOutput struct {
	Id              string                        `json:"id"`
	AccommodationId string                        `json:"accommodation_id"`
	Name            string                        `json:"name"`
	Guests          uint8                         `json:"guests"`
	Beds            Beds                          `json:"beds"`
	Facilities      FacilitiesAccommodationDetail `json:"facilities"`
	AvailableRooms  uint8                         `json:"available_rooms"`
	Price           string                        `json:"price"`
	DiscountId      string                        `json:"discount_id"`
	Images          []string                      `json:"images"`
}

type DeleteAccommodationDetailInput struct {
	Id              string `json:"id" validate:"required"`
	AccommodationId string `json:"accommodation_id" validate:"required"`
}
