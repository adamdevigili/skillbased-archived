package models

import (
	"fmt"
	"net/http"
)

type Errors struct {
	Errors []Error `json:"errors"`
}

type Error struct {
	Status    int    `json:"status"`
	Title     string `json:"title"`
	Detail    string `json:"detail"`
	RequestID string `json:"request_id"`
}

func GenNotFoundError(resource, resourceID, requestID string) Error {
	return Error{
		Status:    http.StatusNotFound,
		Title:     fmt.Sprintf("%s with ID '%s' not found", resource, resourceID),
		Detail:    "please check the provided ID is correct and try again",
		RequestID: requestID,
	}
}
