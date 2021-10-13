package handlers

import (
	"net/http"

	"github.com/adamdevigili/skillbased/api/pkg/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// Health returns a simple JSON response, indicating that the server is alive
func (h *Handler) Health(c echo.Context) error {
	log.Info()
	return c.JSON(http.StatusOK, models.Health{
		Status: "healthy",
	})
}
