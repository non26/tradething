package handler

import (
	"net/http"
	"tradething/app/bn/future/accumulation/handler/res"
	"tradething/app/bn/future/accumulation/service"

	"github.com/labstack/echo/v4"
	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

type getAllHandler struct {
	service service.IService
}

func NewGetAllHandler(service service.IService) *getAllHandler {
	return &getAllHandler{service: service}
}

func (h *getAllHandler) Handler(c echo.Context) error {
	accumulations, err := h.service.GetAll(c.Request().Context())
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.FailCode, err.Error(), nil)
		return response.SendResponse(http.StatusInternalServerError, c)
	}

	responseData := &res.GetAllRes{}
	responseData.FromDomain(accumulations)
	response := appresponse.NewAppResponse(appresponse.SuccessCode, appresponse.SuccessMsg, responseData)
	return response.SendResponse(http.StatusOK, c)
}
