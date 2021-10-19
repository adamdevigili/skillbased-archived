package handlers

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type (
	Handler struct {
		DB *gorm.DB
	}
)

func getRequestID(c echo.Context) string {
	return c.Response().Header().Get(echo.HeaderXRequestID)
}
