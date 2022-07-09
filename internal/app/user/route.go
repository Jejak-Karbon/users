package user

import (
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	g.GET("/:id", h.GetByID)
	g.POST("/update-password", h.Update,middleware.Authentication)
}
