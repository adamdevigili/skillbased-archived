package server

import (
	"fmt"

	"github.com/adamdevigili/skillbased/api/pkg/constants"
	"github.com/adamdevigili/skillbased/api/pkg/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

// InitRoutes sets up the API routes, handlers, and middleware
func InitRoutes(e *echo.Echo, db *gorm.DB) {
	h := &handlers.Handler{
		DB: db,
	}

	// api := e.Group("/api")

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.RequestID())

	v1 := e.Group("/v1")

	v1.GET("/health", h.Health)

	// Teams
	v1.POST("/teams", h.CreateTeam)
	v1.GET(fmt.Sprintf("/teams/:%s", constants.URIKeyID), h.GetTeam)
	v1.GET("/teams", h.ListTeams)
	v1.DELETE(fmt.Sprintf("/teams/:%s", constants.URIKeyID), h.DeleteTeam)

	v1.PUT("/teams/generate", h.GenerateTeams)

	// Sports
	v1.POST("/sports", h.CreateSport)
	v1.GET("/sports", h.ListSports)
	v1.GET(fmt.Sprintf("/sports/:%s", constants.URIKeyID), h.GetSport)
	v1.DELETE(fmt.Sprintf("/sports/:%s", constants.URIKeyID), h.DeleteSport)
	v1.PATCH(fmt.Sprintf("/sports/:%s", constants.URIKeyID), h.UpdateSport)

	// Players
	v1.GET(fmt.Sprintf("/players/:%s", constants.URIKeyID), h.GetPlayer)
}
