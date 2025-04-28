package vo

import "mime/multipart"

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
	Name                 string                  `form:"name"`
	Country              string                  `form:"country"`
	City                 string                  `form:"city"`
	District             string                  `form:"district"`
	Image                []*multipart.FileHeader `form:"image"`
	Description          string                  `form:"description"`
	Facilities           Facilities              `form:"-"`
	GoogleMap            string                  `form:"google_map"`
	PropertySurroundings PropertySurroundings    `form:"property_surrounds"`
	Rules                string                  `form:"rules"`
}

type CreateAccommodationOutput struct {
	Id                   string               `json:"id"`
	ManagerId            string               `json:"manager_id"`
	Name                 string               `json:"name"`
	City                 string               `json:"city"`
	Country              string               `json:"country"`
	District             string               `json:"district"`
	Image                string               `json:"image"`
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
	Image                string               `json:"image"`
	Description          string               `json:"description"`
	Rating               string               `json:"rating"`
	Facilities           Facilities           `json:"facilities"`
	GoogleMap            string               `json:"google_map"`
	PropertySurroundings PropertySurroundings `json:"property_surrounds"`
	Rules                string               `json:"rules"`
}

type UpdateAccommodationInput struct {
	Id                   string                  `form:"id"`
	Name                 string                  `form:"name"`
	Country              string                  `form:"country"`
	City                 string                  `form:"city"`
	District             string                  `form:"district"`
	Image                []*multipart.FileHeader `form:"image"`
	Description          string                  `form:"description"`
	Facilities           Facilities              `form:"-"`
	GoogleMap            string                  `form:"google_map"`
	PropertySurroundings PropertySurroundings    `form:"property_surrounds"`
	Rules                string                  `form:"rules"`
}

type UpdateAccommodationOutput struct {
	Id                   string               `json:"id"`
	ManagerId            string               `json:"manager_id"`
	Name                 string               `json:"name"`
	City                 string               `json:"city"`
	Country              string               `json:"country"`
	District             string               `json:"district"`
	Image                string               `json:"image"`
	Description          string               `json:"description"`
	Rating               string               `json:"rating"`
	Facilities           Facilities           `json:"facilities"`
	GoogleMap            string               `json:"google_map"`
	PropertySurroundings PropertySurroundings `json:"property_surrounds"`
	Rules                string               `json:"rules"`
}

type DeleteAccommodationInput struct {
	Id string `json:"id"`
}
