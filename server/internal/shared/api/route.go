package api

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/riyanamanda/ekinerja/internal/shared/config"
	"github.com/riyanamanda/ekinerja/internal/shared/database"
)

func RouteSetups(app *echo.Echo, cfg *config.Config) {
	_, err := database.GetDatabase(cfg.Database)
	if err != nil {
		log.Fatal("Database connection error: ", err)
	}
}
