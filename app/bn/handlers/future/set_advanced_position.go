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

func (h *SetAdvancedPositionHandler) Handler(c echo.Context) (response interface{}, httpStatus int, err error) {
	request := new(req.SetAdvancedPositionReqs)
	if err := c.Bind(request); err != nil {
		return nil, http.StatusBadRequest, err
	}

	if err := request.Validate(); err != nil {
		return nil, http.StatusBadRequest, err
	}

	response, err = h.process.SetAdvancedPosition(c.Request().Context(), request.ToDomain())
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return response, http.StatusOK, nil
}
