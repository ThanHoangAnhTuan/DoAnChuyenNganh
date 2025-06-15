package global

import (
	"context"
	"database/sql"

	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/logger"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/setting"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mysql  *sql.DB
	Redis  *redis.Client
	Kafka  *kafka.Writer

	Ctx = context.Background()
)
