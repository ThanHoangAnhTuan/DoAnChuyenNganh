package vo

import "mime/multipart"

type RegisterInput struct {
	VerifyKey     string `json:"verify_key" validate:"required"`
	VerifyType    uint8  `json:"verify_type" validate:"required"`
	VerifyPurpose string `json:"verify_purpose"`
}

type VerifyOTPInput struct {
	VerifyKey  string `json:"verify_key" validate:"required"`
	VerifyCode string `json:"verify_code" validate:"required"`
}

type VerifyOTPOutput struct {
	Token string `json:"token" validate:"required"`
}

type UpdatePasswordRegisterInput struct {
	Token    string `json:"token" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UpdatePasswordRegisterOutput struct {
	UserID string `json:"user_id" validate:"required"`
}

type LoginInput struct {
	UserAccount  string `json:"account" validate:"required"`
	UserPassword string `json:"password" validate:"required"`
}

type LoginOutput struct {
	Token string `json:"token" validate:"required"`
}

type GetUserInfoOutput struct {
	ID       string `json:"id"`
	Account  string `json:"account"`
	Username string `json:"username"`
	Image    string `json:"image"`
	Phone    string `json:"phone"`
	Gender   string `json:"gender"`
	Birthday string `json:"birthday"`
}

type UpdateUserInfoOutput struct {
	ID       string `json:"id"`
	Account  string `json:"account"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Gender   string `json:"gender"`
	Birthday string `json:"birthday"`
	// Email    string `json:"email"`
}

type UpdateUserInfoInput struct {
	// Account  string `json:"account"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Gender   uint8  `json:"gender"`
	Birthday string `json:"birthday"`
	// Email    string `json:"email"`
}

type UploadUserAvatarInput struct {
	Avatar *multipart.FileHeader `form:"avatar"`
}
