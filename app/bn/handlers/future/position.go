package handlers

import (
	"net/http"
	req "tradething/app/bn/handlers/future/req"
	process "tradething/app/bn/process/future"

	"github.com/labstack/echo/v4"
	apphandler "github.com/non26/tradepkg/pkg/bn/app_handler"
)

// type IPositionHandler interface {
// 	Handler(c echo.Context) error
// }

type positionHandler struct {
	processPosition process.IFuture
}

func NewPositionHandler(processPosition process.IFuture) apphandler.IHandler {
	return &positionHandler{processPosition}
}

func (p *positionHandler) Handler(c echo.Context) (response interface{}, httpStatus int, err error) {
	request := new(req.Position)
	if err := c.Bind(request); err != nil {
		return nil, http.StatusBadRequest, err
	}
	request.Transform()
	if err := request.Validate(); err != nil {
		return nil, http.StatusBadRequest, err
	}

	response, err = p.processPosition.PlaceOrder(c.Request().Context(), request.ToDomain())
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return response, http.StatusOK, nil
}
