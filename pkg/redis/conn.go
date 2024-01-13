package redis

import (
	"fmt"

	"github.com/go-redis/redis"

	"bookwormia/pkg/config"
)

func GetRedisClient(config *config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.GetString("redis_host"), config.GetString("redis_port")),
		Password: config.GetString("redis_password"),
		DB:       config.GetInt("redis_db"),
	})
}
