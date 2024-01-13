package middleware

import (
	"bookwormia/pkg/utils"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

const SECRETKEY = "sdfihsiudf89yfisuhf784yt7gsujfgweu7t89hgfuasf74ggh"

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		tokenString := ctx.Request().Header.Get("Authorization")
		if tokenString == "" {
			return next(ctx)
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRETKEY), nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token signature"})
			}
			return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid token"})
		}

		if !token.Valid {
			return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token"})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to retreive token"})
		}

		userId := utils.ParseInt64(claims["user_id"])
		if userId == 0 {
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "user Id not found"})
		}

		ctx.Set("user_id", userId)

		return next(ctx)
	}
}
