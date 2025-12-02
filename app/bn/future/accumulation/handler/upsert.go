package handler

import (
	"net/http"
	"tradething/app/bn/future/accumulation/handler/req"
	"tradething/app/bn/future/accumulation/service"

	"github.com/labstack/echo/v4"
	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

type upsertHandler struct {
	service service.IService
}

func NewUpsertHandler(service service.IService) *upsertHandler {
	return &upsertHandler{service: service}
}

func (h *upsertHandler) GetReqBody(c echo.Context) (*req.UpsertReq, error) {
	req := &req.UpsertReq{}
	if err := c.Bind(req); err != nil {
		return nil, err
	}
	return req, nil
}

func (h *upsertHandler) Handler(c echo.Context) error {
	reqBody, err := h.GetReqBody(c)
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.InvalidRequestErrorCode, err.Error(), nil)
		return response.SendResponse(http.StatusBadRequest, c)
	}

	err = h.service.Upsert(c.Request().Context(), reqBody.ToDomain())
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.FailCode, err.Error(), nil)
		return response.SendResponse(http.StatusInternalServerError, c)
	}

	response := appresponse.NewAppResponse(appresponse.SuccessCode, appresponse.SuccessMsg, nil)
	return response.SendResponse(http.StatusOK, c)
}
