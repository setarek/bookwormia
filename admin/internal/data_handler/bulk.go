package data_handler

import (
	"bookwormia/pkg/utils"
	"encoding/json"
)

const (
	NEWBOOKS    = "BOOKS:NEW"
	EDITBOOKS   = "BOOKS:EDIT"
	DELETEBOOKS = "BOOKS:DELETE"
)

func (d dataHandler) CreateBulk(header []string, records [][]string) error {

	var books []Book
	for _, record := range records {
		book := Book{}
		for i, field := range header {
			switch field {
			case "Name":
				book.Name = record[i]
			case "Description":
				book.Description = record[i]

			}

		}
		books = append(books, book)
	}

	b, err := json.Marshal(books)
	if err != nil {
		return err
	}

	if _, err := d.repository.LPush(NEWBOOKS, b); err != nil {
		return err
	}
	return nil
}

func (d dataHandler) UpdateBulk(header []string, records [][]string) error {

	var books []Book
	book := Book{}
	for _, record := range records {
		for i, field := range header {
			switch field {
			case "Id":
				book.Id = utils.ParseInt64(record[i])
			case "Name":
				book.Name = record[i]
			case "Description":
				book.Description = record[i]
			}
		}
		books = append(books, book)
	}

	b, err := json.Marshal(books)
	if err != nil {
		return err
	}

	if _, err := d.repository.LPush(EDITBOOKS, b); err != nil {
		return err
	}
	return nil
}

func (d dataHandler) DeleteBulk(ids []int64) error {

	b, err := json.Marshal(ids)
	if err != nil {
		return err
	}

	if _, err := d.repository.LPush(DELETEBOOKS, b); err != nil {
		return err
	}
	return nil
}
