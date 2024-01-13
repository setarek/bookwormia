package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	rsErr "bookwormia/pkg/error"
	"bookwormia/pkg/logger"
	"bookwormia/pkg/utils"
)

func (u *UserHandler) Login(ctx echo.Context) error {
	var request UserRequest
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Message: rsErr.ErrorInvalidBodyRequest.Error(),
		})
	}

	if request.Email == "" || request.Password == "" {
		logger.Logger.Warn().Msg("fill both email and password")
		return ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Message: rsErr.ErrorInvalidBodyRequest.Error(),
		})
	}

	// todo: check email is valid

	userId, err := u.Repository.CreateOrGetUser(ctx.Request().Context(), request.Email, request.Password)
	if err != nil {
		logger.Logger.Error().Err(err).Msg("error while generating jwt token")
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: rsErr.ServerErr.Error(),
		})
	}

	token, err := utils.GenerateJWTToken(userId, u.Config.GetString("secret_key"))
	if err != nil {
		logger.Logger.Error().Err(err).Msg("error while generating jwt token")
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: rsErr.ServerErr.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, Respone{
		Success: true,
		Data: UserReponse{
			Token: token,
		},
	})

}
