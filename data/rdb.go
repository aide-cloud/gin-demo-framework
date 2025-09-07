package data

import (
	"context"
	"gin-demo-framework/config"
	"log/slog"

	"github.com/redis/go-redis/v9"
)

var _rdb *redis.Client

func InitRDB(c *config.RedisConfig) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: c.Password,
		DB:       c.DB,
	})
	res := rdb.Ping(context.Background())
	if err := res.Err(); err != nil {
		slog.Error("Redis connect failed", "error", err)
		return
	}
	slog.Info("Redis connected", "res", res)
	_rdb = rdb

}

func GetRDB() *redis.Client {
	return _rdb
}
