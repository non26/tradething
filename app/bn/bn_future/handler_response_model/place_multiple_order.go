package bnfuture

type PlaceMultipleOrderHandlerResponse struct {
	Result []PlaceSignleOrderHandlerResponse `json:"result"`
}
