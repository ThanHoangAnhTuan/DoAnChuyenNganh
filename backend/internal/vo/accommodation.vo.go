package vo

type Facilities struct {
	WiFi         bool `json:"wifi"`
	AirCondition bool `json:"air_condition"`
	TV           bool `json:"tv"`
}

type PropertySurroundings struct {
	Restaurant bool `json:"restaurant"`
	Bar        bool `json:"bar"`
}

type CreateAccommodationInput struct {
	Name                 string               `json:"name" validate:"required"`
	Country              string               `json:"country" validate:"required"`
	City                 string               `json:"city" validate:"required"`
	District             string               `json:"district" validate:"required"`
	Description          string               `json:"description" validate:"required"`
	Facilities           Facilities           `json:"facilities" validate:"required"`
	GoogleMap            string               `json:"google_map" validate:"required"`
	PropertySurroundings PropertySurroundings `json:"property_surrounds" validate:"required"`
	Rules                string               `json:"rules" validate:"required"`
}

type CreateAccommodationOutput struct {
	Id                   string               `json:"id"`
	ManagerId            string               `json:"manager_id"`
	Name                 string               `json:"name"`
	City                 string               `json:"city"`
	Country              string               `json:"country"`
	District             string               `json:"district"`
	Images               []string             `json:"images"`
	Description          string               `json:"description"`
	Rating               string               `json:"rating"`
	Facilities           Facilities           `json:"facilities"`
	GoogleMap            string               `json:"google_map"`
	PropertySurroundings PropertySurroundings `json:"property_surrounds"`
	Rules                string               `json:"rules"`
}

type GetAccommodations struct {
	Id                   string               `json:"id"`
	ManagerId            string               `json:"manager_id"`
	Name                 string               `json:"name"`
	City                 string               `json:"city"`
	Country              string               `json:"country"`
	District             string               `json:"district"`
	Images               []string             `json:"images"`
	Description          string               `json:"description"`
	Rating               string               `json:"rating"`
	Facilities           Facilities           `json:"facilities"`
	GoogleMap            string               `json:"google_map"`
	PropertySurroundings PropertySurroundings `json:"property_surrounds"`
	Rules                string               `json:"rules"`
}

type UpdateAccommodationInput struct {
	Id                   string               `json:"id" validate:"required"`
	Name                 string               `json:"name"`
	Country              string               `json:"country"`
	City                 string               `json:"city"`
	District             string               `json:"district"`
	Description          string               `json:"description"`
	Facilities           Facilities           `json:"facilities"`
	GoogleMap            string               `json:"google_map"`
	PropertySurroundings PropertySurroundings `json:"property_surrounds"`
	Rules                string               `json:"rules"`
}

type UpdateAccommodationOutput struct {
	Id                   string               `json:"id"`
	ManagerId            string               `json:"manager_id"`
	Name                 string               `json:"name"`
	City                 string               `json:"city"`
	Country              string               `json:"country"`
	District             string               `json:"district"`
	Images               []string             `json:"images"`
	Description          string               `json:"description"`
	Rating               string               `json:"rating"`
	Facilities           Facilities           `json:"facilities"`
	GoogleMap            string               `json:"google_map"`
	PropertySurroundings PropertySurroundings `json:"property_surrounds"`
	Rules                string               `json:"rules"`
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
	Id        string `json:"id"`
	Name      string `json:"name"`
	City      string `json:"city"`
	Country   string `json:"country"`
	District  string `json:"district"`
	Image     string `json:"image"`
	Rating    string `json:"rating"`
	GoogleMap string `json:"google_map"`
}

// get accommodation by id

type GetAccommodationByIdInput struct {
	Id string `uri:"id"`
}

type GetAccommodationByIdOutput struct {
	Id                   string               `json:"id"`
	ManagerId            string               `json:"manager_id"`
	Name                 string               `json:"name"`
	City                 string               `json:"city"`
	Country              string               `json:"country"`
	District             string               `json:"district"`
	Images               []string             `json:"images"`
	Description          string               `json:"description"`
	Rating               string               `json:"rating"`
	Facilities           Facilities           `json:"facilities"`
	GoogleMap            string               `json:"google_map"`
	PropertySurroundings PropertySurroundings `json:"property_surrounds"`
	Rules                string               `json:"rules"`
}
