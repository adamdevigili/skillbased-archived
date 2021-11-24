package handlers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func getRequestID(c echo.Context) string {
	return c.Response().Header().Get(echo.HeaderXRequestID)
}
