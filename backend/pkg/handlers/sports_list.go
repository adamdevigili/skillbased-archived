package handlers

import (
	"net/http"

	"github.com/adamdevigili/skillbased.io/pkg/db"
	"github.com/adamdevigili/skillbased.io/pkg/models"
	"github.com/labstack/echo/v4"
)

func (h *Handler) ListSports(c echo.Context) error {
	sportList := make([]models.Sport, len(db.SportsMem))
	for _, s := range db.SportsMem {
		sportList = append(sportList, *s)
	}

	return c.JSON(http.StatusOK, sportList)
}
