package handler

import (
	rsErr "bookwormia/pkg/error"
	"bookwormia/pkg/logger"
	"bookwormia/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h BookHandler) AddBookmark(ctx echo.Context) error {

	userIdStr := ctx.Get("user_id")
	if userIdStr == nil {
		logger.Logger.Error().Msg("unauthorized user")
		return ctx.JSON(http.StatusUnprocessableEntity, ErrorResponse{
			Message: rsErr.ErrorUnauthorizeUser.Error(),
		})
	}
	userId := utils.ParseInt64(userIdStr)

	var request BookmarkRequest
	if err := ctx.Bind(&request); err != nil {
		logger.Logger.Error().Err(err).Msg("error while binding request")
		return ctx.JSON(http.StatusUnprocessableEntity, ErrorResponse{
			Message: rsErr.EmptyBodyRequest.Error(),
		})
	}

	if err := h.Repository.BookMark(ctx.Request().Context(), request.BookId, int64(userId), request.Maked); err != nil {
		logger.Logger.Error().Err(err).Msg("error while bookmark")
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: rsErr.ServerErr.Error(),
		})
	}

	return ctx.JSON(http.StatusAccepted, Respone{
		Success: true,
	})

}
