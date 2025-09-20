package pangkat

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/riyanamanda/ekinerja/internal/features/pangkat/dto"
	"github.com/riyanamanda/ekinerja/internal/features/pangkat/model"
	"github.com/riyanamanda/ekinerja/internal/shared/response"
	"github.com/riyanamanda/ekinerja/internal/shared/validation"
	"gorm.io/gorm"
)

type pangkatHandler struct {
	service model.PangkatService
}

func NewPangkatHandler(app *echo.Group, service model.PangkatService) {
	Handler := &pangkatHandler{service: service}

	app.GET("/pangkat", Handler.GetAll)
	app.POST("/pangkat", Handler.Create)
	app.GET("/pangkat/:id", Handler.GetById)
	app.PUT("/pangkat/:id", Handler.Update)
	app.DELETE("/pangkat/:id", Handler.Delete)
}

func (h *pangkatHandler) GetAll(c echo.Context) error {
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

func (h *pangkatHandler) Create(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	var request dto.PangkatRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, response.CreateErrorResponse(err.Error()))
	}

	if validationErrors := validation.Validate(request); len(validationErrors) > 0 {
		return c.JSON(http.StatusBadRequest, response.CreateErrorResponse(validationErrors))
	}

	if err := h.service.Create(ctx, request); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return c.JSON(http.StatusConflict, response.CreateErrorResponse("Pangkat dengan nama tersebut sudah ada"))
		}
		return c.JSON(http.StatusInternalServerError, response.CreateErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusCreated, response.CreateSuccessResponse("Pangkat berhasil ditambahkan"))
}

func (h *pangkatHandler) GetById(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.CreateErrorResponse("invalid id"))
	}
	pangkat, err := h.service.GetById(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.CreateErrorResponse(err.Error()))
	}
	if pangkat == nil {
		return c.JSON(http.StatusNoContent, map[string]any{})
	}
	return c.JSON(http.StatusOK, pangkat)
}

func (h *pangkatHandler) Update(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.CreateErrorResponse("invalid id"))
	}

	var request dto.PangkatRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, response.CreateErrorResponse(err.Error()))
	}

	if validationErrors := validation.Validate(request); len(validationErrors) > 0 {
		return c.JSON(http.StatusBadRequest, response.CreateErrorResponse(validationErrors))
	}

	if err = h.service.Update(ctx, id, request); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, response.CreateErrorResponse("Pangkat tidak ditemukan"))
		} else if errors.Is(err, gorm.ErrDuplicatedKey) {
			return c.JSON(http.StatusConflict, response.CreateErrorResponse("Pangkat dengan nama tersebut sudah ada"))
		}
		return c.JSON(http.StatusInternalServerError, response.CreateErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusAccepted, response.CreateSuccessResponse("Pangkat berhasil diperbaharui"))
}

func (h *pangkatHandler) Delete(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.CreateErrorResponse("invalid id"))
	}

	if err = h.service.Delete(ctx, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, response.CreateErrorResponse("Pangkat tidak ditemukan"))
		}
		return c.JSON(http.StatusInternalServerError, response.CreateErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, response.CreateSuccessResponse("Pangkat berhasil dihapus"))
}
