package bidang

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/riyanamanda/ekinerja/internal/features/bidang/dto"
	"github.com/riyanamanda/ekinerja/internal/features/bidang/model"
	"github.com/riyanamanda/ekinerja/internal/shared/response"
	"github.com/riyanamanda/ekinerja/internal/shared/validation"
	"gorm.io/gorm"
)

type bidangHandler struct {
	service model.BidangService
}

func NewBidangHandler(app *echo.Group, bidangService model.BidangService) {
	handler := &bidangHandler{
		service: bidangService,
	}

	app.GET("/bidang", handler.GetAll)
	app.POST("/bidang", handler.Create)
	app.GET("/bidang/:id", handler.GetById)
	app.PUT("/bidang/:id", handler.Update)
	app.DELETE("/bidang/:id", handler.Delete)

}

func (h *bidangHandler) GetAll(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	pageStr := c.QueryParam("page")
	sizeStr := c.QueryParam("size")
	page, size := 1, 10
	if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
		page = p
	}
	if p, err := strconv.Atoi(sizeStr); err == nil && p > 0 {
		size = p
	}
	list, total, err := h.service.GetAll(ctx, page, size)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.CreateErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, response.CreatePaginationResponse(list, page, size, total))
}

func (h *bidangHandler) Create(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	var request dto.BidangRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, response.CreateErrorResponse(err.Error()))
	}

	if validationErrors := validation.Validate(request); len(validationErrors) > 0 {
		return c.JSON(http.StatusBadRequest, response.CreateErrorResponse(validationErrors))
	}

	if err := h.service.Create(ctx, request); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return c.JSON(http.StatusConflict, response.CreateErrorResponse("Bidang dengan nama tersebut sudah ada"))
		}
		return c.JSON(http.StatusInternalServerError, response.CreateErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusCreated, response.CreateSuccessResponse("Bidang berhasil ditambahkan"))
}

func (h *bidangHandler) GetById(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.CreateErrorResponse("invalid id"))
	}
	bidang, err := h.service.GetById(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.CreateErrorResponse(err.Error()))
	}
	if bidang == nil {
		return c.JSON(http.StatusNoContent, map[string]any{})
	}
	return c.JSON(http.StatusOK, bidang)
}

func (h *bidangHandler) Update(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.CreateErrorResponse(err.Error()))
	}
	var request dto.BidangRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, response.CreateErrorResponse(err.Error()))
	}

	if validationErrors := validation.Validate(request); len(validationErrors) > 0 {
		return c.JSON(http.StatusBadRequest, response.CreateErrorResponse(validationErrors))
	}

	if err := h.service.Update(ctx, id, request); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, response.CreateErrorResponse("Bidang tidak ditemukan"))
		}
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return c.JSON(http.StatusConflict, response.CreateErrorResponse("Bidang dengan nama tersebut sudah ada"))
		}
		return c.JSON(http.StatusInternalServerError, response.CreateErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, response.CreateSuccessResponse("Bidang berhasil diupdate"))
}

func (h *bidangHandler) Delete(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.CreateErrorResponse(err.Error()))
	}
	if err := h.service.Delete(ctx, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, response.CreateErrorResponse("Bidang tidak ditemukan"))
		}
		return c.JSON(http.StatusInternalServerError, response.CreateErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, response.CreateSuccessResponse("Bidang berhasil dihapus"))
}
