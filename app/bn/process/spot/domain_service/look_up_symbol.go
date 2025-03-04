package domainservice

type LookUpSymbol struct {
	symbol   string
	counting int64
}

func NewLookUpSymbol() *LookUpSymbol {
	return &LookUpSymbol{}
}

func (l *LookUpSymbol) SetSymbol(symbol string) {
	l.symbol = symbol
}

func (l *LookUpSymbol) SetCounting(counting int64) {
	l.counting = counting
}

func (l *LookUpSymbol) GetSymbol() string {
	return l.symbol
}

func (l *LookUpSymbol) GetCounting() int64 {
	return l.counting
}
