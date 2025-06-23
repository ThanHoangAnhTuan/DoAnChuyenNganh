package vo

type ManagerRegisterInput struct {
	UserAccount  string `json:"account" validate:"required,email"`
	UserPassword string `json:"password" validate:"required"`
	Username     string `json:"username" validate:"required"`
}

type ManagerLoginInput struct {
	UserAccount  string `json:"account" validate:"required"`
	UserPassword string `json:"password" validate:"required"`
}

type ManagerLoginOutput struct {
	Token    string `json:"token" validate:"required"`
	Account  string `json:"account" validate:"required"`
	UserName string `json:"user_name" validate:"required"`
}

type ManagerInfor struct {
	Account  string `json:"account" validate:"required"`
	UserName string `json:"user_name" validate:"required"`
}
