package initializer

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/global"
	"go.uber.org/zap"
)

func InitRedis() {
	r := global.Config.Redis

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", r.Host, r.Port),
		Password: r.Password, // no password set
		DB:       r.Database, // use default DB
		PoolSize: r.PoolSize,
	})

	_, err := rdb.Ping(global.Ctx).Result()
	if err != nil {
		global.Logger.Error("InitRedis error: ", zap.Error(err))
	}
	global.Logger.Info("InitRedis success")
	global.Redis = rdb
}
