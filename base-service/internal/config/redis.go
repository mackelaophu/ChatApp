package config

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

// import(
//
//	"github.com/redis/go-redis/v9"
//
// )
type RedisConfig struct {
	Rdb *redis.Client
	Ctx *context.Context
}

var redisConfig *RedisConfig

func ConnectRedis() *RedisConfig {
	if redisConfig == nil {
		redisConfig = newRedisConfig()
	}
	return redisConfig
}
func newRedisConfig() *RedisConfig {
	Rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PWD"),
		DB:       0,
	})
	Ctx := context.Background()
	return &RedisConfig{Rdb: Rdb, Ctx: &Ctx}
}
