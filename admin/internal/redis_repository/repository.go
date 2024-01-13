package redis_repository

import "github.com/go-redis/redis"

type RedisRepository struct {
	rc *redis.Client
}

func NewRedisRepository(rc *redis.Client) *RedisRepository {
	return &RedisRepository{rc: rc}
}

func (r *RedisRepository) LPush(key string, value interface{}) (int64, error) {
	return r.rc.LPush(key, value).Result()
}

func (r *RedisRepository) RPop(key string) (string, error) {
	result, err := r.rc.RPop(key).Result()
	if err != nil && err != redis.Nil {
		return "", err
	} else if err == redis.Nil {
		return "", nil
	}
	return result, nil
}
