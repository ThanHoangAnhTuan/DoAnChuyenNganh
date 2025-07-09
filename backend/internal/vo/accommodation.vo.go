package vo

import "encoding/json"

type FacilitiesOutput struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

type Rule struct {
	CheckIn                 string `json:"check_in"`
	CheckOut                string `json:"check_out"`
	Cancellation            string `json:"cancellation"`
	RefundableDamageDeposit uint32 `json:"refundable_damage_deposit"`
	AgeRestriction          bool   `json:"age_restriction"`
	Pet                     bool   `json:"pet"`
}

type CreateAccommodationInput struct {
	Name        string   `json:"name" validate:"required"`
	Country     string   `json:"country" validate:"required"`
	City        string   `json:"city" validate:"required"`
	District    string   `json:"district" validate:"required"`
	Address     string   `json:"address" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Facilities  []string `json:"facilities" validate:"required"`
	GoogleMap   string   `json:"google_map" validate:"required"`
	Rating      uint8    `json:"rating" validate:"required"`
	Rules       Rule     `json:"rules"`
}

type CreateAccommodationOutput struct {
	ID          string             `json:"id"`
	ManagerID   string             `json:"manager_id"`
	Name        string             `json:"name"`
	City        string             `json:"city"`
	Country     string             `json:"country"`
	District    string             `json:"district"`
	Address     string             `json:"address"`
	Images      []string           `json:"images"`
	Description string             `json:"description"`
	Rating      uint8              `json:"rating"`
	Facilities  []FacilitiesOutput `json:"facilities"`
	GoogleMap   string             `json:"google_map"`
	Rules       Rule               `json:"rules"`
	IsVerified  bool               `json:"is_verified"`
	IsDeleted   bool               `json:"is_deleted"`
}

type GetAccommodationsInput struct {
	City string `form:"city"`
	BasePaginationInput
}

type GetAccommodationsOutput struct {
	ID          string             `json:"id"`
	ManagerID   string             `json:"manager_id"`
	Name        string             `json:"name"`
	City        string             `json:"city"`
	Country     string             `json:"country"`
	District    string             `json:"district"`
	Address     string             `json:"address"`
	Images      []string           `json:"images"`
	Description string             `json:"description"`
	Rating      uint8              `json:"rating"`
	Facilities  []FacilitiesOutput `json:"facilities"`
	GoogleMap   string             `json:"google_map"`
	Rules       Rule               `json:"rules"`
	IsVerified  bool               `json:"is_verified"`
	IsDeleted   bool               `json:"is_deleted"`
}

type UpdateAccommodationInput struct {
	ID          string   `json:"id" validate:"required"`
	Name        string   `json:"name"`
	Country     string   `json:"country"`
	City        string   `json:"city"`
	District    string   `json:"district"`
	Address     string   `json:"address"`
	Description string   `json:"description"`
	Facilities  []string `json:"facilities"`
	GoogleMap   string   `json:"google_map"`
	Rules       Rule     `json:"rules"`
}

type UpdateAccommodationOutput struct {
	ID          string             `json:"id"`
	ManagerID   string             `json:"manager_id"`
	Name        string             `json:"name"`
	City        string             `json:"city"`
	Country     string             `json:"country"`
	District    string             `json:"district"`
	Address     string             `json:"address"`
	Images      []string           `json:"images"`
	Description string             `json:"description"`
	Rating      uint8              `json:"rating"`
	Facilities  []FacilitiesOutput `json:"facilities"`
	GoogleMap   string             `json:"google_map"`
	Rules       Rule               `json:"rules"`
}

type DeleteAccommodationInput struct {
	ID string `json:"id" validate:"required"`
}

// get accommodation by city
type GetAccommodationByCityInput struct {
	City string `uri:"city"`
	BasePaginationInput
}

type GetAccommodationsByCityOutput struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	City      string   `json:"city"`
	Country   string   `json:"country"`
	District  string   `json:"district"`
	Address   string   `json:"address"`
	Images    []string `json:"images"`
	Rating    uint8    `json:"rating"`
	GoogleMap string   `json:"google_map"`
}

type GetAccommodationInput struct {
	ID string `uri:"id"`
}

type GetAccommodationOutput struct {
	ID          string             `json:"id"`
	ManagerID   string             `json:"manager_id"`
	Name        string             `json:"name"`
	City        string             `json:"city"`
	Country     string             `json:"country"`
	District    string             `json:"district"`
	Address     string             `json:"address"`
	Images      []string           `json:"images"`
	Description string             `json:"description"`
	Rating      uint8              `json:"rating"`
	Facilities  []FacilitiesOutput `json:"facilities"`
	GoogleMap   string             `json:"google_map"`
	Rules       Rule               `json:"rules"`
}

type AccommodationData struct {
	ID          string
	ManagerID   string
	Country     string
	Name        string
	City        string
	District    string
	Description string
	Facilities  json.RawMessage
	Address     string
	GgMap       string
	Rules       json.RawMessage
	Rating      uint8
	IsVerified  bool
	IsDeleted   bool
}
