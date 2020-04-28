package handlers

import (
	"net/http"

	"github.com/adamdevigili/balancer.team/pkg/db"
	"github.com/adamdevigili/balancer.team/pkg/models"
	"github.com/labstack/echo"
)

func (h *Handler) ListTeams(c echo.Context) error {
	teamList := make([]models.Team, len(db.SportsMem))
	for _, t := range db.TeamsMem {
		teamList = append(teamList, *t)
	}

	return c.JSON(http.StatusOK, teamList)
}
