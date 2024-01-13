package repository

import (
	"bookwormia/book/internal/model"
	"context"
)

type IBook interface {
	CreateBook(ctx context.Context, name string, description string) error
	UpdateBook(ctx context.Context, id int64, name string, description string) error
	DeleteBook(ctx context.Context, id int64) error

	GetBooks(ctx context.Context, userId int, pageNumber int, pageSize int) ([]BookInfo, error)
	GetBookDetails(ctx context.Context, id int64) (model.Book, error)
	GetBookReview(ctx context.Context, id int64) ([]BookReview, error)
	GetBookAvgScore(ctx context.Context, bookId int64) (float64, error)

	UpsertScore(ctx context.Context, userId int64, bookId int64, score int64) error
	UpsertComment(ctx context.Context, userId int64, bookId int64, comment string) error
	GetBookCommentByUserId(ctx context.Context, userId int64, bookId int64) (string, error)

	BookMark(ctx context.Context, bookId int64, userId int64, marked bool) error
	GetDisticScore(ctx context.Context, bookId int64) (DistictScore, error)
	GetBookScoreByUserId(ctx context.Context, userId int64, bookId int64) (int64, error)
}
