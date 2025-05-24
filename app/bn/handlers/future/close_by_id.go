package handlers

import (
	"net/http"
	req "tradething/app/bn/handlers/future/req"
	process "tradething/app/bn/process/future"

	"github.com/labstack/echo/v4"
)

type ICloseByIdHandler interface {
	Handler(c echo.Context) (response interface{}, httpStatus int, err error)
}

type closeByIdHandler struct {
	processCloseById process.IFuture
}

func NewCloseByIdHandler(processCloseById process.IFuture) ICloseByIdHandler {
	return &closeByIdHandler{processCloseById}
}

func (h *closeByIdHandler) Handler(c echo.Context) (response interface{}, httpStatus int, err error) {
	request := new(req.ClosePositionByClientIds)
	if err := c.Bind(request); err != nil {
		return nil, http.StatusBadRequest, err
	}

	response, err = h.processCloseById.ClosePositionByClientIds(c.Request().Context(), request.ClientIds)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return response, http.StatusOK, nil
}
