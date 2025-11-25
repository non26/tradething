package handler

import (
	"net/http"
	"tradething/app/bn/future/BFF/account_management/handler/req"
	"tradething/app/bn/future/BFF/account_management/service"

	"github.com/labstack/echo/v4"
	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

type updateHandler struct {
	service service.IAccountManagementService
}

func NewUpdateHandler(service service.IAccountManagementService) *updateHandler {
	return &updateHandler{service: service}
}

func (h *updateHandler) GetReqBody(c echo.Context) (*req.UpdateReq, error) {
	req := &req.UpdateReq{}
	if err := c.Bind(req); err != nil {
		return nil, err
	}
	return req, nil
}

func (h *updateHandler) Handler(c echo.Context) error {
	reqBody, err := h.GetReqBody(c)
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.InvalidRequestErrorCode, err.Error(), nil)
		return response.SendResponse(http.StatusBadRequest, c)
	}

	err = h.service.UpdateSubAccount(c.Request().Context(), reqBody.ToDomain())
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.FailCode, err.Error(), nil)
		return response.SendResponse(http.StatusInternalServerError, c)
	}

	response := appresponse.NewAppResponse(appresponse.SuccessCode, appresponse.SuccessMsg, nil)
	return response.SendResponse(http.StatusOK, c)
}
