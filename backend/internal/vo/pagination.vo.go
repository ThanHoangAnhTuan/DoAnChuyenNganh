package vo

type PaginationInput interface {
	GetPage() int32
	GetLimit() int32
}

type BasePaginationInput struct {
	Page  *int32 `form:"page,omitempty"`
	Limit *int32 `form:"limit,omitempty"`
}

type BasePaginationOutput struct {
	Page       int32 `json:"page"`
	Limit      int32 `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int64 `json:"total_pages"`
}

func (b *BasePaginationInput) GetPage() int32 {
	if b.Page == nil || *b.Page <= 0 {
		return 1
	}
	return *b.Page
}

func (b *BasePaginationInput) GetLimit() int32 {
	if b.Limit == nil || *b.Limit <= 0 {
		return 10
	}
	return *b.Limit
}
