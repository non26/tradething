package handlers

import (
	"net/http"
	req "tradething/app/bn/handlers/future/req"
	process "tradething/app/bn/process/future"

	"github.com/labstack/echo/v4"
)

type GetAdvancedPositionHandler struct {
	process process.IFuture
}

func NewGetAdvancedPositionHandler(process process.IFuture) *GetAdvancedPositionHandler {
	return &GetAdvancedPositionHandler{process: process}
}

func (h *GetAdvancedPositionHandler) Handler(c echo.Context) error {
	request := new(req.GetAdvancedPositionReq)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	response, err := h.process.GetAdvancedPosition(c.Request().Context(), request.ClientId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, response)
}
