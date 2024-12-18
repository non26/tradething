package bnfuture

import svcfuture "tradething/app/bn/bn_future/service_model"

type CloseBySymbolsHandlerRequest struct {
	Data []CloseBySymbolsHandlerRequestData
}

type CloseBySymbolsHandlerRequestData struct {
	Symbol       string `json:"symbol"`
	PositionSide string `json:"position_side"`
	AmountQ      string `json:"amount_q"`
}

func (c *CloseBySymbolsHandlerRequest) ToServiceModel() *svcfuture.CloseBySymbolsServiceRequest {
	serviceModel := svcfuture.NewCloseBySymbolsServiceRequest()
	for _, data := range c.Data {
		rq_data := svcfuture.CloseBySymbolsServiceRequestData{}
		rq_data.SetSymbol(data.Symbol)
		rq_data.SetPositionSide(data.PositionSide)
		rq_data.SetAmountQ(data.AmountQ)
		serviceModel.SetData(rq_data)
	}
	return serviceModel
}
