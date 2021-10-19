package models

import (
	"fmt"
	"net/http"
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
func GenNotFoundError(resource, resourceID, requestID string) Error {
	return Error{
		Status:    http.StatusNotFound,
		Message:   fmt.Sprintf("%s with ID '%s' not found", resource, resourceID),
		Detail:    "please check the provided ID is correct and try again",
		RequestID: requestID,
	}
}
