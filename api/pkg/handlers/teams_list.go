package handlers

import (
	"net/http"

	"github.com/adamdevigili/skillbased.io/pkg/db"
	"github.com/adamdevigili/skillbased.io/pkg/models"
	"github.com/labstack/echo/v4"
)

func (h *Handler) ListTeams(c echo.Context) error {
	teamList := make([]models.Team, len(db.SportsMem))
	for _, t := range db.TeamsMem {
		teamList = append(teamList, *t)
	}

	return c.JSON(http.StatusOK, teamList)
}
