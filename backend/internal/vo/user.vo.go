package vo

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
