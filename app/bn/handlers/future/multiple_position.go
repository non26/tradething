package handlers

import (
	"net/http"
	req "tradething/app/bn/handlers/future/req"
	process "tradething/app/bn/process/future"

	"github.com/labstack/echo/v4"
)

type IMultiplePositionHandler interface {
	Handler(c echo.Context) error
}

type multiplePositionHandler struct {
	processMultiplePosition process.IFuture
}

func NewMultiplePositionHandler(processMultiplePosition process.IFuture) IMultiplePositionHandler {
	return &multiplePositionHandler{processMultiplePosition}
}

func (h *multiplePositionHandler) Handler(c echo.Context) error {
	request := new(req.MultiplePosition)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	response, err := h.processMultiplePosition.MultiplePosition(c.Request().Context(), request.ToDomain())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, response)
}
