package services

import (
	"context"

	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
)

type (
	IStats interface {
		GetMonthlyEarnings(ctx context.Context) (codeStatus int, out []*vo.GetMonthlyEarningsOuput, err error)
	}
)

var localStats IStats

func Stats() IStats {
	if localStats == nil {
		panic("Implement localStats not found for interface IStats")
	}
	return localStats
}

func InitStats(i IStats) {
	localStats = i
}
