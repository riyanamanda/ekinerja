package role

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/riyanamanda/ekinerja/internal/features/role/dto"
	"github.com/riyanamanda/ekinerja/internal/features/role/model"
	"github.com/riyanamanda/ekinerja/internal/shared/response"
	"github.com/riyanamanda/ekinerja/internal/shared/validation"
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
	app.POST("/role", handler.Create)
	app.GET("/role/:id", handler.GetByID)
	app.PUT("/role/:id", handler.Update)
	app.DELETE("/role/:id", handler.Delete)
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

func (h *roleHandler) Create(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	var request dto.RoleRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, response.CreateErrorResponse("invalid request"))
	}

	if validationErrors := validation.Validate(request); len(validationErrors) > 0 {
		return c.JSON(http.StatusBadRequest, response.CreateErrorResponse(validationErrors))
	}

	if err := h.service.Create(ctx, request); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return c.JSON(http.StatusConflict, response.CreateErrorResponse("role dengan nama tersebut sudah ada"))
		}
		return c.JSON(http.StatusInternalServerError, response.CreateErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusCreated, response.CreateSuccessResponse("role berhasil ditambahkan"))
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

func (h *roleHandler) Update(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.CreateErrorResponse("invalid id parameter"))
	}

	var request dto.RoleRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, response.CreateErrorResponse("invalid request"))
	}

	if validationErrors := validation.Validate(request); len(validationErrors) > 0 {
		return c.JSON(http.StatusBadRequest, response.CreateErrorResponse(validationErrors))
	}

	if err = h.service.Update(ctx, id, request); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, response.CreateErrorResponse("role tidak ditemukan"))
		}
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return c.JSON(http.StatusConflict, response.CreateErrorResponse("role dengan nama tersebut sudah ada"))
		}
		return c.JSON(http.StatusInternalServerError, response.CreateErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, response.CreateSuccessResponse("role berhasil diperbarui"))
}

func (h *roleHandler) Delete(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.CreateErrorResponse("invalid id parameter"))
	}

	if err = h.service.Delete(ctx, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, response.CreateErrorResponse("role tidak ditemukan"))
		}
		return c.JSON(http.StatusInternalServerError, response.CreateErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, response.CreateSuccessResponse("role berhasil dihapus"))
}
