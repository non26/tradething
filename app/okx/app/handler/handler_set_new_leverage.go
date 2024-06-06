package handler

import (
	"net/http"
	handlermodel "tradething/app/okx/app/model/handlermodel"
	"tradething/app/okx/app/service"
	"tradething/common"

	"github.com/labstack/echo/v4"
)

type ISetLeverage interface {
	Handler(c echo.Context) error
}

type setLeverage struct {
	okxservice service.IOkxAppService
}

func NewSetLeverage(
	okxservice service.IOkxAppService,
) ISetLeverage {
	return &setLeverage{
		okxservice,
	}
}
func (s *setLeverage) Handler(c echo.Context) error {
	body := new(handlermodel.SetNewLeverageHandlerRequest)
	err := c.Bind(body)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			&common.CommonResponse{
				Code:    common.FailCode,
				Message: err.Error(),
			},
		)
	}

	data, err := s.okxservice.SetNewLeverage(body)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			&common.CommonResponse{
				Code:    common.FailCode,
				Message: err.Error(),
				Data:    nil,
			},
		)
	}

	return c.JSON(
		http.StatusOK,
		common.CommonResponse{
			Code:    common.SuccessCode,
			Message: "Success",
			Data:    data,
		},
	)
}
