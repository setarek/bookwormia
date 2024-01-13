package handler

import (
	rsErr "bookwormia/pkg/error"
	"bookwormia/pkg/logger"
	"bookwormia/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h BookHandler) AddReview(ctx echo.Context) error {

	userIdStr := ctx.Get("user_id")
	if userIdStr == nil {
		logger.Logger.Error().Msg("unauthorized user")
		return ctx.JSON(http.StatusUnprocessableEntity, ErrorResponse{
			Message: rsErr.ErrorUnauthorizeUser.Error(),
		})
	}

	userId := utils.ParseInt64(userIdStr)

	var request BookReviewRequest
	if err := ctx.Bind(&request); err != nil {
		logger.Logger.Error().Err(err).Msg("error while binding request")
		return ctx.JSON(http.StatusUnprocessableEntity, ErrorResponse{
			Message: rsErr.EmptyBodyRequest.Error(),
		})
	}

	if request.Comment == "" && request.Score == 0 {
		logger.Logger.Warn().Msg("please fill at least one of comment or score")
		return ctx.JSON(http.StatusUnprocessableEntity, ErrorResponse{
			Message: rsErr.EmptyBodyRequest.Error(),
		})
	}

	if request.Comment != "" {

		comment, err := h.Repository.GetBookCommentByUserId(ctx.Request().Context(), userId, request.BookId)
		if err != nil {
			logger.Logger.Error().Err(err).Msg("error while getting user has any comment in this book")
			return ctx.JSON(http.StatusInternalServerError, ErrorResponse{
				Message: rsErr.ServerErr.Error(),
			})
		}

		if err := h.Repository.UpsertComment(ctx.Request().Context(), int64(userId), request.BookId, request.Comment); err != nil {
			logger.Logger.Error().Err(err).Msg("error while add new comment to database")
			return ctx.JSON(http.StatusInternalServerError, ErrorResponse{
				Message: rsErr.ServerErr.Error(),
			})
		}

		// todo: instead of check postgres that user wrote a comment for this book
		// and add comment count to redis directly, queue it then consume in worker to add to redis

		if comment == "" && err == nil {
			if err := h.Cache.AddNewComment(request.BookId); err != nil {
				logger.Logger.Error().Err(err).Msg("error while add new score to cache")
				return ctx.JSON(http.StatusInternalServerError, ErrorResponse{
					Message: rsErr.ServerErr.Error(),
				})
			}
		}

	}

	if request.Score != 0 {
		if request.Score < 1 || request.Score > 5 {
			logger.Logger.Error().Msg("invalid score. it must be between one and five")
			return ctx.JSON(http.StatusInternalServerError, ErrorResponse{
				Message: rsErr.ErrorScoreEmail.Error(),
			})
		}

		score, err := h.Repository.GetBookScoreByUserId(ctx.Request().Context(), userId, request.BookId)
		if err != nil {
			logger.Logger.Error().Err(err).Msg("error while getting user has any score in this book")
			return ctx.JSON(http.StatusInternalServerError, ErrorResponse{
				Message: rsErr.ServerErr.Error(),
			})
		}

		if err := h.Repository.UpsertScore(ctx.Request().Context(), int64(userId), request.BookId, request.Score); err != nil {
			logger.Logger.Error().Err(err).Msg("error while add new score to database")
			return ctx.JSON(http.StatusInternalServerError, ErrorResponse{
				Message: rsErr.ServerErr.Error(),
			})
		}

		// todo: instead of check postgres at line 75that user scored this book before and
		//add score to redis directly, queue it then consume in worker to add to redis

		if score == 0 {
			if err := h.Cache.AddNewScore(request.BookId); err != nil {
				logger.Logger.Error().Err(err).Msg("error while add new score to cache")
				return ctx.JSON(http.StatusInternalServerError, ErrorResponse{
					Message: rsErr.ServerErr.Error(),
				})
			}
		}
	}

	// todo: at first check user had bookmark this book or not
	if err := h.Repository.BookMark(ctx.Request().Context(), request.BookId, int64(userId), false); err != nil {
		logger.Logger.Error().Err(err).Msg("error while undo bookmared book")
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: rsErr.ServerErr.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, Respone{
		Success: true,
	})

}
