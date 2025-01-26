package bnfuture

type PlacePosition struct {
	Symbol     string              `json:"symbol"`
	Quantity   string              `json:"quantity"`
	InValidate *InvalidatePosition `json:"invalidate_position"`
	Validate   *ValidatePosition   `json:"validate_position"`
}
