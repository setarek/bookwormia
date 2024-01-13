package cache

type ICache interface {
	AddNewScore(bookId int64) error
	AddNewComment(bookId int64) error

	GetCommentCount(bookId int64) (int64, error)
	GetScoreCount(bookId int64) (int64, error)

	RPop(key string) (string, error)
}
