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
	h := &handlers.Handler{
		DBConn: dbConn,
	}
	e.GET("/health", h.Health)

	apiGroup := e.Group("/v1")

	// Middleware
	apiGroup.Use(middleware.RequestIDMiddleware())
	apiGroup.Use(echomiddleware.Logger())
	apiGroup.Use(echomiddleware.Recover())

	// Teams
	apiGroup.POST("/teams", h.CreateTeam)
	apiGroup.GET(fmt.Sprintf("/teams/:%s", constants.URIKeyID), h.GetTeam)
	apiGroup.DELETE(fmt.Sprintf("/teams/:%s", constants.URIKeyID), h.DeleteTeam)

	apiGroup.POST("/teams/generate", h.GenerateTeams)

	// Sports
	apiGroup.POST("/sports", h.CreateSport)
	apiGroup.GET("/sports", h.ListSports)
	apiGroup.GET(fmt.Sprintf("/sports/:%s", constants.URIKeyID), h.GetSport)
	apiGroup.DELETE(fmt.Sprintf("/sports/:%s", constants.URIKeyID), h.DeleteSport)

	// Players
	apiGroup.GET(fmt.Sprintf("/players/:%s", constants.URIKeyID), h.GetPlayer)
}