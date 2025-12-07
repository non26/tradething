package handler

import (
	"net/http"
	req "tradething/app/bn/future/position/handler/req/current_position"
	res "tradething/app/bn/future/position/handler/res/current_position"
	"tradething/app/bn/future/position/service"

	"github.com/labstack/echo/v4"
	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

type getHandler struct {
	service service.ICurrentPositionService
}

func NewGetHandler(service service.ICurrentPositionService) *getHandler {
	return &getHandler{service: service}
}

func (h *getHandler) GetReqBody(c echo.Context) (*req.GetCurrentPositionReq, error) {
	req := &req.GetCurrentPositionReq{}
	if err := c.Bind(req); err != nil {
		return nil, err
	}
	return req, nil
}

func (h *getHandler) Handler(c echo.Context) error {
	reqBody, err := h.GetReqBody(c)
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.InvalidRequestErrorCode, err.Error(), nil)
		return response.SendResponse(http.StatusBadRequest, c)
	}

	positions, err := h.service.Get(c.Request().Context(), reqBody.Symbol, reqBody.AccountId, reqBody.PositionSide)
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.FailCode, err.Error(), nil)
		return response.SendResponse(http.StatusInternalServerError, c)
	}

	responseData := &res.GetCurrentPositionRes{}
	responseData = responseData.FromDomain(positions)
	var response *appresponse.AppResponse
	if responseData == nil {
		response = appresponse.NewAppResponse(appresponse.SuccessCode, appresponse.SuccessMsg, nil)
	} else {
		response = appresponse.NewAppResponse(appresponse.SuccessCode, appresponse.SuccessMsg, responseData)
	}
	return response.SendResponse(http.StatusOK, c)
}
