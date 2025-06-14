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

func (s *StatsImpl) GetDailyEarnings(ctx *gin.Context) (codeStatus int, out []*vo.GetDailyEarningsOutput, err error) {
	clientTZ := ctx.GetHeader("X-Timezone")
	now := s.getCurrentTimeForClient(clientTZ)

	return s.GetDailyEarningsByMonth(ctx, &vo.GetDailyEarningsByMonthInput{
		Year:  now.Year(),
		Month: int(now.Month()),
	})
}

func (s *StatsImpl) GetDailyEarningsByMonth(ctx *gin.Context, in *vo.GetDailyEarningsByMonthInput) (codeStatus int, out []*vo.GetDailyEarningsOutput, err error) {
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

	clientTZ := ctx.GetHeader("X-Timezone")
	startEpoch, endEpoch := s.getMonthTimeRange(in.Year, in.Month, clientTZ)

	dailyEarnings, err := s.sqlc.DailyEarnings(ctx, database.DailyEarningsParams{
		ManagerID: userID,
		StartTime: startEpoch,
		EndTime:   endEpoch,
	})

	if err != nil {
		return response.ErrCodeGetMonthlyEarningFailed, nil, err
	}

	for _, dailyEarning := range dailyEarnings {
		out = append(out, &vo.GetDailyEarningsOutput{
			Day:          dailyEarning.Day.Format("02-01-2006"),
			TotalOrders:  dailyEarning.TotalOrders,
			TotalRevenue: dailyEarning.TotalRevenue,
		})
	}

	return response.ErrCodeGetMonthlyEarningSuccess, out, nil
}

func (s *StatsImpl) GetMonthlyEarningsByYear(ctx *gin.Context, in *vo.GetMonthlyEarningsByYearInput) (codeStatus int, out []*vo.GetMonthlyEarningsOutput, err error) {
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
	clientTZ := ctx.GetHeader("X-Timezone")
	startEpoch, endEpoch := s.getYearTimeRange(in.Year, clientTZ)

	monthlyEarnings, err := s.sqlc.MonthlyEarnings(ctx, database.MonthlyEarningsParams{
		ManagerID: userID,
		StartTime: startEpoch,
		EndTime:   endEpoch,
	})

	fmt.Printf("monthlyEarnings: %v", monthlyEarnings)

	if err != nil {
		return response.ErrCodeGetMonthlyEarningFailed, nil, err
	}

	for _, monthlyEarning := range monthlyEarnings {
		out = append(out, &vo.GetMonthlyEarningsOutput{
			Month:        monthlyEarning.Month,
			TotalOrders:  monthlyEarning.TotalOrders,
			TotalRevenue: monthlyEarning.TotalRevenue,
		})
	}

	return response.ErrCodeGetMonthlyEarningSuccess, out, nil
}

func (s *StatsImpl) GetMonthlyEarnings(ctx *gin.Context) (codeStatus int, out []*vo.GetMonthlyEarningsOutput, err error) {
	clientTZ := ctx.GetHeader("X-Timezone")
	now := s.getCurrentTimeForClient(clientTZ)

	return s.GetMonthlyEarningsByYear(ctx, &vo.GetMonthlyEarningsByYearInput{
		Year: now.Year(),
	})
}

func (s *StatsImpl) getCurrentTimeForClient(clientTimezone string) time.Time {
	if clientTimezone == "" {
		return time.Now().UTC()
	}

	loc, err := time.LoadLocation(clientTimezone)
	if err != nil {
		return time.Now().UTC()
	}

	return time.Now().In(loc)
}

func (s *StatsImpl) getYearTimeRange(year int, clientTimezone string) (startEpoch, endEpoch uint64) {
	loc := s.getLocation(clientTimezone)

	startOfYear := time.Date(year, time.January, 1, 0, 0, 0, 0, loc)
	endOfYear := time.Date(year, time.December, 31, 23, 59, 59, 999999999, loc)

	return uint64(startOfYear.UTC().UnixMilli()), uint64(endOfYear.UTC().UnixMilli())
}

func (s *StatsImpl) getMonthTimeRange(year, month int, clientTimezone string) (startEpoch, endEpoch uint64) {
	loc := s.getLocation(clientTimezone)

	startOfMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, loc)
	endOfMonth := startOfMonth.AddDate(0, 1, 0).Add(-1 * time.Nanosecond) // Last nanosecond of month

	return uint64(startOfMonth.UTC().UnixMilli()), uint64(endOfMonth.UTC().UnixMilli())
}

func (s *StatsImpl) getLocation(clientTimezone string) *time.Location {
	if clientTimezone == "" {
		return time.UTC
	}

	loc, err := time.LoadLocation(clientTimezone)
	if err != nil {
		return time.UTC
	}

	return loc
}

func NewStatsImpl(sqlc *database.Queries) *StatsImpl {
	return &StatsImpl{
		sqlc: sqlc,
	}
}
