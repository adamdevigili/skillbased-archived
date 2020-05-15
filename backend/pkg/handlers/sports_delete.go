package handlers

import (
	"net/http"

	"github.com/adamdevigili/skillbased.io/pkg/constants"

	"github.com/adamdevigili/skillbased.io/pkg/db"
	"github.com/labstack/echo/v4"
)

func (h *Handler) DeleteSport(c echo.Context) error {
	db.DeleteSport(h.DBConn, c.Param(constants.URIKeyID))

	return c.NoContent(http.StatusNoContent)
}
