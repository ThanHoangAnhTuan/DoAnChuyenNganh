package vo

type Pagination struct {
	Page       int   `json:"page,omitempty"`
	Limit      int   `json:"limit,omitempty"`
	Total      int64 `json:"total,omitempty"`
	TotalPages int   `json:"total_pages,omitempty"`
}

func (g *GetReviewsInput) GetPage() int {
	if g.Page == nil || *g.Page <= 0 {
		return 1
	}
	return *g.Page
}

// GetLimit trả về limit với default value là 10
func (g *GetReviewsInput) GetLimit() int {
	if g.Limit == nil || *g.Limit <= 0 {
		return 10
	}
	return *g.Limit
}
