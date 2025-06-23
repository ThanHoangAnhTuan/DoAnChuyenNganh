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

type GetAccommodationsOfManagerInput struct {
	ManagerID string `uri:"id" validate:"required"`
	BasePaginationInput
}
type GetAccommodationsOfManagerOutput struct {
	ID          string             `json:"id"`
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

type VerifyAccommodationInput struct {
	AccommodationID string `json:"accommodation_id"`
	Status          uint8  `json:"status"` // 1: verify, 0: unverify
}

type VerifyAccommodationOutput struct {
}
