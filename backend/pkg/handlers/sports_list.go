package handlers

import (
	"net/http"

	"github.com/adamdevigili/skillbased.io/pkg/models"
	"github.com/labstack/gommon/log"

	"github.com/adamdevigili/skillbased.io/pkg/db"
	"github.com/labstack/echo/v4"
)

func (h *Handler) ListSports(c echo.Context) error {
	sportList, err := db.ListSports(h.DBConn)
	if err != nil {
		e := models.Errors{Errors: []models.Error{
			{
				Status:    http.StatusInternalServerError,
				Title:     "internal server error",
				Detail:    "error when storing sport in database",
				RequestID: c.Response().Header().Get(echo.HeaderXRequestID),
			},
		}}
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, e)
	}

	return c.JSON(http.StatusOK, sportList)
}
