package domainservice

type LookupOpeningPosition struct {
	isFound      bool
	amountB      string
	clientId     string
	symbol       string
	positionSide string
	side         string
}

func NewLookupOpeningPosition() *LookupOpeningPosition {
	return &LookupOpeningPosition{}
}

func (l *LookupOpeningPosition) IsFound() bool {
	return l.isFound
}

func (l *LookupOpeningPosition) GetAmountB() string {
	return l.amountB
}

func (l *LookupOpeningPosition) GetClientId() string {
	return l.clientId
}

func (l *LookupOpeningPosition) GetSymbol() string {
	return l.symbol
}

func (l *LookupOpeningPosition) GetPositionSide() string {
	return l.positionSide
}

func (l *LookupOpeningPosition) GetSide() string {
	return l.side
}

func (l *LookupOpeningPosition) SetIsFound(isFound bool) {
	l.isFound = isFound
}

func (l *LookupOpeningPosition) SetAmountB(amountB string) {
	l.amountB = amountB
}

func (l *LookupOpeningPosition) SetClientId(clientId string) {
	l.clientId = clientId
}

func (l *LookupOpeningPosition) SetSymbol(symbol string) {
	l.symbol = symbol
}

func (l *LookupOpeningPosition) SetPositionSide(positionSide string) {
	l.positionSide = positionSide
}

func (l *LookupOpeningPosition) SetSide(side string) {
	l.side = side
}
