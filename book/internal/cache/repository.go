package cache

import "github.com/go-redis/redis"

func (c Cache) RPop(key string) (string, error) {
	result, err := c.Client.RPop(key).Result()
	if err != nil && err != redis.Nil {
		return "", err
	} else if err == redis.Nil {
		return "", nil
	}
	return result, nil
}
