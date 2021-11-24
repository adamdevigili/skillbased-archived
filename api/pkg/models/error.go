package models

import (
	"fmt"
	"net/http"

	"github.com/adamdevigili/skillbased/api/pkg/constants"
	"github.com/labstack/echo/v4"
)

// Errors is the collection of errors returned to a caller
type Errors struct {
	Errors []Error `json:"errors"`
}

// Error contains the details of a specific error
type Error struct {
	Status    int    `json:"status"`
	Message   string `json:"message"`
	Detail    string `json:"detail"`
	RequestID string `json:"request_id"`
}

// GenNotFoundError generates a "not found" HTTP error to return to a caller
func GenGenericError(c echo.Context, resource, action string) Error {
	return Error{
		Status:    http.StatusInternalServerError,
		Message:   "internal server error",
		Detail:    "error when fetching sport from database",
		RequestID: c.Get(constants.RequestIDKey).(string),
	}
}

// GenNotFoundError generates a "not found" HTTP error to return to a caller
func GenNotFoundError(c echo.Context, resource, resourceID string) Error {
	return Error{
		Status:    http.StatusNotFound,
		Message:   fmt.Sprintf("%s with ID '%s' not found", resource, resourceID),
		Detail:    "please check the provided ID is correct and try again",
		RequestID: c.Get(constants.RequestIDKey).(string),
	}
}
