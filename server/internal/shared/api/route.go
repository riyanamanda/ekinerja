package api

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/riyanamanda/ekinerja/internal/features/bidang"
	"github.com/riyanamanda/ekinerja/internal/features/jabatan"
	"github.com/riyanamanda/ekinerja/internal/features/pangkat"
	"github.com/riyanamanda/ekinerja/internal/shared/config"
	"github.com/riyanamanda/ekinerja/internal/shared/database"
)

func RouteSetups(app *echo.Echo, cfg *config.Config) {
	conn, err := database.GetDatabase(cfg.Database)
	if err != nil {
		log.Fatal("Database connection error: ", err)
	}

	api := app.Group("/api")

	pangkatRepository := pangkat.NewPangkatRepository(conn)
	pangkatService := pangkat.NewPangkatService(pangkatRepository)
	pangkat.NewPangkatHandler(api, pangkatService)

	jabatanRepository := jabatan.NewJabatanRepository(conn)
	jabatanService := jabatan.NewJabatanService(jabatanRepository)
	jabatan.NewJabatanHandler(api, jabatanService)

	bidangRepository := bidang.NewBidangRepository(conn)
	bidangService := bidang.NewBidangService(bidangRepository)
	bidang.NewBidangHandler(api, bidangService)
}
