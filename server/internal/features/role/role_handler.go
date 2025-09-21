package role

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/riyanamanda/ekinerja/internal/features/role/model"
	"github.com/riyanamanda/ekinerja/internal/shared/response"
	"gorm.io/gorm"
)

type roleHandler struct {
	service model.RoleService
}

func NewRoleHandler(app *echo.Group, service model.RoleService) {
	handler := &roleHandler{
		service: service,
	}

	app.GET("/role", handler.GetAll)
	app.GET("/role/:id", handler.GetByID)
}

func (h *roleHandler) GetAll(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	roles, err := h.service.GetAll(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.CreateErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, roles)
}

func (h *roleHandler) GetByID(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.CreateErrorResponse("invalid id parameter"))
	}
	role, err := h.service.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNoContent, map[string]any{})
		}
		return c.JSON(http.StatusInternalServerError, response.CreateErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, role)
}
