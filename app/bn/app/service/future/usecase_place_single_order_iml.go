package service

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	bnmodel "tradething/app/bn/app/model/bnservicemodel/future"
	model "tradething/app/bn/app/model/handlermodel/future"
	"tradething/common"
)

func (bfs *binanceFutureService) PlaceSingleOrder(
	ctx context.Context,
	request *model.PlaceSignleOrderHandlerRequest,
) (*model.PlaceSignleOrderHandlerResponse, error) {

	request.Symbol = strings.ToUpper(request.Symbol)
	request.Side = strings.ToUpper(request.Side)
	request.PositionSide = strings.ToUpper(request.PositionSide)

	bnResponse, err := bfs.binanceService.PlaceSingleOrder(
		ctx,
		request,
	)
	if err != nil {
		return nil, err
	}
	defer bnResponse.Body.Close()

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

	res := &model.PlaceSignleOrderHandlerResponse{
		Symbol:   request.Symbol,
		Quantity: request.EntryQuantity,
	}
	return res, nil
}
