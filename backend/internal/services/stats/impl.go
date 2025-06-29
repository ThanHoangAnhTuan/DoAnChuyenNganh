package stats

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/database"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/response"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils"
)

type serviceImpl struct {
	sqlc *database.Queries
}

func New(sqlc *database.Queries) Service {
	return &serviceImpl{
		sqlc: sqlc,
	}
}

func (s *serviceImpl) GetDailyEarnings(ctx *gin.Context) (codeStatus int, out []*vo.GetDailyEarningsOutput, err error) {
	clientTZ := ctx.GetHeader("X-Timezone")
	now := s.getCurrentTimeForClient(clientTZ)

	return s.GetDailyEarningsByMonth(ctx, &vo.GetDailyEarningsByMonthInput{
		Year:  now.Year(),
		Month: int(now.Month()),
	})
}

func (s *serviceImpl) GetDailyEarningsByMonth(ctx *gin.Context, in *vo.GetDailyEarningsByMonthInput) (codeStatus int, out []*vo.GetDailyEarningsOutput, err error) {
	// TODO: get user from context
	userID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		return response.ErrCodeUnauthorized, nil, fmt.Errorf("userID not found in context")
	}

	// TODO: check user is manager
	manager, err := s.sqlc.CheckUserManagerExistsByID(ctx, userID)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("error for get manager: %s", err)
	}

	if !manager {
		return response.ErrCodeForbidden, nil, fmt.Errorf("manager not found")
	}

	clientTZ := ctx.GetHeader("X-Timezone")
	startEpoch, endEpoch := s.getMonthTimeRange(in.Year, in.Month, clientTZ)

	dailyEarnings, err := s.sqlc.DailyEarnings(ctx, database.DailyEarningsParams{
		ManagerID: userID,
		StartTime: startEpoch,
		EndTime:   endEpoch,
	})

	if err != nil {
		return response.ErrCodeInternalServerError, nil, err
	}

	for _, dailyEarning := range dailyEarnings {

		out = append(out, &vo.GetDailyEarningsOutput{
			Day:          dailyEarning.Day.Format("02-01-2006"),
			TotalOrders:  dailyEarning.TotalOrders,
			TotalRevenue: dailyEarning.TotalRevenue,
		})
	}

	return response.ErrCodeStatsSuccess, out, nil
}

func (s *serviceImpl) GetMonthlyEarningsByYear(ctx *gin.Context, in *vo.GetMonthlyEarningsByYearInput) (codeStatus int, out []*vo.GetMonthlyEarningsOutput, err error) {
	// TODO: get user from context
	userID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		return response.ErrCodeUnauthorized, nil, fmt.Errorf("userID not found in context")
	}

	// TODO: check user is manager
	manager, err := s.sqlc.CheckUserManagerExistsByID(ctx, userID)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("error for get manager: %s", err)
	}

	if !manager {
		return response.ErrCodeForbidden, nil, fmt.Errorf("manager not found")
	}

	// TODO: get monthly earnings of manager
	clientTZ := ctx.GetHeader("X-Timezone")
	startEpoch, endEpoch := s.getYearTimeRange(in.Year, clientTZ)

	monthlyEarnings, err := s.sqlc.MonthlyEarnings(ctx, database.MonthlyEarningsParams{
		ManagerID: userID,
		StartTime: startEpoch,
		EndTime:   endEpoch,
	})

	if err != nil {
		return response.ErrCodeInternalServerError, nil, err
	}

	for _, monthlyEarning := range monthlyEarnings {
		out = append(out, &vo.GetMonthlyEarningsOutput{
			Month:        monthlyEarning.Month,
			TotalOrders:  monthlyEarning.TotalOrders,
			TotalRevenue: monthlyEarning.TotalRevenue,
		})
	}

	return response.ErrCodeStatsSuccess, out, nil
}

func (s *serviceImpl) GetMonthlyEarnings(ctx *gin.Context) (codeStatus int, out []*vo.GetMonthlyEarningsOutput, err error) {
	clientTZ := ctx.GetHeader("X-Timezone")
	now := s.getCurrentTimeForClient(clientTZ)

	return s.GetMonthlyEarningsByYear(ctx, &vo.GetMonthlyEarningsByYearInput{
		Year: now.Year(),
	})
}

func (s *serviceImpl) ExportDailyEarningsCSV(ctx *gin.Context, in *vo.GetDailyEarningsByMonthInput) (codeStatus int, csvData []byte, err error) {
	statusCode, dailyEarnings, err := s.GetDailyEarningsByMonth(ctx, in)
	if err != nil {
		return statusCode, nil, err
	}

	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	header := []string{"Date", "Total Orders", "Total Revenue"}
	if err := writer.Write(header); err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("failed to write CSV header: %w", err)
	}

	var totalOrders int64
	var totalRevenue float64

	for _, earning := range dailyEarnings {
		record := []string{
			earning.Day,
			strconv.FormatInt(earning.TotalOrders, 10),
			earning.TotalRevenue,
		}
		if err := writer.Write(record); err != nil {
			return response.ErrCodeInternalServerError, nil, fmt.Errorf("failed to write CSV record: %w", err)
		}

		rev, err := strconv.ParseFloat(earning.TotalRevenue, 64)
		if err != nil {
			return response.ErrCodeInternalServerError, nil, err
		}

		totalOrders += earning.TotalOrders
		totalRevenue += rev
	}

	// Write summary
	summaryRecord := []string{
		"TOTAL",
		strconv.FormatInt(totalOrders, 10),
		fmt.Sprintf("%.2f", totalRevenue),
	}
	if err := writer.Write(summaryRecord); err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("failed to write CSV summary: %w", err)
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("CSV writer error: %w", err)
	}

	return response.ErrCodeExportSuccess, buf.Bytes(), nil
}

func (s *serviceImpl) ExportMonthlyEarningsCSV(ctx *gin.Context, in *vo.GetMonthlyEarningsByYearInput) (codeStatus int, csvData []byte, err error) {
	// Get monthly earnings data
	statusCode, monthlyEarnings, err := s.GetMonthlyEarningsByYear(ctx, in)
	if err != nil {
		return statusCode, nil, err
	}

	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	// Write header
	header := []string{"Month", "Total Orders", "Total Revenue"}
	if err := writer.Write(header); err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("failed to write CSV header: %w", err)
	}

	// Write data
	var totalOrders int64
	var totalRevenue float64

	monthNames := []string{
		"", "January", "February", "March", "April", "May", "June",
		"July", "August", "September", "October", "November", "December",
	}

	for _, earning := range monthlyEarnings {
		monthName := monthNames[earning.Month]
		record := []string{
			monthName,
			strconv.FormatInt(earning.TotalOrders, 10),
			earning.TotalRevenue,
		}
		if err := writer.Write(record); err != nil {
			return response.ErrCodeInternalServerError, nil, fmt.Errorf("failed to write CSV record: %w", err)
		}

		rev, err := strconv.ParseFloat(earning.TotalRevenue, 64)
		if err != nil {
			return response.ErrCodeInternalServerError, nil, err
		}

		totalOrders += earning.TotalOrders
		totalRevenue += rev
	}

	// Write summary
	summaryRecord := []string{
		"TOTAL",
		strconv.FormatInt(totalOrders, 10),
		fmt.Sprintf("%.2f", totalRevenue),
	}
	if err := writer.Write(summaryRecord); err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("failed to write CSV summary: %w", err)
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("CSV writer error: %w", err)
	}

	return response.ErrCodeExportSuccess, buf.Bytes(), nil
}

func (s *serviceImpl) getCurrentTimeForClient(clientTimezone string) time.Time {
	if clientTimezone == "" {
		return time.Now().UTC()
	}

	loc, err := time.LoadLocation(clientTimezone)
	if err != nil {
		return time.Now().UTC()
	}

	return time.Now().In(loc)
}

func (s *serviceImpl) getYearTimeRange(year int, clientTimezone string) (startEpoch, endEpoch uint64) {
	loc := s.getLocation(clientTimezone)

	startOfYear := time.Date(year, time.January, 1, 0, 0, 0, 0, loc)
	endOfYear := time.Date(year, time.December, 31, 23, 59, 59, 999999999, loc)

	return uint64(startOfYear.UTC().UnixMilli()), uint64(endOfYear.UTC().UnixMilli())
}

func (s *serviceImpl) getMonthTimeRange(year, month int, clientTimezone string) (startEpoch, endEpoch uint64) {
	loc := s.getLocation(clientTimezone)

	startOfMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, loc)
	endOfMonth := startOfMonth.AddDate(0, 1, 0).Add(-1 * time.Nanosecond) // Last nanosecond of month

	return uint64(startOfMonth.UTC().UnixMilli()), uint64(endOfMonth.UTC().UnixMilli())
}

func (s *serviceImpl) getLocation(clientTimezone string) *time.Location {
	if clientTimezone == "" {
		return time.UTC
	}

	loc, err := time.LoadLocation(clientTimezone)
	if err != nil {
		return time.UTC
	}

	return loc
}
