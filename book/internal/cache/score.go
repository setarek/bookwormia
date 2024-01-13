package cache

import (
	"bookwormia/pkg/utils"
	"fmt"

	"github.com/go-redis/redis"
)

const (
	BookScore = "book:%d:score"
)

type Cache struct {
	Client *redis.Client
}

func NewCache(client *redis.Client) ICache {
	return Cache{
		Client: client,
	}
}

func (c Cache) AddNewScore(bookId int64) error {
	if err := c.Client.Incr(fmt.Sprintf(BookScore, bookId)).Err(); err != nil {
		return err
	}

	return nil
}

func (c Cache) GetScoreCount(bookId int64) (int64, error) {
	result, err := c.Client.Get(fmt.Sprintf(BookScore, bookId)).Result()
	if err == redis.Nil {
		return 0, nil
	} else if err != nil {
		return 0, err
	}
	return utils.ParseInt64(result), nil
}
