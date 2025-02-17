package handlers

import (
	"net/http"
	"tradething/app/bn/handlers/spot/req"
	process "tradething/app/bn/process/spot"

	"github.com/labstack/echo/v4"
)

type ICloseByIdHandler interface {
	Handler(c echo.Context) error
}

type closeByIdHandler struct {
	processCloseById process.ISpot
}

func NewCloseByIdHandler(
	processCloseById process.ISpot,
) ICloseByIdHandler {
	return &closeByIdHandler{processCloseById}
}

func (h *closeByIdHandler) Handler(c echo.Context) error {
	request := req.CloseByClientIds{}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	response, err := h.processCloseById.CloseOrderByClientIds(c.Request().Context(), request.ClientIds)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, response)
}
