package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	m "bookwormia/book/internal/middleware"
)

func New() *echo.Echo {
	router := echo.New()
	router.Logger.SetLevel(log.DEBUG)
	router.Use(middleware.Logger())
	router.Use(m.JWTMiddleware)
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{echo.POST, echo.GET},
	}))
	return router
}
