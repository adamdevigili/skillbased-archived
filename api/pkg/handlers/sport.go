package handlers

import (
	"net/http"
	"strings"

	"github.com/adamdevigili/skillbased.io/pkg/constants"

	"github.com/rs/xid"

	"github.com/adamdevigili/skillbased.io/pkg/db"
	"github.com/adamdevigili/skillbased.io/pkg/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
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

	s.ID = xid.New()

	if err := db.InsertSport(h.DB, s); err != nil {
		e := models.Errors{Errors: []models.Error{
			{
				Status:    http.StatusInternalServerError,
				Title:     "internal server error",
				Detail:    "error when storing sport in database",
				RequestID: c.Response().Header().Get(echo.HeaderXRequestID),
			},
		}}

		log.Errorf("Unable to create sport: %v", err)
		return c.JSON(http.StatusInternalServerError, e)
	}

	return c.JSON(http.StatusCreated, s)
}

func (h *Handler) GetSport(c echo.Context) error {
	id := c.Param(constants.URIKeyID)

	sport, err := db.GetSport(h.DB, id)
	if err != nil {
		log.Errorf("Unable to fetch sport: %v", err)

		code := http.StatusInternalServerError
		e := models.Errors{Errors: []models.Error{}}

		if strings.Contains(err.Error(), "not found") {
			e.Errors = append(e.Errors, models.GenNotFoundError("sport", id, c.Get(constants.RequestIDKey).(string)))
			code = http.StatusNotFound
		} else {
			e.Errors = append(e.Errors, models.Error{
				Status:    http.StatusInternalServerError,
				Title:     "internal server error",
				Detail:    "error when fetching sport from database",
				RequestID: c.Response().Header().Get(echo.HeaderXRequestID),
			})
		}

		return c.JSON(code, e)
	}

	return c.JSON(http.StatusOK, sport)
}

func (h *Handler) DeleteSport(c echo.Context) error {
	id := c.Param(constants.URIKeyID)

	if err := db.DeleteSport(h.DB, id); err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(
				http.StatusNotFound,
				models.GenNotFoundError("sport", id, c.Get(constants.RequestIDKey).(string)),
			)
		}

		e := models.Errors{Errors: []models.Error{
			{
				Status:    http.StatusInternalServerError,
				Title:     "internal server error",
				Detail:    "error when deleting sport from database",
				RequestID: c.Response().Header().Get(echo.HeaderXRequestID),
			},
		}}

		log.Errorf("Unable to fetch sport: %v", err)
		return c.JSON(http.StatusInternalServerError, e)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) ListSports(c echo.Context) error {
	sportList, err := db.ListSports(h.DB)
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
