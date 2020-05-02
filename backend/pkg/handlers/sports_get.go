package handlers

import (
	"net/http"

	"github.com/adamdevigili/skillbased.io/pkg/constants"

	"github.com/adamdevigili/skillbased.io/pkg/db"
	"github.com/adamdevigili/skillbased.io/pkg/models"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetSport(c echo.Context) error {
	id := c.Param("id")
	sport, ok := db.SportsMem[id]
	if !ok {
		return c.JSON(http.StatusNotFound, models.GenNotFoundError("sport", id, c.Get(constants.RequestIDKey).(string)))
	}

	return c.JSON(http.StatusOK, sport)
}
