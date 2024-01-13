package cache

import (
	"bookwormia/pkg/utils"
	"fmt"

	"github.com/go-redis/redis"
)

const (
	BookCommentCount = "book:%d:comment"
)

func (c Cache) AddNewComment(bookId int64) error {
	if err := c.Client.Incr(fmt.Sprintf(BookCommentCount, bookId)).Err(); err != nil {
		return err
	}
	return nil
}

func (c Cache) GetCommentCount(bookId int64) (int64, error) {
	result, err := c.Client.Get(fmt.Sprintf(BookCommentCount, bookId)).Result()
	if err == redis.Nil {
		return 0, nil
	} else if err != nil {
		return 0, err
	}
	return utils.ParseInt64(result), nil
}
