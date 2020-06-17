package handlers

import (
	"fmt"
	"net/http"

	"github.com/adamdevigili/skillbased.io/pkg/constants"
	"github.com/adamdevigili/skillbased.io/pkg/core"
	"github.com/adamdevigili/skillbased.io/pkg/db"
	"github.com/adamdevigili/skillbased.io/pkg/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/rs/xid"
)

func (h *Handler) CreateTeam(c echo.Context) error {
	t := &models.Team{}

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

	xid := xid.New().String()
	t.ID = xid

	db.TeamsMem[xid] = t

	return c.JSON(http.StatusCreated, t)
}

func (h *Handler) GetTeam(c echo.Context) error {
	id := c.Param(constants.URIKeyID)
	team, ok := db.TeamsMem[id]
	if !ok {
		return c.JSON(http.StatusNotFound, models.GenNotFoundError("team", id, c.Get(constants.RequestIDKey).(string)))
	}

	return c.JSON(http.StatusOK, team)
}

func (h *Handler) ListTeams(c echo.Context) error {
	teamList := make([]models.Team, len(db.SportsMem))
	for _, t := range db.TeamsMem {
		teamList = append(teamList, *t)
	}

	return c.JSON(http.StatusOK, teamList)
}

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
				Title:     fmt.Sprintf("sport with id %s not found", t.SportID),
				Detail:    "please check the provided ID is correct and try again",
				RequestID: c.Request().Header.Get(echo.HeaderXRequestID),
			},
		}}
		return c.JSON(http.StatusBadRequest, e)
	}

	return c.JSON(http.StatusCreated, core.GenerateTeams(*t, *sport))
}

func (h *Handler) UpdateTeam(c echo.Context) error {
	return nil
}

func (h *Handler) DeleteTeam(c echo.Context) error {
	id := c.Param("id")
	if _, ok := db.TeamsMem[id]; !ok {
		return c.JSON(http.StatusNotFound, models.GenNotFoundError("team", id, c.Get(constants.RequestIDKey).(string)))
	}

	delete(db.TeamsMem, id)
	return c.NoContent(http.StatusNoContent)
}
