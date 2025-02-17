package req

type MultipleOrders struct {
	Orders []Trade `json:"orders"`
}
