package server

import (
	"fmt"
	"net/http"

	"github.com/adamdevigili/skillbased/api/pkg/constants"
	"github.com/adamdevigili/skillbased/api/pkg/handlers"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// InitRoutes sets up the API routes, handlers, and middleware
func InitRoutes(e *echo.Echo, db *gorm.DB, static *http.Handler) {
	h := &handlers.Handler{
		DB: db,
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.RequestID())

	e.GET("/*", echo.WrapHandler(*static))

	apiGroup := e.Group("/v1")

	apiGroup.GET("/health", h.Health)

	// Middleware
	// apiGroup.Use(echomiddleware.Logger())
	// apiGroup.Use(echomiddleware.Recover())
	// apiGroup.Use(echomiddleware.CORS())

	// Teams
	apiGroup.POST("/teams", h.CreateTeam)
	apiGroup.GET(fmt.Sprintf("/teams/:%s", constants.URIKeyID), h.GetTeam)
	apiGroup.DELETE(fmt.Sprintf("/teams/:%s", constants.URIKeyID), h.DeleteTeam)

	apiGroup.PUT("/teams/generate", h.GenerateTeams)

	// Sports
	apiGroup.POST("/sports", h.CreateSport)
	apiGroup.GET("/sports", h.ListSports)
	apiGroup.GET(fmt.Sprintf("/sports/:%s", constants.URIKeyID), h.GetSport)
	apiGroup.DELETE(fmt.Sprintf("/sports/:%s", constants.URIKeyID), h.DeleteSport)
	apiGroup.PATCH(fmt.Sprintf("/sports/:%s", constants.URIKeyID), h.UpdateSport)

	// Players
	apiGroup.GET(fmt.Sprintf("/players/:%s", constants.URIKeyID), h.GetPlayer)
}
