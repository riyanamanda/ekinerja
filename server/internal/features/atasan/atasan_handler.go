package atasan

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/riyanamanda/ekinerja/internal/features/atasan/dto"
	"github.com/riyanamanda/ekinerja/internal/features/atasan/model"
	"github.com/riyanamanda/ekinerja/internal/shared/response"
	"github.com/riyanamanda/ekinerja/internal/shared/validation"
	"gorm.io/gorm"
)

type atasanHandler struct {
	service model.AtasanService
}

func NewAtasanHandler(app *echo.Group, service model.AtasanService) {
	handler := &atasanHandler{service: service}

	app.GET("/atasan", handler.GetAll)
	app.POST("/atasan", handler.Create)
	app.GET("/atasan/:id", handler.GetByID)
	app.PUT("/atasan/:id", handler.Update)
	app.DELETE("/atasan/:id", handler.Delete)
}

func (h *atasanHandler) GetAll(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	pageStr := c.QueryParam("page")
	sizeStr := c.QueryParam("size")
	page, size := 1, 10
	if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
		page = p
	}
	if s, err := strconv.Atoi(sizeStr); err == nil && s > 0 {
		size = s
	}
	list, total, err := h.service.GetAll(ctx, page, size)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.CreateErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, response.CreatePaginationResponse(list, page, size, total))
}

func (h *atasanHandler) Create(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	var request dto.AtasanCreateRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, response.CreateErrorResponse(err.Error()))
	}

	if validationErrors := validation.Validate(request); len(validationErrors) > 0 {
		return c.JSON(http.StatusBadRequest, response.CreateErrorResponse(validationErrors))
	}

	if err := h.service.Create(ctx, &request); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return c.JSON(http.StatusConflict, response.CreateErrorResponse("Atasan dengan nama yang sama sudah ada"))
		}
		return c.JSON(http.StatusInternalServerError, response.CreateErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusCreated, response.CreateSuccessResponse("Atasan berhasil ditambahkan"))
}

func (h *atasanHandler) GetByID(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.CreateErrorResponse("Invalid ID"))
	}

	atasan, err := h.service.GetByID(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.CreateErrorResponse(err.Error()))
	}
	if atasan == nil {
		return c.JSON(http.StatusOK, map[string]any{})
	}
	return c.JSON(http.StatusOK, atasan)
}

func (h *atasanHandler) Update(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.CreateErrorResponse("Invalid ID"))
	}

	var request dto.AtasanUpdateRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, response.CreateErrorResponse(err.Error()))
	}

	if validationErrors := validation.Validate(request); len(validationErrors) > 0 {
		return c.JSON(http.StatusBadRequest, response.CreateErrorResponse(validationErrors))
	}

	if err := h.service.Update(ctx, id, request); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, response.CreateErrorResponse("Atasan tidak ditemukan"))
		} else if errors.Is(err, gorm.ErrDuplicatedKey) {
			return c.JSON(http.StatusConflict, response.CreateErrorResponse("Atasan dengan nama yang sama sudah ada"))
		}
		return c.JSON(http.StatusInternalServerError, response.CreateErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, response.CreateSuccessResponse("Atasan berhasil diperbarui"))
}

func (h *atasanHandler) Delete(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.CreateErrorResponse("Invalid ID"))
	}

	if err := h.service.Delete(ctx, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, response.CreateErrorResponse("Atasan tidak ditemukan"))
		}
		return c.JSON(http.StatusInternalServerError, response.CreateErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, response.CreateSuccessResponse("Atasan berhasil dihapus"))
}
