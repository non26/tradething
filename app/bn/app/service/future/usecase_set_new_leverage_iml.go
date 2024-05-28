package service

import (
	"context"
	"encoding/json"
	"errors"
	bnmodel "tradetoolv2/app/bn/app/model/bnservicemodel/future"
	model "tradetoolv2/app/bn/app/model/handlermodel/future"
	"tradetoolv2/common"
)

func (bfs *binanceFutureService) SetNewLeverage(
	ctx context.Context,
	request *model.SetLeverageHandlerRequest) error {
	bnResponse, err := bfs.binanceService.SetNewLeverage(
		ctx,
		request,
	)
	if err != nil {
		return err
	}
	defer bnResponse.Body.Close()

	if bnResponse.StatusCode != 200 {
		bnResponseError := new(bnmodel.ResponseBinanceFutureError)
		json.NewDecoder(bnResponse.Body).Decode(bnResponseError)
		msg := common.FormatMessageOtherThanHttpStatus200(
			bfs.binanceFutureServiceName,
			bnResponse.StatusCode,
			bnResponseError.Code,
			bnResponseError.Message,
		)
		return errors.New(msg)

	}
	return nil
}
