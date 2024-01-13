package handler

import "github.com/labstack/echo/v4"

func (h *UserHandler) Register(v1 *echo.Group) {
	c := v1.Group("/login")
	c.POST("", h.Login)
}
