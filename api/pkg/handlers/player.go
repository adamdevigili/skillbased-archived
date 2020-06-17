package handlers

import (
	"net/http"

	"github.com/adamdevigili/skillbased.io/pkg/constants"
	"github.com/adamdevigili/skillbased.io/pkg/db"
	"github.com/adamdevigili/skillbased.io/pkg/models"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetPlayer(c echo.Context) error {
	id := c.Param(constants.URIKeyID)
	player, ok := db.PlayersMem[id]
	if !ok {
		return c.JSON(http.StatusNotFound, models.GenNotFoundError("player", id, c.Get(constants.RequestIDKey).(string)))
	}

	return c.JSON(http.StatusOK, player)
}
