package handlers

import (
	"net/http"
	req "tradething/app/bn/handlers/future/req"
	process "tradething/app/bn/process/future"

	"github.com/labstack/echo/v4"
)

type ManagePositionHandler struct {
	process process.IFuture
}

func NewManagePositionHandler(process process.IFuture) *ManagePositionHandler {
	return &ManagePositionHandler{process: process}
}

func (h *ManagePositionHandler) Handler(c echo.Context) (response interface{}, httpStatus int, err error) {
	request := new(req.ManagePositionReq)
	if err := c.Bind(request); err != nil {
		return nil, http.StatusBadRequest, err
	}

	response, err = h.process.ManagePosition(c.Request().Context(), request.ClosePosition, request.AdvancedPosition)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return response, http.StatusOK, nil
}
