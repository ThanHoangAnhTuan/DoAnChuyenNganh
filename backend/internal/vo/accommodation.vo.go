package vo

type Rule struct {
	CheckIn                 string `json:"check_in" validate:"required"`
	CheckOut                string `json:"check_out" validate:"required"`
	Cancellation            string `json:"cancellation" validate:"required"`
	RefundableDamageDeposit uint32 `json:"refundable_damage_deposit" validate:"required"`
	AgeRestriction          bool   `json:"age_restriction" validate:"boolean"`
	Pet                     bool   `json:"pet" validate:"boolean"`
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
	Rules       Rule     `json:"rules" validate:"required"`
}

type CreateAccommodationOutput struct {
	Id          string   `json:"id"`
	ManagerId   string   `json:"manager_id"`
	Name        string   `json:"name"`
	City        string   `json:"city"`
	Country     string   `json:"country"`
	District    string   `json:"district"`
	Address     string   `json:"address"`
	Images      []string `json:"images"`
	Description string   `json:"description"`
	Rating      uint8    `json:"rating"`
	Facilities  []string `json:"facilities"`
	GoogleMap   string   `json:"google_map"`
	Rules       Rule     `json:"rules"`
}

type GetAccommodations struct {
	Id          string   `json:"id"`
	ManagerId   string   `json:"manager_id"`
	Name        string   `json:"name"`
	City        string   `json:"city"`
	Country     string   `json:"country"`
	District    string   `json:"district"`
	Address     string   `json:"address"`
	Images      []string `json:"images"`
	Description string   `json:"description"`
	Rating      uint8    `json:"rating"`
	Facilities  []string `json:"facilities"`
	GoogleMap   string   `json:"google_map"`
	Rules       Rule     `json:"rules"`
}

type UpdateAccommodationInput struct {
	Id          string   `json:"id" validate:"required"`
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
	Id          string   `json:"id"`
	ManagerId   string   `json:"manager_id"`
	Name        string   `json:"name"`
	City        string   `json:"city"`
	Country     string   `json:"country"`
	District    string   `json:"district"`
	Address     string   `json:"address"`
	Images      []string `json:"images"`
	Description string   `json:"description"`
	Rating      uint8    `json:"rating"`
	Facilities  []string `json:"facilities"`
	GoogleMap   string   `json:"google_map"`
	Rules       Rule     `json:"rules"`
}

type DeleteAccommodationInput struct {
	Id string `uri:"id" validate:"required"`
}

// get accommodation by city
type GetAccommodationByCityInput struct {
	City string `uri:"city"`
}

// get accommodation by city
type GetAccommodationsByCity struct {
	Id        string   `json:"id"`
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
	Id string `uri:"id"`
}

type GetAccommodationByIdOutput struct {
	Id          string   `json:"id"`
	ManagerId   string   `json:"manager_id"`
	Name        string   `json:"name"`
	City        string   `json:"city"`
	Country     string   `json:"country"`
	District    string   `json:"district"`
	Address     string   `json:"address"`
	Images      []string `json:"images"`
	Description string   `json:"description"`
	Rating      uint8    `json:"rating"`
	Facilities  []string `json:"facilities"`
	GoogleMap   string   `json:"google_map"`
	Rules       Rule     `json:"rules"`
}
