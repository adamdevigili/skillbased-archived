package server

import (
	"github.com/adamdevigili/balancer.team/pkg/middleware"
	"github.com/labstack/echo"
	echomiddleware "github.com/labstack/echo/middleware"

	"github.com/adamdevigili/balancer.team/pkg/handlers"
)

func InitRoutes(e *echo.Echo) {
	// Middleware
	e.Use(middleware.RequestIDMiddleware())
	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.Recover())

	h := &handlers.Handler{}

	// Teams
	e.POST("/teams", h.CreateTeam)
	e.GET("/teams/:id", h.GetTeam)
	e.DELETE("/teams/:id", h.DeleteTeam)

	e.POST("/teams/generate", h.GenerateTeams)

	// Sports
	e.GET("/sports", h.ListSports)
	e.POST("/sports", h.CreateSport)

	// Players
	e.GET("/players/:id", h.GetPlayer)

}
