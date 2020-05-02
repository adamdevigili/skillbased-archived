package handlers

import (
	"github.com/adamdevigili/skillbased.io/pkg/db"
	"github.com/adamdevigili/skillbased.io/pkg/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/rs/xid"

	"net/http"
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
