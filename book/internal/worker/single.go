package worker

import (
	"bookwormia/book/internal/model"
	"bookwormia/pkg/logger"
	"context"
	"encoding/json"
	"time"
)

const (
	NEWBOOK    = "BOOK:NEW"
	EDITBOOK   = "BOOK:EDIT"
	DELETEBOOK = "BOOK:DELETE"
)

func (w *Worker) startCreateOneBookWorker() {
	logger.Logger.Info().Msg("create one book worker started")

	for {
		select {
		case <-w.stopChan:
			logger.Logger.Info().Msg("create one book stoped")
			return
		default:
			bookPayload, err := w.RedisRepository.RPop(NEWBOOK)
			if err == nil && bookPayload == "" {
				time.Sleep(1 * time.Second)
				continue
			} else if err != nil {
				logger.Logger.Error().Err(err).Msg("error while retrieving book message")
				time.Sleep(1 * time.Second)
				continue
			}

			var book model.Book
			err = json.Unmarshal([]byte(bookPayload), &book)
			if err != nil {
				logger.Logger.Error().Err(err).Msg("error while unmarshalling book message")
			}

			if err := w.BookRepository.CreateBook(context.Background(), book.Name, book.Description); err != nil {
				logger.Logger.Error().Err(err).Msg("error while creating book message")
			}
		}
	}
}

func (w *Worker) startUpdateOneBookWorker() {
	logger.Logger.Info().Msg("update one book worker started")

	for {
		select {
		case <-w.stopChan:
			logger.Logger.Info().Msg("updating one book stoped")
			return
		default:
			bookPayload, err := w.RedisRepository.RPop(EDITBOOK)
			if err == nil && bookPayload == "" {
				time.Sleep(1 * time.Second)
				continue
			} else if err != nil {
				logger.Logger.Error().Err(err).Msg("error while retrieving book message")
				time.Sleep(1 * time.Second)
				continue
			}

			var book model.Book
			err = json.Unmarshal([]byte(bookPayload), &book)
			if err != nil {
				logger.Logger.Error().Err(err).Msg("error while unmarshalling book message")
			}
			if err := w.BookRepository.UpdateBook(context.Background(), book.Id, book.Name, book.Description); err != nil {
				logger.Logger.Error().Err(err).Msg("error while updating book message")
			}
		}
	}
}

func (w *Worker) startDeleteOneBookWorker() {
	logger.Logger.Info().Msg("delete one book worker started")

	for {
		select {
		case <-w.stopChan:
			logger.Logger.Info().Msg("deleting one book stoped")
			return
		default:
			bookIdPayload, err := w.RedisRepository.RPop(DELETEBOOK)
			if err == nil && bookIdPayload == "" {
				time.Sleep(1 * time.Second)
				continue
			} else if err != nil {
				logger.Logger.Error().Err(err).Msg("error while retrieving book message")
				time.Sleep(1 * time.Second)
				continue
			}

			var bookId int64
			err = json.Unmarshal([]byte(bookIdPayload), &bookId)
			if err != nil {
				logger.Logger.Error().Err(err).Msg("error while unmarshalling book message")
			}

			if err := w.BookRepository.DeleteBook(context.Background(), bookId); err != nil {
				logger.Logger.Error().Err(err).Msg("error while deleting book message")
			}
		}
	}
}
