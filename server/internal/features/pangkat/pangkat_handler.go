package pangkat

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/riyanamanda/ekinerja/internal/features/pangkat/dto"
)

type pangkatHandler struct {
	service PangkatService
}

func NewPangkatHandler(app *echo.Group, service PangkatService) {
	Handler := &pangkatHandler{service: service}

	app.GET("/pangkat", Handler.GetAll)
	app.POST("/pangkat", Handler.Save)
}

func (h *pangkatHandler) GetAll(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	responses, err := h.service.GetAll(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, responses)
}

func (h *pangkatHandler) Save(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	var request dto.PangkatRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.service.Save(ctx, request); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "success"})
}
