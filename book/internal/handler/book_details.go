package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	rsError "bookwormia/pkg/error"
	"bookwormia/pkg/logger"
	"bookwormia/pkg/utils"
)

func (h BookHandler) GetBookDetails(ctx echo.Context) error {

	if ctx.Param("id") == "" {
		return ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Message: rsError.ErrorNoQueryParam.Error(),
		})
	}

	bookId := utils.ParseInt64(ctx.Param("id"))

	book, err := h.Repository.GetBookDetails(ctx.Request().Context(), bookId)
	if err != nil {
		logger.Logger.Error().Err(err).Msg("error while getting book details")
		return ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Message: rsError.ErrorEmptyQuery.Error(),
		})
	}

	bookReview, err := h.Repository.GetBookReview(ctx.Request().Context(), bookId)
	if err != nil {
		logger.Logger.Error().Err(err).Msg("error while getting book reviews")
		return ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Message: rsError.ServerErr.Error(),
		})
	}

	// todo: implement a worker to calculate score per book every hour and store it in redis
	// then get average score from redis
	avgScore, err := h.Repository.GetBookAvgScore(ctx.Request().Context(), bookId)
	if err != nil {
		logger.Logger.Error().Err(err).Msg("error while getting average score")
		return ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Message: rsError.ServerErr.Error(),
		})
	}

	// todo: implement a worker to calculate distinct score per book every hour and store it in redis
	// then get distinct score from redis
	distinctScore, err := h.Repository.GetDisticScore(ctx.Request().Context(), bookId)
	if err != nil {
		logger.Logger.Error().Err(err).Msg("error while getting distinct score")
		return ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Message: rsError.ServerErr.Error(),
		})
	}

	scoreCount, err := h.Cache.GetScoreCount(bookId)
	if err != nil {
		logger.Logger.Error().Err(err).Msg("error while getting score count")
		return ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Message: rsError.ServerErr.Error(),
		})
	}

	commentCount, err := h.Cache.GetCommentCount(bookId)
	if err != nil {
		logger.Logger.Error().Err(err).Msg("error while getting score count")
		return ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Message: rsError.ServerErr.Error(),
		})
	}

	response := Respone{
		Success: true,
		Data: map[string]interface{}{
			"name":           book.Name,
			"description":    book.Description,
			"review":         NewReviewRespone(bookReview),
			"average_score":  avgScore,
			"distinct_score": distinctScore,
			"score_count":    scoreCount,
			"comment_count":  commentCount,
		},
	}

	return ctx.JSON(http.StatusOK, response)
}
