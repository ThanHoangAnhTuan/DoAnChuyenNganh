package vo

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
}

type GetAccommodations struct {
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
}

// get accommodation by city
type GetAccommodationsByCity struct {
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

// get accommodation by id

type GetAccommodationByIdInput struct {
	ID string `uri:"id"`
}

type GetAccommodationByIdOutput struct {
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
