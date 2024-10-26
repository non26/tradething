package bnfuture

import (
	"fmt"
	"strings"
	bnSvcfuture "tradething/app/bn/bn_future/bnservice_request_model"
	repomodel "tradething/app/bn/bn_future/repository_model"
)

type PlaceSignleOrderServiceRequest struct {
	positionSide  string
	side          string
	entryQuantity string // amountQ
	symbol        string
	leverageLevel int
	clientOrderId string
}

func (p *PlaceSignleOrderServiceRequest) GetPositionSide() string {
	return p.positionSide
}

func (p *PlaceSignleOrderServiceRequest) SetPositionSide(positionSide string) {
	p.positionSide = strings.ToUpper(positionSide)
}

func (p *PlaceSignleOrderServiceRequest) GetSide() string {
	return p.side
}

func (p *PlaceSignleOrderServiceRequest) SetSide(side string) {
	p.side = strings.ToUpper(side)
}

func (p *PlaceSignleOrderServiceRequest) GetEntryQuantity() string {
	return p.entryQuantity
}

func (p *PlaceSignleOrderServiceRequest) SetEntryQuantity(entryQuantity string) {
	p.entryQuantity = entryQuantity
}

func (p *PlaceSignleOrderServiceRequest) GetSymbol() string {
	return p.symbol
}

func (p *PlaceSignleOrderServiceRequest) SetSymbol(symbol string) {
	p.symbol = symbol
}

func (p *PlaceSignleOrderServiceRequest) GetLeverageLevel() int {
	return p.leverageLevel
}

func (p *PlaceSignleOrderServiceRequest) SetLeverageLevel(leverageLevel int) {
	p.leverageLevel = leverageLevel
}

func (p *PlaceSignleOrderServiceRequest) GetClientOrderId() string {
	return p.clientOrderId
}

func (p *PlaceSignleOrderServiceRequest) SetClientOrderId(clientOrderId string) {
	p.clientOrderId = clientOrderId
}

func (p *PlaceSignleOrderServiceRequest) ToBinanceServiceModel() *bnSvcfuture.PlaceSignleOrderBinanceServiceRequest {
	m := bnSvcfuture.PlaceSignleOrderBinanceServiceRequest{
		PositionSide:  p.positionSide,
		Side:          p.side,
		EntryQuantity: p.entryQuantity,
		Symbol:        p.symbol,
		ClientOrderId: p.clientOrderId,
	}
	return &m
}

func (p *PlaceSignleOrderServiceRequest) ToRepositoryModel() *repomodel.BinanceFutureOpeningPositionTable {
	m := repomodel.NewBinanceFutureOpeningPositionTable()
	m.Symbol = p.symbol
	m.PositionSide = p.positionSide
	m.AmountQ = p.entryQuantity
	m.Leverage = fmt.Sprintf("%v", p.leverageLevel)
	return m
}
