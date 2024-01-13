package repository

import "database/sql"

type BookRepository struct {
	DB *sql.DB
}

func NewBookRepository(db *sql.DB) IBook {
	return &BookRepository{
		DB: db,
	}
}

func (b *BookRepository) AddNewBook() error {
	return nil
}
