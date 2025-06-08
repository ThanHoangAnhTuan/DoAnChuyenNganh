package impl

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/database"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/response"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils"
)

type StatsImpl struct {
	sqlc *database.Queries
}

func (s *StatsImpl) GetMonthlyEarnings(ctx *gin.Context) (codeStatus int, out []*vo.GetMonthlyEarningsOuput, err error) {
	// TODO: get user from context
	userID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		return response.ErrCodeUnauthorized, nil, fmt.Errorf("userID not found in context")
	}

	// TODO: check user is manager
	manager, err := s.sqlc.CheckUserManagerExistsByID(ctx, userID)
	if err != nil {
		return response.ErrCodeGetManagerFailed, nil, fmt.Errorf("error for get manager: %s", err)
	}

	if !manager {
		return response.ErrCodeManagerNotFound, nil, fmt.Errorf("manager not found")
	}

	// TODO: get monthly earnings of manager
	currentYear := time.Now().Year()
	monthlyEarnings, err := s.sqlc.MonthlyEarnings(ctx, database.MonthlyEarningsParams{
		UserID:    userID,
		CreatedAt: uint64(currentYear),
	})
	if err != nil {
		return response.ErrCodeGetMonthlyEarningFailed, nil, err
	}

	for _, monthlyEarning := range monthlyEarnings {
		out = append(out, &vo.GetMonthlyEarningsOuput{
			Month:   monthlyEarning.RevenueMonth,
			Revenue: monthlyEarning.TotalRevenue,
		})
	}

	return response.ErrCodeGetMonthlyEarningSuccess, out, nil
}

func NewStatsImpl(sqlc *database.Queries) *StatsImpl {
	return &StatsImpl{
		sqlc: sqlc,
	}
}
