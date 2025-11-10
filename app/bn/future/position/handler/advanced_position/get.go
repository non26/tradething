package handler

import (
	"net/http"
	req "tradething/app/bn/future/position/handler/req/advanced_position"
	res "tradething/app/bn/future/position/handler/res/advanced_position"
	"tradething/app/bn/future/position/service"

	"github.com/labstack/echo/v4"
	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

type getHandler struct {
	service service.IAdvancedPositionService
}

func NewGetHandler(service service.IAdvancedPositionService) *getHandler {
	return &getHandler{service: service}
}

func (h *getHandler) GetReqBody(c echo.Context) (*req.GetAdvancedPositionReq, error) {
	req := &req.GetAdvancedPositionReq{}
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

	positions, err := h.service.Get(c.Request().Context(), reqBody.ClientId)
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.FailCode, err.Error(), nil)
		return response.SendResponse(http.StatusInternalServerError, c)
	}

	responseData := &res.GetAdvancedPositionRes{}
	responseData = responseData.FromDomain(positions)
	response := appresponse.NewAppResponse(appresponse.SuccessCode, appresponse.SuccessMsg, responseData)
	return response.SendResponse(http.StatusOK, c)
}
