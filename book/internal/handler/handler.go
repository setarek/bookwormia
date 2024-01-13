package handler

import (
	"bookwormia/book/internal/cache"
	"bookwormia/book/internal/repository"
)

type BookHandler struct {
	Repository repository.IBook
	Cache      cache.ICache
}

func NewHandler(bookRepository repository.IBook, cache cache.ICache) *BookHandler {
	return &BookHandler{
		Repository: bookRepository,
		Cache:      cache,
	}
}
