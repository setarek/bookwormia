package repository

import (
	"context"
	"database/sql"
)

func (b *BookRepository) UpsertComment(ctx context.Context, userId int64, bookId int64, comment string) error {

	query := `
		insert into book_comments (user_id, book_id, comment) values ($1, $2, $3) 
		on conflict (user_id, book_id) do update set comment = excluded.comment;
	`
	if _, err := b.DB.ExecContext(ctx, query, userId, bookId, comment); err != nil {
		return err
	}
	return nil
}

func (b *BookRepository) GetBookCommentByUserId(ctx context.Context, userId int64, bookId int64) (string, error) {
	query := `
		select comment from book_comments where user_id = $1 and book_id = $2;
	`
	var comment string
	if err := b.DB.QueryRowContext(ctx, query, userId, bookId).Scan(&comment); err == sql.ErrNoRows {
		return "", nil
	} else if err != nil {
		return "", err
	}
	return comment, nil
}
