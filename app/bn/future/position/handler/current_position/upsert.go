package handler

import (
	"net/http"
	req "tradething/app/bn/future/position/handler/req/current_position"
	"tradething/app/bn/future/position/service"

	"github.com/labstack/echo/v4"
	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

type upsertHandler struct {
	service service.ICurrentPositionService
}

func NewUpsertHandler(service service.ICurrentPositionService) *upsertHandler {
	return &upsertHandler{service: service}
}

func (h *upsertHandler) GetReqBody(c echo.Context) (*req.UpsertCurrentPositionReq, error) {
	req := &req.UpsertCurrentPositionReq{}
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
