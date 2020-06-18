package handlers

import (
	"net/http"

	"github.com/adamdevigili/skillbased.io/pkg/models"
	"github.com/labstack/echo/v4"
)

// Health returns a simple JSON response, indicating that the server is alive
func (h *Handler) Health(c echo.Context) error {
	return c.JSON(http.StatusOK, models.Health{
		Status: "up and running",
	})
}
