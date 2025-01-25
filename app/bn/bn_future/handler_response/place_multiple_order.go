package bnfuture

type PlaceMultiplePosition struct {
	Result PlaceMultiplePositionData
}

type PlaceMultiplePositionData struct {
	Success []string
	Failed  []PlaceMultiplePositionFailed
}

type PlaceMultiplePositionFailed struct {
	Symbol string
	Error  string
}
