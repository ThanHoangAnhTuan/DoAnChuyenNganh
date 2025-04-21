package utiltime

import "time"

func GetTimeNow() uint64 {
	now := time.Now()
	epochMillis := now.UnixMilli()
	return uint64(epochMillis)
}
