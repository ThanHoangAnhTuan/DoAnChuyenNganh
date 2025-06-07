package pagination

import (
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
)

func NewPagination(page, limit int, total int64) *vo.Pagination {
	totalPages := int((total + int64(limit) - 1) / int64(limit))

	return &vo.Pagination{
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: totalPages,
	}
}
