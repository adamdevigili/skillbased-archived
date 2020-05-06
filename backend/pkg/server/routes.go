package server

import (
	"github.com/adamdevigili/skillbased.io/pkg/middleware"
	"github.com/jackc/pgx"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"

	"github.com/adamdevigili/skillbased.io/pkg/handlers"
)

func InitRoutes(e *echo.Echo, dbConn *pgx.ConnPool) {
	// Middleware
	e.Use(middleware.RequestIDMiddleware())
	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.Recover())

	h := &handlers.Handler{
		DBConn: dbConn,
	}

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
