package jabatan

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/riyanamanda/ekinerja/internal/features/jabatan/dto"
	"github.com/riyanamanda/ekinerja/internal/features/jabatan/model"
	"github.com/riyanamanda/ekinerja/internal/shared/response"
	"github.com/riyanamanda/ekinerja/internal/shared/validation"
	"gorm.io/gorm"
)

type jabatanHandler struct {
	service model.JabatanService
}

func NewJabatanHandler(app *echo.Group, service model.JabatanService) {
	handler := &jabatanHandler{
		service: service,
	}

	app.GET("/jabatan", handler.GetAll)
	app.POST("/jabatan", handler.Create)
	app.GET("/jabatan/:id", handler.GetById)
	app.PUT("/jabatan/:id", handler.Update)
	app.DELETE("/jabatan/:id", handler.Delete)
}

func (h *jabatanHandler) GetAll(c echo.Context) error {
	_, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	pageStr := c.QueryParam("page")
	perPageStr := c.QueryParam("per_page")
	page := 1
	perPage := 10
	if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
		page = p
	}
	if pp, err := strconv.Atoi(perPageStr); err == nil && pp > 0 {
		perPage = pp
	}
	list, total, err := h.service.GetAll(c.Request().Context(), page, perPage)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	resp := response.CreatePaginationResponse(list, page, perPage, total)
	return c.JSON(http.StatusOK, resp)
}

func (h *jabatanHandler) Create(c echo.Context) error {
	_, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	var request dto.JabatanRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, response.CreateErrorResponse(err.Error()))
	}

	if validationErrors := validation.Validate(request); len(validationErrors) > 0 {
		return c.JSON(http.StatusBadRequest, response.CreateErrorResponse(validationErrors))
	}

	if err := h.service.Save(c.Request().Context(), request); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return c.JSON(http.StatusBadRequest, response.CreateErrorResponse("jabatan dengan nama tersebut sudah ada"))
		}
		return c.JSON(http.StatusInternalServerError, response.CreateErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, response.CreateSuccessResponse("jabatan berhasil ditambahkan"))
}

func (h *jabatanHandler) GetById(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.CreateErrorResponse(err.Error()))
	}
	jabatan, err := h.service.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, response.CreateErrorResponse("jabatan tidak ditemukan"))
		}
		return c.JSON(http.StatusInternalServerError, response.CreateErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, jabatan)
}

func (h *jabatanHandler) Update(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.CreateErrorResponse(err.Error()))
	}
	var request dto.JabatanRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, response.CreateErrorResponse(err.Error()))
	}
	if validationErrors := validation.Validate(request); len(validationErrors) > 0 {
		return c.JSON(http.StatusBadRequest, response.CreateErrorResponse(validationErrors))
	}
	err = h.service.Update(ctx, id, request)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, response.CreateErrorResponse("jabatan tidak ditemukan"))
		} else if errors.Is(err, gorm.ErrDuplicatedKey) {
			return c.JSON(http.StatusConflict, response.CreateErrorResponse("jabatan dengan nama tersebut sudah ada"))
		}
		return c.JSON(http.StatusInternalServerError, response.CreateErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusAccepted, response.CreateSuccessResponse("jabatan berhasil diperbarui"))
}

func (h *jabatanHandler) Delete(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.CreateErrorResponse(err.Error()))
	}
	if err := h.service.Delete(ctx, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, response.CreateErrorResponse("jabatan tidak ditemukan"))
		}
		return c.JSON(http.StatusInternalServerError, response.CreateErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusAccepted, response.CreateSuccessResponse("jabatan berhasil dihapus"))
}
