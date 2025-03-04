package domainservice

type lookUpSymbol struct {
	symbol        string
	countingLong  int64
	countingShort int64
}

func NewLookUpSymbol() *lookUpSymbol {
	return &lookUpSymbol{}
}

func (l *lookUpSymbol) GetSymbol() string {
	return l.symbol
}

func (l *lookUpSymbol) GetCountingLong() int64 {
	return l.countingLong
}

func (l *lookUpSymbol) GetCountingShort() int64 {
	return l.countingShort
}

func (l *lookUpSymbol) SetSymbol(symbol string) {
	l.symbol = symbol
}

func (l *lookUpSymbol) SetCountingLong(countingLong int64) {
	l.countingLong = countingLong
}

func (l *lookUpSymbol) SetCountingShort(countingShort int64) {
	l.countingShort = countingShort
}
