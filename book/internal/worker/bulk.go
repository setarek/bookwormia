package worker

import (
	"bookwormia/book/internal/model"
	"bookwormia/pkg/logger"
	"context"
	"encoding/json"
	"time"
)

const (
	NEWBOOKS    = "BOOKS:NEW"
	EDITBOOKS   = "BOOKS:EDIT"
	DELETEBOOKS = "BOOKS:DELETE"
)

func (w *Worker) startCreateBulkBookWorker() {
	logger.Logger.Info().Msg("create one book worker started")

	for {
		select {
		case <-w.stopChan:
			logger.Logger.Info().Msg("create one book stoped")
			return
		default:
			booksPayload, err := w.RedisRepository.RPop(NEWBOOKS)
			if err == nil && booksPayload == "" {
				time.Sleep(1 * time.Second)
				continue
			} else if err != nil {
				logger.Logger.Error().Err(err).Msg("error while retrieving book message")
				time.Sleep(1 * time.Second)
				continue
			}

			var books []model.Book
			err = json.Unmarshal([]byte(booksPayload), &books)
			if err != nil {
				logger.Logger.Error().Err(err).Msg("error while unmarshalling book message")
			}

			for _, book := range books {
				if err := w.BookRepository.CreateBook(context.Background(), book.Name, book.Description); err != nil {
					logger.Logger.Error().Err(err).Msg("error while creating book message")
				}

			}
		}
	}
}

func (w *Worker) startUpdateBulkBookWorker() {
	logger.Logger.Info().Msg("updating books worker started")

	for {
		select {
		case <-w.stopChan:
			logger.Logger.Info().Msg("updating books stoped")
			return
		default:

			booksPayload, err := w.RedisRepository.RPop(EDITBOOKS)
			if err == nil && booksPayload == "" {
				time.Sleep(1 * time.Second)
				continue
			} else if err != nil {
				logger.Logger.Error().Err(err).Msg("error while retrieving books message")
				time.Sleep(1 * time.Second)
				continue
			}

			var books []model.Book
			err = json.Unmarshal([]byte(booksPayload), &books)
			if err != nil {
				logger.Logger.Error().Err(err).Msg("error while unmarshalling books message")
			}

			for _, book := range books {
				if err := w.BookRepository.UpdateBook(context.Background(), book.Id, book.Name, book.Description); err != nil {
					logger.Logger.Error().Err(err).Msg("error while updating book message")
				}
			}

		}
	}
}

func (w *Worker) startDeleteBulkBookWorker() {
	logger.Logger.Info().Msg("deleting books worker started")

	for {
		select {
		case <-w.stopChan:
			logger.Logger.Info().Msg("deleting one books stoped")
			return
		default:
			bookIdsPayload, err := w.RedisRepository.RPop(DELETEBOOKS)
			if err == nil && bookIdsPayload == "" {
				time.Sleep(1 * time.Second)
				continue
			} else if err != nil {
				logger.Logger.Error().Err(err).Msg("error while retrieving books message")
				time.Sleep(1 * time.Second)
				continue
			}

			var bookIds []int64
			err = json.Unmarshal([]byte(bookIdsPayload), &bookIds)
			if err != nil {
				logger.Logger.Error().Err(err).Msg("error while unmarshalling book message")
			}

			for _, bookId := range bookIds {
				if err := w.BookRepository.DeleteBook(context.Background(), bookId); err != nil {
					logger.Logger.Error().Err(err).Msg("error while deleting book message")
				}
			}
		}
	}
}
