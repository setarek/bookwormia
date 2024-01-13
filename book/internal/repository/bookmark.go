package repository

import "context"

func (b BookRepository) BookMark(ctx context.Context, bookId int64, userId int64, marked bool) error {

	query := `
		insert into book_marks (user_id, book_id, marked) values ($1, $2, $3) 
		on conflict (user_id, book_id) do update set marked = excluded.marked;
	`

	if _, err := b.DB.ExecContext(ctx, query, userId, bookId, marked); err != nil {
		return err
	}

	return nil
}
