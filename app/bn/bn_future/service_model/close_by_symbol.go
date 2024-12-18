package bnfuture

type CloseBySymbolsServiceRequest struct {
	data []CloseBySymbolsServiceRequestData
}

func NewCloseBySymbolsServiceRequest() *CloseBySymbolsServiceRequest {
	return &CloseBySymbolsServiceRequest{
		data: []CloseBySymbolsServiceRequestData{},
	}
}

func (c *CloseBySymbolsServiceRequest) SetData(data CloseBySymbolsServiceRequestData) {
	c.data = append(c.data, data)
}

func (c *CloseBySymbolsServiceRequest) GetData() []CloseBySymbolsServiceRequestData {
	return c.data
}

type CloseBySymbolsServiceRequestData struct {
	symbol       string
	positionSide string
	amountQ      string
}

func (c *CloseBySymbolsServiceRequestData) SetSymbol(symbol string) {
	c.symbol = symbol
}

func (c *CloseBySymbolsServiceRequestData) GetSymbol() string {
	return c.symbol
}

func (c *CloseBySymbolsServiceRequestData) SetPositionSide(positionSide string) {
	c.positionSide = positionSide
}

func (c *CloseBySymbolsServiceRequestData) GetPositionSide() string {
	return c.positionSide
}

func (c *CloseBySymbolsServiceRequestData) SetAmountQ(amountQ string) {
	c.amountQ = amountQ
}

func (c *CloseBySymbolsServiceRequestData) GetAmountQ() string {
	return c.amountQ
}
