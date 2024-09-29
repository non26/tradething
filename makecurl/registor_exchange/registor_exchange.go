package registorexchange

type IRegistorExchange interface {
	RegistorNewExchanges(exchange string)
}

type registorExchange struct {
	exchanges []string
}

func NewRegistorExchange() IRegistorExchange {
	return &registorExchange{
		exchanges: make([]string, 0, 10),
	}
}

func (rg *registorExchange) RegistorNewExchanges(exchange string) {
	rg.exchanges = append(rg.exchanges, exchange)
}
