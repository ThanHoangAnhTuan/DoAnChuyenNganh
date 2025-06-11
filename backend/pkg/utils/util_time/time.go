package utiltime

import (
	"context"
	"fmt"
	"time"

	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils/timezone"
)

func GetTimeNow() uint64 {
	now := time.Now()
	epochMillis := now.UnixMilli()
	return uint64(epochMillis)
}

func ConvertISOToUnixTimestamp(dateStr string) (uint64, error) {
	t, err := time.Parse("02-01-2006", dateStr)
	if err != nil {
		return 0, fmt.Errorf("không thể phân tích chuỗi ngày: %v", err)
	}

	utcTime := time.Date(
		t.Year(),
		t.Month(),
		t.Day(),
		0,
		0,
		0,
		0,
		time.UTC,
	)

	unixMilliseconds := uint64(utcTime.UnixNano() / int64(time.Millisecond))

	return unixMilliseconds, nil
}

func FormatVNPayTime(vnpayTime string) (string, error) {
	// Định dạng của chuỗi thời gian từ VNPAY
	const vnpayLayout = "20060102150405" // yyyyMMddHHmmss

	// Load múi giờ GMT+7 (Asia/Ho_Chi_Minh)
	loc, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	if err != nil {
		return "", fmt.Errorf("failed to load location: %s", err)
	}

	// Parse chuỗi thời gian với múi giờ GMT+7
	t, err := time.ParseInLocation(vnpayLayout, vnpayTime, loc)
	if err != nil {
		return "", fmt.Errorf("failed to parse time: %s", err)
	}

	// Định dạng đầu ra dễ đọc, ví dụ: "13:05, 24/09/2015"
	formattedTime := t.Format("15:04, 02/01/2006")

	return formattedTime, nil
}

func ConvertUnixTimestampToISO(ctx context.Context, timestamp int64) (string, error) {
	utcTime := time.Unix(0, timestamp*int64(time.Millisecond))

	timezoneStr, ok := timezone.GetTimezone(ctx)
	if !ok {
		return utcTime.UTC().Format("02-01-2006"), nil
	}

	loc, err := timezone.GetLocation(ctx)
	if err != nil {
		return "", fmt.Errorf("không thể tải múi giờ %s: %v", timezoneStr, err)
	}

	return utcTime.In(loc).Format("02-01-2006"), nil
}
