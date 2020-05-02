package handlers

import (
	"fmt"
	"net/http"

	"github.com/adamdevigili/skillbased.io/pkg/core"
	"github.com/adamdevigili/skillbased.io/pkg/db"
	"github.com/adamdevigili/skillbased.io/pkg/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func (h *Handler) GenerateTeams(c echo.Context) error {
	t := &models.GenerateTeamRequest{}

	if err := c.Bind(t); err != nil {
		e := models.Errors{Errors: []models.Error{
			{
				Status:    http.StatusBadRequest,
				Title:     "failed to bind JSON",
				Detail:    "please check your JSON structure",
				RequestID: c.Request().Header.Get(echo.HeaderXRequestID),
			},
		}}
		log.Error("ded")
		return c.JSON(http.StatusBadRequest, e)
	}

	sport, ok := db.SportsMem[t.SportID]
	if !ok {
		e := models.Errors{Errors: []models.Error{
			{
				Status:    http.StatusBadRequest,
				Title:     fmt.Sprintf("team with id %s not found", t.SportID),
				Detail:    "please check the provided ID is correct and try again",
				RequestID: c.Request().Header.Get(echo.HeaderXRequestID),
			},
		}}
		return c.JSON(http.StatusBadRequest, e)
	}

	return c.JSON(http.StatusCreated, core.GenerateTeams(*t, *sport))
}
