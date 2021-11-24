package handlers

import (
	"net/http"
	"strings"

	"github.com/adamdevigili/skillbased/api/pkg/constants"
	"github.com/adamdevigili/skillbased/api/pkg/db"
	"github.com/adamdevigili/skillbased/api/pkg/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/segmentio/ksuid"
)

// CreateSport creates a sport
func (h *Handler) CreateSport(c echo.Context) error {
	s := &models.Sport{}

	if err := c.Bind(s); err != nil {
		e := models.Errors{Errors: []models.Error{
			{
				Status:    http.StatusBadRequest,
				Message:   "failed to bind JSON",
				Detail:    "please check your JSON structure",
				RequestID: c.Response().Header().Get(echo.HeaderXRequestID),
			},
		}}
		return c.JSON(http.StatusBadRequest, e)
	}

	s.ID = ksuid.New().String()

	if err := db.InsertSport(h.DB, s); err != nil {
		e := models.Errors{Errors: []models.Error{
			{
				Status:    http.StatusInternalServerError,
				Message:   "internal server error",
				Detail:    "error when storing sport in database",
				RequestID: c.Response().Header().Get(echo.HeaderXRequestID),
			},
		}}

		log.Errorf("Unable to create sport: %v", err)
		return c.JSON(http.StatusInternalServerError, e)
	}

	return c.JSON(http.StatusCreated, s)
}

// GetSport retrieves an existing sport
func (h *Handler) GetSport(c echo.Context) error {
	id := c.Param(constants.URIKeyID)
	log.Infof("Fetching sport: %s", id)

	sport, err := db.GetSport(h.DB, id)
	if err != nil {
		log.Errorf("Unable to fetch sport: %v", err)

		code := http.StatusInternalServerError
		e := models.Errors{Errors: []models.Error{}}

		if strings.Contains(err.Error(), "not found") {
			e.Errors = append(e.Errors, models.GenNotFoundError("sport", id, getRequestID(c)))
			code = http.StatusNotFound
		} else {
			e.Errors = append(e.Errors, models.Error{
				Status:    http.StatusInternalServerError,
				Message:   "internal server error",
				Detail:    "error when fetching sport from database",
				RequestID: c.Response().Header().Get(echo.HeaderXRequestID),
			})
		}

		return c.JSON(code, e)
	}

	return c.JSON(http.StatusOK, sport)
}

// ListSports list all existing sports
func (h *Handler) ListSports(c echo.Context) error {
	sportList, err := db.ListSports(h.DB)
	if err != nil {
		e := models.Errors{Errors: []models.Error{
			{
				Status:    http.StatusInternalServerError,
				Message:   "internal server error",
				Detail:    "error when storing sport in database",
				RequestID: c.Response().Header().Get(echo.HeaderXRequestID),
			},
		}}
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, e)
	}

	return c.JSON(http.StatusOK, sportList)
}

// UpdateSport updates an existing sport
func (h *Handler) UpdateSport(c echo.Context) error {
	s := &models.Sport{}

	if err := c.Bind(s); err != nil {
		e := models.Errors{Errors: []models.Error{
			{
				Status:    http.StatusBadRequest,
				Message:   "failed to bind JSON",
				Detail:    "please check your JSON structure",
				RequestID: c.Response().Header().Get(echo.HeaderXRequestID),
			},
		}}
		return c.JSON(http.StatusBadRequest, e)
	}

	id := c.Param(constants.URIKeyID)

	s.ID = id

	sport, err := db.UpdateSport(h.DB, s)
	if err != nil {
		log.Errorf("Unable to update sport: %v", err)

		code := http.StatusInternalServerError
		e := models.Errors{Errors: []models.Error{}}

		if strings.Contains(err.Error(), "not found") {
			e.Errors = append(e.Errors, models.GenNotFoundError("sport", id, c.Get(constants.RequestIDKey).(string)))
			code = http.StatusNotFound
		} else {
			e.Errors = append(e.Errors, models.Error{
				Status:    http.StatusInternalServerError,
				Message:   "internal server error",
				Detail:    "error when fetching sport from database",
				RequestID: c.Response().Header().Get(echo.HeaderXRequestID),
			})
		}

		return c.JSON(code, e)
	}

	return c.JSON(http.StatusOK, sport)
}

// DeleteSport deletes an existing sport
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
				Message:   "internal server error",
				Detail:    "error when deleting sport from database",
				RequestID: c.Response().Header().Get(echo.HeaderXRequestID),
			},
		}}

		log.Errorf("Unable to fetch sport: %v", err)
		return c.JSON(http.StatusInternalServerError, e)
	}

	return c.NoContent(http.StatusNoContent)
}
