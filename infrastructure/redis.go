package infrastructure

import (
	"IkezawaYuki/a-root-backend/config"
	"github.com/redis/go-redis/v9"
)

func GetRedisConnection() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: config.Env.RedisAddr,
		DB:   0,
	})
}
