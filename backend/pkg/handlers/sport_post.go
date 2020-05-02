package handlers

import (
	"net/http"

	"github.com/adamdevigili/skillbased.io/pkg/db"

	"github.com/adamdevigili/skillbased.io/pkg/models"
	"github.com/labstack/echo/v4"
	"github.com/rs/xid"
)

func (h *Handler) CreateSport(c echo.Context) error {
	s := &models.Sport{}

	if err := c.Bind(s); err != nil {
		e := models.Errors{Errors: []models.Error{
			{
				Status:    http.StatusBadRequest,
				Title:     "failed to bind JSON",
				Detail:    "please check your JSON structure",
				RequestID: c.Response().Header().Get(echo.HeaderXRequestID),
			},
		}}
		return c.JSON(http.StatusBadRequest, e)
	}

	xid := xid.New().String()
	s.ID = xid

	db.SportsMem[xid] = s

	return c.JSON(http.StatusCreated, s)
}
