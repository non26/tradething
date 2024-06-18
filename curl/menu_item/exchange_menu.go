package menuitem

type IExchangeMenuItem interface {
	GetExchangeName() string
	GetExchangeId() int
}

type exchangeMenuItem struct {
	exchange_name string
	exchange_id   int
}

func NewExchangeMenu(name string, id int) *exchangeMenuItem {
	return &exchangeMenuItem{
		exchange_name: name,
		exchange_id:   id,
	}
}

func (e *exchangeMenuItem) GetExchangeName() string {
	return e.exchange_name
}

func (e *exchangeMenuItem) GetExchangeId() int {
	return e.exchange_id
}
