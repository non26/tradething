package bnfuture

type GetAllBinanceFee struct {
	Level        string `json:"level"`
	Maker        string `json:"maker"`
	Taker        string `json:"taker"`
	MakerBnbDisc string `json:"maker_bnb_disc"`
	TakerBnbDisc string `json:"taker_bnb_disc"`
}

type GetAllBinanceFeeHandlerResponse struct {
	Fees []GetAllBinanceFee `json:"fees"`
}
