package handlers

import (
	"fmt"
	"net/http"

	"github.com/Pallinder/go-randomdata"
	"github.com/adamdevigili/balancer.team/pkg/core"
	"github.com/adamdevigili/balancer.team/pkg/db"
	"github.com/adamdevigili/balancer.team/pkg/models"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/rs/xid"
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

	teams := core.GenerateTeams(t.Players, *sport, t.NumberOfTeams)

	// Initialize the teams with new IDs, names, set their sports, and store them in memory
	for i := range teams {
		teams[i].Sport = *sport
		teams[i].ID = xid.New().String()
		teams[i].Name = randomdata.SillyName()
		db.TeamsMem[teams[i].ID] = &teams[i]
	}

	return c.JSON(http.StatusCreated, teams)
}
