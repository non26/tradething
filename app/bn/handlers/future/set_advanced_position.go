package handlers

import (
	"net/http"
	req "tradething/app/bn/handlers/future/req"
	process "tradething/app/bn/process/future"

	"github.com/labstack/echo/v4"
)

type SetAdvancedPositionHandler struct {
	process process.IFuture
}

func NewSetAdvancedPositionHandler(process process.IFuture) *SetAdvancedPositionHandler {
	return &SetAdvancedPositionHandler{process: process}
}

func (h *SetAdvancedPositionHandler) Handler(c echo.Context) error {
	request := new(req.SetAdvancedPositionReqs)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := request.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	response, err := h.process.SetAdvancedPosition(c.Request().Context(), request.ToDomain())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, response)
}
