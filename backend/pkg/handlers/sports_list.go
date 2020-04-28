package handlers

import (
	"net/http"

	"github.com/adamdevigili/balancer.team/pkg/db"
	"github.com/adamdevigili/balancer.team/pkg/models"
	"github.com/labstack/echo"
)

func (h *Handler) ListSports(c echo.Context) error {
	sportList := make([]models.Sport, len(db.SportsMem))
	for _, s := range db.SportsMem {
		sportList = append(sportList, *s)
	}

	return c.JSON(http.StatusOK, sportList)
}
