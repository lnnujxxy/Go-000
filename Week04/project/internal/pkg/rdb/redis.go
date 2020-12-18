package rdb

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"gogeekbang/internal/pkg/config"
)

func NewRedis(conf config.Redis) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Password: conf.Password,
		DB:       conf.DB,
		PoolSize: conf.PoolSize,
	})
}
