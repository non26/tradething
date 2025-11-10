package handler

import (
	"net/http"
	res "tradething/app/bn/future/position/handler/res/current_position"
	"tradething/app/bn/future/position/service"

	"github.com/labstack/echo/v4"
	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

type getallHandler struct {
	service service.ICurrentPositionService
}

func NewGetAllHandler(service service.ICurrentPositionService) *getallHandler {
	return &getallHandler{service: service}
}

func (h *getallHandler) Handler(c echo.Context) error {
	positions, err := h.service.GetAll(c.Request().Context())
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.FailCode, err.Error(), nil)
		return response.SendResponse(http.StatusInternalServerError, c)
	}

	responseData := &res.GetAllCurrentPositionRes{}
	responseData = responseData.FromDomain(positions)
	response := appresponse.NewAppResponse(appresponse.SuccessCode, appresponse.SuccessMsg, responseData)
	return response.SendResponse(http.StatusOK, c)
}
