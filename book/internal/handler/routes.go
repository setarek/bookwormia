package handler

import "github.com/labstack/echo/v4"

func (h *BookHandler) Register(v1 *echo.Group) {
	c := v1.Group("/books")
	c.GET("/list", h.GetList)
	c.GET("/:id", h.GetBookDetails)

	c.POST("/review", h.AddReview)
	c.POST("/bookmark", h.AddBookmark)
}
