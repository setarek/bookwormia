package repository

import (
	"context"
	"fmt"
)

func (b BookRepository) UpsertComment(ctx context.Context, userId int64, bookId int64, comment string) error {

	query := `
		insert into book_comments (user_id, book_id, comment) values ($1, $2, $3) 
		on conflict (user_id, book_id) do update set comment = excluded.comment;
	`

	fmt.Println("=============", query)

	if _, err := b.DB.ExecContext(ctx, query, userId, bookId, comment); err != nil {
		return err
	}
	return nil
}
