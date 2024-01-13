package repository

import (
	"bookwormia/book/internal/model"
	"context"
	"database/sql"
	"fmt"
)

type BookRepository struct {
	DB *sql.DB
}

func NewBookRepository(db *sql.DB) IBook {
	return &BookRepository{
		DB: db,
	}
}

type BookInfo struct {
	BookId           int64  `json:"book_id"`
	BookName         string `json:"book_name"`
	BookmarkCount    int64  `json:"bookmark_count"`
	BookMarkedByUser int    `json:"bookmarked_by_user"`
}

func (b *BookRepository) GetBooks(ctx context.Context, userId int, pageNumber int, pageSize int) ([]BookInfo, error) {

	query := `
		select b.id AS book_id, b.name AS book_name, count(distinct case when bm.marked then bm.user_id end) as bookmark_count,
		count(distinct case when bm.user_id = $1 and bm.marked then bm.user_id end) as bookmarked_by_user from books b left join 
		book_marks bm on b.id = bm.book_id group by b.id, b.name, b.description limit $2 offset $3;
	`
	offset := (pageSize - 1) * pageNumber
	rows, err := b.DB.Query(query, userId, pageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var booksInfo []BookInfo
	for rows.Next() {
		var info BookInfo
		err := rows.Scan(&info.BookId, &info.BookName, &info.BookmarkCount, &info.BookMarkedByUser)
		if err != nil {
			return nil, err
		}
		booksInfo = append(booksInfo, info)
	}

	return booksInfo, nil

}

func (b *BookRepository) GetBookDetails(ctx context.Context, id int64) (model.Book, error) {

	query := `
		select name, description from books where id = $1;
	`
	var book model.Book
	if err := b.DB.QueryRow(query, id).Scan(&book.Name, &book.Description); err != nil {
		return model.Book{}, err
	}

	return book, nil
}

type BookReview struct {
	UserId  sql.NullInt64  `json:"user_id"`
	Comment sql.NullString `json:"comment"`
	Score   sql.NullInt64  `json:"score"`
}

func (b *BookRepository) GetBookReview(ctx context.Context, id int64) ([]BookReview, error) {

	query := `
		select u.id as user_id, bc.comment, bs.score from users u left join book_comments bc ON u.id = bc.user_id AND bc.book_id = $1 left join book_scores bs on u.id = bs.user_id and bs.book_id = $1 where bc.comment is not null or bs.score is not null;
	`

	fmt.Println("=====================", query)
	rows, err := b.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var booksReview []BookReview
	for rows.Next() {
		var review BookReview
		err := rows.Scan(&review.UserId, &review.Comment, &review.Score)
		if err != nil {
			return nil, err
		}
		booksReview = append(booksReview, review)
	}
	return booksReview, nil
}

func (b *BookRepository) CreateBook(ctx context.Context, name string, description string) error {
	query := `
		insert into books (name, description) values ($1, $2);
	`

	if _, err := b.DB.ExecContext(ctx, query, name, description); err != nil {
		return err
	}
	return nil
}

func (b *BookRepository) UpdateBook(ctx context.Context, id int64, name string, description string) error {
	query := "update books set"

	if name != "" {
		query += fmt.Sprintf(" name = '%s',", name)
	}

	if description != "" {
		query += fmt.Sprintf(" description = '%s',", description)
	}
	query = query[:len(query)-1]

	query += fmt.Sprintf(" where id = %d;", id)

	if _, err := b.DB.ExecContext(ctx, query); err != nil {
		return err
	}
	return nil
}

func (b *BookRepository) DeleteBook(ctx context.Context, id int64) error {
	query := `
		delete from books where id = $1 
	`

	if _, err := b.DB.ExecContext(ctx, query, id); err != nil {
		return err
	}
	return nil
}
