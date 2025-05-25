package timezone

import (
	"context"
	"time"
)

type contextKey string

const (
	TimezoneKey contextKey = "timezone"
)

func WithTimezone(ctx context.Context, timezone string) context.Context {
	return context.WithValue(ctx, TimezoneKey, timezone)
}

func GetTimezone(ctx context.Context) (string, bool) {
	timezone, ok := ctx.Value(TimezoneKey).(string)
	return timezone, ok
}

func GetLocation(ctx context.Context) (*time.Location, error) {
	timezone, ok := GetTimezone(ctx)
	if !ok {
		return time.UTC, nil
	}
	return time.LoadLocation(timezone)
}
