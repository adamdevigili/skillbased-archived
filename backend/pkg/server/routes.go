package server

import (
	"fmt"

	"github.com/adamdevigili/skillbased.io/pkg/constants"
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
	e.GET(fmt.Sprintf("/teams/:%s", constants.URIKeyID), h.GetTeam)
	e.DELETE(fmt.Sprintf("/teams/:%s", constants.URIKeyID), h.DeleteTeam)

	e.POST("/teams/generate", h.GenerateTeams)

	// Sports
	e.POST("/sports", h.CreateSport)
	e.GET("/sports", h.ListSports)
	e.GET(fmt.Sprintf("/sports/:%s", constants.URIKeyID), h.GetSport)
	e.DELETE(fmt.Sprintf("/sports/:%s", constants.URIKeyID), h.DeleteSport)

	// Players
	e.GET(fmt.Sprintf("/players/:%s", constants.URIKeyID), h.GetPlayer)
}
