package data_handler

import (
	redisRepository "bookwormia/admin/internal/redis_repository"
	"bookwormia/pkg/logger"
	"encoding/json"
)

const (
	NEWBOOK    = "BOOK:NEW"
	EDITBOOK   = "BOOK:EDIT"
	DELETEBOOK = "BOOK:DELETE"
)

type dataHandler struct {
	repository redisRepository.RedisRepository
}

func NewDataHandler(r redisRepository.RedisRepository) DataHandler {
	return &dataHandler{
		repository: r,
	}
}
func (d dataHandler) CreateOne(book Book) error {
	b, err := json.Marshal(book)
	if err != nil {
		logger.Logger.Error().Err(err).Msg("error while marshalling book object")
	}

	if _, err := d.repository.LPush(NEWBOOK, b); err != nil {
		return err
	}
	return nil
}

func (d dataHandler) UpdateOne(book Book) error {

	b, err := json.Marshal(book)
	if err != nil {
		logger.Logger.Error().Err(err).Msg("error while marshalling book object")
	}

	if _, err := d.repository.LPush(EDITBOOK, b); err != nil {
		return err
	}
	return nil
}

func (d dataHandler) DeleteOne(id int64) error {
	if _, err := d.repository.LPush(DELETEBOOK, id); err != nil {
		return err
	}
	return nil
}
