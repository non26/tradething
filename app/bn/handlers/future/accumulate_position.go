package handlers

import (
	"net/http"
	req "tradething/app/bn/handlers/future/req"
	process "tradething/app/bn/process/future"

	"github.com/labstack/echo/v4"
)

type IAccumulatePositionHandler interface {
	Handler(c echo.Context) (response interface{}, httpStatus int, err error)
}

type accumulatePositionHandler struct {
	accumulatePosition process.IFuture
}

func NewAccumulatePositionHandler(process process.IFuture) IAccumulatePositionHandler {
	return &accumulatePositionHandler{process}
}

func (h *accumulatePositionHandler) Handler(c echo.Context) (response interface{}, httpStatus int, err error) {
	request := new(req.AccumulatePositionReq)
	if err := c.Bind(request); err != nil {
		return nil, http.StatusBadRequest, err
	}

	response, err = h.accumulatePosition.AccumulatePosition(c.Request().Context(), request.ToDomain())
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return response, http.StatusOK, nil
}
