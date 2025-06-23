package vo

type AdminRegisterInput struct {
	UserAccount  string `json:"account" validate:"required,email"`
	UserPassword string `json:"password" validate:"required"`
}

type AdminLoginInput struct {
	UserAccount  string `json:"account" validate:"required"`
	UserPassword string `json:"password" validate:"required"`
}

type AdminLoginOutput struct {
	Token    string `json:"token" validate:"required"`
	Account  string `json:"account" validate:"required"`
	UserName string `json:"user_name" validate:"required"`
}

type AdminInfor struct {
	Account  string `json:"account" validate:"required"`
	UserName string `json:"user_name" validate:"required"`
}

type GetManagerInput struct {
	BasePaginationInput
}

type GetManagerOutput struct {
	ID        string `json:"id"`
	Account   string `json:"account"`
	Username  string `json:"username"`
	IsDeleted bool   `json:"is_deleted"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
