package vo

type RegisterInput struct {
	VerifyKey     string `json:"verify_key" validate:"required"`
	VerifyType    int    `json:"verify_type" validate:"required"`
	VerifyPurpose string `json:"verify_purpose" validate:"required"`
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
	UserId string `json:"user_id" validate:"required"`
}

type LoginInput struct {
	UserAccount  string `json:"account" validate:"required"`
	UserPassword string `json:"password" validate:"required"`
}

type LoginOutput struct {
	Token string `json:"token" validate:"required"`
}
