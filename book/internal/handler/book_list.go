package handler

import (
	"bookwormia/pkg/logger"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	rsErr "bookwormia/pkg/error"
)

const (
	defaultPageNumber = 0
	defaultPageSize   = 3

	baseUrl = "localhost:8080"
)

func (h BookHandler) GetList(ctx echo.Context) error {
	pageNumber := defaultPageNumber
	if ctx.QueryParam("page_number") != "" {
		page, _ := strconv.Atoi(ctx.QueryParam("page_number"))
		pageNumber = page
	}

	pageSize := defaultPageSize
	if ctx.QueryParam("page_size") != "" {
		size, _ := strconv.Atoi(ctx.QueryParam("page_size"))
		pageSize = size
	}

	books, err := h.Repository.GetBooks(ctx.Request().Context(), 1, pageNumber, pageSize)
	if len(books) == 0 {
		logger.Logger.Info().Msg("these isn't any book!")
		return ctx.JSON(http.StatusNotFound, Respone{
			Success: false,
		})
	}
	if err != nil {
		logger.Logger.Error().Err(err).Msg("error while getting book list")
		return ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Message: rsErr.ServerErr.Error(),
		})
	}

	response := NewBookListResponse(baseUrl, books)
	return ctx.JSON(http.StatusOK, Respone{
		Success: true,
		Data:    response,
	})
}
