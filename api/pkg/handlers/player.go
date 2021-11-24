package handlers

import (
	"net/http"

	"github.com/adamdevigili/skillbased/api/pkg/constants"
	"github.com/adamdevigili/skillbased/api/pkg/db"
	"github.com/adamdevigili/skillbased/api/pkg/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func (h *Handler) GetPlayer(c echo.Context) error {
	id := c.Param(constants.URIKeyID)
	player, ok := db.PlayersMem[id]
	if !ok {
		return c.JSON(http.StatusNotFound, models.GenNotFoundError(c, "player", id))
	}

	return c.JSON(http.StatusOK, player)
}

// ListSports list all existing sports
func (h *Handler) ListPlayers(c echo.Context) error {
	playerList, err := db.ListPlayers(h.DB)
	if err != nil {
		e := models.Errors{Errors: []models.Error{
			{
				Status:    http.StatusInternalServerError,
				Message:   "internal server error",
				Detail:    "error when listing players",
				RequestID: c.Response().Header().Get(echo.HeaderXRequestID),
			},
		}}
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, e)
	}

	return c.JSON(http.StatusOK, playerList)
}
