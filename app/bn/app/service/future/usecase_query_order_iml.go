package service

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	bnmodel "tradetoolv2/app/bn/app/model/bnservicemodel/future"
	model "tradetoolv2/app/bn/app/model/handlermodel/future"
	"tradetoolv2/common"
)

func (bfs *binanceFutureService) QueryOrder(
	ctx context.Context,
	request *model.QueryOrderBinanceHandlerRequest,
) (*model.QueryOrderBinanceHandlerResponse, error) {
	request.Symbol = strings.ToUpper(request.Symbol)

	bnResponse, err := bfs.binanceService.QueryOrder(
		ctx,
		request,
	)
	if err != nil {
		return nil, err
	}

	if bnResponse.StatusCode != http.StatusOK {
		bnResponseError := new(bnmodel.ResponseBinanceFutureError)
		json.NewDecoder(bnResponse.Body).Decode(bnResponseError)
		msg := common.FormatMessageOtherThanHttpStatus200(
			bfs.binanceFutureServiceName,
			bnResponse.StatusCode,
			bnResponseError.Code,
			bnResponseError.Message,
		)
		return nil, errors.New(msg)
	}

	bnResponseError := new(bnmodel.QueryOrderBinanceServiceResponse)
	json.NewDecoder(bnResponse.Body).Decode(bnResponseError)

	return bnResponseError.ToHandlerResponse(), nil
}
