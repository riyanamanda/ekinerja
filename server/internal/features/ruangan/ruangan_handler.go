package ruangan

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/riyanamanda/ekinerja/internal/features/ruangan/dto"
	"github.com/riyanamanda/ekinerja/internal/features/ruangan/model"
	"github.com/riyanamanda/ekinerja/internal/shared/response"
	"github.com/riyanamanda/ekinerja/internal/shared/validation"
	"gorm.io/gorm"
)

type ruanganHandler struct {
	service model.RuanganService
}

func NewRuanganHandler(app *echo.Group, service model.RuanganService) {
	handler := &ruanganHandler{service: service}

	app.GET("/ruangan", handler.GetAll)
	app.POST("/ruangan", handler.Create)
	app.GET("/ruangan/:id", handler.GetByID)
	app.PUT("/ruangan/:id", handler.Update)
	app.DELETE("/ruangan/:id", handler.Delete)
}

func (h *ruanganHandler) GetAll(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	pageStr := c.QueryParam("page")
	sizeStr := c.QueryParam("size")
	page, size := 1, 10
	if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
		page = p
	}
	if pp, err := strconv.Atoi(sizeStr); err == nil && pp > 0 {
		size = pp
	}
	list, total, err := h.service.GetAll(ctx, page, size)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.CreateErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, response.CreatePaginationResponse(list, page, size, total))
}

func (h *ruanganHandler) Create(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	var request dto.RuanganRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, response.CreateErrorResponse(err.Error()))
	}

	if validationErrors := validation.Validate(request); len(validationErrors) > 0 {
		return c.JSON(http.StatusBadRequest, response.CreateErrorResponse(validationErrors))
	}

	if err := h.service.Create(ctx, request); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return c.JSON(http.StatusConflict, response.CreateErrorResponse("Ruangan dengan nama tersebut sudah ada"))
		}
		return c.JSON(http.StatusInternalServerError, response.CreateErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, response.CreateSuccessResponse("Ruangan berhasil ditambahkan"))
}

func (h *ruanganHandler) GetByID(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.CreateErrorResponse("Invalid ID"))
	}

	ruangan, err := h.service.GetById(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.CreateErrorResponse(err.Error()))
	}
	if ruangan == nil {
		return c.JSON(http.StatusNoContent, map[string]any{})
	}
	return c.JSON(http.StatusOK, ruangan)
}

func (h *ruanganHandler) Update(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.CreateErrorResponse("Invalid ID"))
	}

	var request dto.RuanganRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, response.CreateErrorResponse(err.Error()))
	}

	if validationErrors := validation.Validate(request); len(validationErrors) > 0 {
		return c.JSON(http.StatusBadRequest, response.CreateErrorResponse(validationErrors))
	}

	if err := h.service.Update(ctx, id, request); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, response.CreateErrorResponse("Ruangan tidak ditemukan"))
		} else if errors.Is(err, gorm.ErrDuplicatedKey) {
			return c.JSON(http.StatusConflict, response.CreateErrorResponse("Ruangan dengan nama tersebut sudah ada"))
		}
		return c.JSON(http.StatusInternalServerError, response.CreateErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusAccepted, response.CreateSuccessResponse("Ruangan berhasil diperbarui"))
}

func (h *ruanganHandler) Delete(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.CreateErrorResponse("Invalid ID"))
	}

	if err := h.service.Delete(ctx, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, response.CreateErrorResponse("Ruangan tidak ditemukan"))
		}
		return c.JSON(http.StatusInternalServerError, response.CreateErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, nil)
}
