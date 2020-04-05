package handlers

import (
	"net/http"

	"github.com/adamdevigili/balancer.team/pkg/constants"

	"github.com/adamdevigili/balancer.team/pkg/db"
	"github.com/adamdevigili/balancer.team/pkg/models"
	"github.com/labstack/echo"
)

func (h *Handler) GetSport(c echo.Context) error {
	id := c.Param("id")
	sport, ok := db.SportsMem[id]
	if !ok {
		return c.JSON(http.StatusNotFound, models.GenNotFoundError("sport", id, c.Get(constants.RequestIDKey).(string)))
	}

	return c.JSON(http.StatusOK, sport)
}
