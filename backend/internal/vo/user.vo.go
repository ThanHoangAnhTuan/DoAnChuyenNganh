package vo

import (
	"encoding/json"
	"mime/multipart"
)

type RegisterInput struct {
	VerifyKey     string `json:"verify_key"`
	VerifyType    int    `json:"verify_type"`
	VerifyPurpose string `json:"verify_purpose"`
}

type VerifyOTPInput struct {
	VerifyKey  string `json:"verify_key"`
	VerifyCode string `json:"verify_code"`
}

type VerifyOTPOutput struct {
	Token string `json:"token"`
}

type UpdatePasswordRegisterInput struct {
	Token    string `json:"token"`
	Password string `json:"password"`
}

type UpdatePasswordRegisterOutput struct {
	UserId string `json:"user_id"`
}

type LoginInput struct {
	UserAccount  string `json:"account"`
	UserPassword string `json:"password"`
}

type LoginOutput struct {
	Token string `json:"token"`
}

type CreateAccommodationInput struct {
	Name                 string                `form:"name"`
	City                 string                `form:"city"`
	Provine              string                `form:"provine"`
	District             string                `form:"district"`
	Images               *multipart.FileHeader `form:"images"`
	Description          string                `form:"description"`
	Facilities           json.RawMessage       `form:"facilities"`
	GoogleMap            string                `form:"google_map"`
	PropertySurroundings json.RawMessage       `form:"property_surrounds"`
	Rules                string                `form:"rules"`
}

type CreateAccommodationOutput struct {
	Id                   string          `json:"id"`
	ManagerId            string          `json:"manager_id"`
	Name                 string          `json:"name"`
	City                 string          `json:"city"`
	Provine              string          `json:"provine"`
	District             string          `json:"district"`
	Images               string          `json:"images"`
	Description          string          `json:"description"`
	Rating               string          `json:"rating"`
	Facilities           json.RawMessage `json:"facilities"`
	GoogleMap            string          `json:"google_map"`
	PropertySurroundings json.RawMessage `json:"property_surrounds"`
	Rules                string          `json:"rules"`
}
