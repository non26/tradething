package handlers

import (
	"net/http"
	req "tradething/app/bn/handlers/future/req"
	process "tradething/app/bn/process/future"

	"github.com/labstack/echo/v4"
)

type ICloseByIdHandler interface {
	Handler(c echo.Context) error
}

type closeByIdHandler struct {
	processCloseById process.IFuture
}

func NewCloseByIdHandler(processCloseById process.IFuture) ICloseByIdHandler {
	return &closeByIdHandler{processCloseById}
}

func (h *closeByIdHandler) Handler(c echo.Context) error {
	request := new(req.ClosePositionByClientIds)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	response, err := h.processCloseById.ClosePositionByClientIds(c.Request().Context(), request.ClientIds)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, response)
}
