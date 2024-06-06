package okxservice

import (
	"net/http"
	model "tradething/app/okx/app/model/okxservicemodel"
	"tradething/config"
)

type IOKXExternalService interface {
	// GetLeverage() (*http.Response, error)
	SetLeverage(
		body *model.SetNewLeverageOKXServiceRequest,
	) (*model.SetNewLeverageOKXserviceResponse, error)
	PlaceASinglePosition(
		body *model.PlaceASinglePositionOKXServiceRequest,
	) (*model.PlaceASinglePositionOKXserviceResponse, error)
	PlaceMultiplePosition(
		body *model.PlaceMultiplePositionOKXServiceRequest,
	) (*http.Response, error)
}

type okxExternalService struct {
	okxFutureUrl *config.OkxFutureUrl
	secret       *config.Secrets
	env          string
}

func NewOKXExternalService(
	okxFutureUrl *config.OkxFutureUrl,
	secret *config.Secrets,
	env string,
) IOKXExternalService {
	return &okxExternalService{
		okxFutureUrl,
		secret,
		env,
	}
}
