package handlers

import (
	"net/http"

	"github.com/adamdevigili/skillbased.io/pkg/models"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Health(c echo.Context) error {
	return c.JSON(http.StatusOK, models.Health{
		Status: "up and running",
	})
}
