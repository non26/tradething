package handler

import (
	"net/http"
	res "tradething/app/bn/future/position/handler/res/advanced_position"
	"tradething/app/bn/future/position/service"

	"github.com/labstack/echo/v4"
	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

type getAllHandler struct {
	service service.IAdvancedPositionService
}

func NewGetAllHandler(service service.IAdvancedPositionService) *getAllHandler {
	return &getAllHandler{service: service}
}

func (h *getAllHandler) Handler(c echo.Context) error {
	positions, err := h.service.GetAll(c.Request().Context())
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.FailCode, err.Error(), nil)
		return response.SendResponse(http.StatusInternalServerError, c)
	}

	responseData := &res.GetAllAdvancedPositionRes{}
	responseData = responseData.FromDomain(positions)
	var response *appresponse.AppResponse
	if responseData == nil {
		response = appresponse.NewAppResponse(appresponse.SuccessCode, appresponse.SuccessMsg, nil)
	} else {
		response = appresponse.NewAppResponse(appresponse.SuccessCode, appresponse.SuccessMsg, responseData)
	}
	return response.SendResponse(http.StatusOK, c)
}
