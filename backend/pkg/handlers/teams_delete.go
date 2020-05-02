package handlers

import (
	"net/http"

	"github.com/adamdevigili/skillbased.io/pkg/constants"

	"github.com/adamdevigili/skillbased.io/pkg/db"

	"github.com/adamdevigili/skillbased.io/pkg/models"
	"github.com/labstack/echo/v4"
)

func (h *Handler) DeleteTeam(c echo.Context) error {
	id := c.Param("id")
	if _, ok := db.TeamsMem[id]; !ok {
		return c.JSON(http.StatusNotFound, models.GenNotFoundError("team", id, c.Get(constants.RequestIDKey).(string)))
	}

	delete(db.TeamsMem, id)
	return c.NoContent(http.StatusNoContent)
}
