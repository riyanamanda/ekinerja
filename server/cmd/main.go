package main

import (
	"github.com/labstack/echo/v4"
	"github.com/riyanamanda/ekinerja/internal/shared/api"
	"github.com/riyanamanda/ekinerja/internal/shared/config"
)

func main() {
	cfg := config.Get()
	app := echo.New()

	api.RouteSetups(app, cfg)

	app.Logger.Fatal(
		app.Start(cfg.Server.Host+":"+cfg.Server.Port),
		"Failed to start server",
	)
}
