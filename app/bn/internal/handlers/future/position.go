package handlers

import (
	"net/http"
	req "tradething/app/bn/internal/handlers/future/req"
	process "tradething/app/bn/internal/process/future"

	"github.com/labstack/echo/v4"
)

type IPositionHandler interface {
	Handler(c echo.Context) error
}

type positionHandler struct {
	processPosition process.IFuture
}

func NewPositionHandler(processPosition process.IFuture) IPositionHandler {
	return &positionHandler{processPosition}
}

func (p *positionHandler) Handler(c echo.Context) error {
	request := new(req.Position)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	request.Transform()
	if err := request.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	response, err := p.processPosition.PlaceOrder(c.Request().Context(), request.ToDomain())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, response)
}
