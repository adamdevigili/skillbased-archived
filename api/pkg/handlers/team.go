package handlers

import (
	"fmt"
	"net/http"

	"github.com/adamdevigili/skillbased/api/pkg/constants"
	"github.com/adamdevigili/skillbased/api/pkg/core"
	"github.com/adamdevigili/skillbased/api/pkg/db"
	"github.com/adamdevigili/skillbased/api/pkg/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// CreateTeam creates a team
func (h *Handler) CreateTeam(c echo.Context) error {
	t := &models.Team{}

	//if err := c.Bind(t); err != nil {
	//	e := models.Errors{Errors: []models.Error{
	//		{
	//			Status:    http.StatusBadRequest,
	//			Title:     "failed to bind JSON",
	//			Detail:    "please check your JSON structure",
	//			RequestID: c.Request().Header.Get(echo.HeaderXRequestID),
	//		},
	//	}}
	//	log.Error("ded")
	//	return c.JSON(http.StatusBadRequest, e)
	//}
	//
	//xid := xid.New().String()
	//t.ID = xid
	//
	//db.TeamsMem[xid] = t

	return c.JSON(http.StatusCreated, t)
}

// GetTeam retrieves an existing sport
func (h *Handler) GetTeam(c echo.Context) error {
	id := c.Param(constants.URIKeyID)
	team, ok := db.TeamsMem[id]
	if !ok {
		return c.JSON(http.StatusNotFound, models.GenNotFoundError(c, "team", id))
	}

	return c.JSON(http.StatusOK, team)
}

// ListSports list all existing teams
func (h *Handler) ListTeams(c echo.Context) error {
	teamList, err := db.ListTeams(h.DB)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, models.GenGenericError(c, "teams", "listing"))
	}

	return c.JSON(http.StatusOK, teamList)
}

// GenerateTeams generates a TeamSet based on the provided players and target sport
func (h *Handler) GenerateTeams(c echo.Context) error {
	t := &models.GenerateTeamRequest{}

	if err := c.Bind(t); err != nil {
		e := models.Errors{Errors: []models.Error{
			{
				Status:    http.StatusBadRequest,
				Message:   "failed to bind JSON",
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
				Message:   fmt.Sprintf("sport with id %s not found", t.SportID),
				Detail:    "please check the provided ID is correct and try again",
				RequestID: c.Request().Header.Get(echo.HeaderXRequestID),
			},
		}}
		return c.JSON(http.StatusBadRequest, e)
	}

	return c.JSON(http.StatusCreated, core.GenerateTeams(*t, *sport))
}

// UpdateTeam updates an existing team
func (h *Handler) UpdateTeam(c echo.Context) error {
	return nil
}

// DeleteTeam deletes an existing team
func (h *Handler) DeleteTeam(c echo.Context) error {
	id := c.Param("id")
	if _, ok := db.TeamsMem[id]; !ok {
		return c.JSON(http.StatusNotFound, models.GenNotFoundError(c, "team", id))
	}

	delete(db.TeamsMem, id)
	return c.NoContent(http.StatusNoContent)
}
