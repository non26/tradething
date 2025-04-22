package domainservice

type LookUpAdvancedPosition struct {
	isFound      bool
	symbol       string
	positionSide string
	side         string
	amountB      string
	clientId     string
}

func NewLookUpAdvancedPosition() *LookUpAdvancedPosition {
	return &LookUpAdvancedPosition{}
}

func (l *LookUpAdvancedPosition) IsFound() bool {
	return l.isFound
}

func (l *LookUpAdvancedPosition) GetSymbol() string {
	return l.symbol
}

func (l *LookUpAdvancedPosition) GetPositionSide() string {
	return l.positionSide
}

func (l *LookUpAdvancedPosition) GetSide() string {
	return l.side
}

func (l *LookUpAdvancedPosition) GetAmountB() string {
	return l.amountB
}

func (l *LookUpAdvancedPosition) GetClientId() string {
	return l.clientId
}

func (l *LookUpAdvancedPosition) SetIsFound(isFound bool) {
	l.isFound = isFound
}

func (l *LookUpAdvancedPosition) SetSymbol(symbol string) {
	l.symbol = symbol
}

func (l *LookUpAdvancedPosition) SetPositionSide(positionSide string) {
	l.positionSide = positionSide
}

func (l *LookUpAdvancedPosition) SetSide(side string) {
	l.side = side
}

func (l *LookUpAdvancedPosition) SetAmountB(amountB string) {
	l.amountB = amountB
}

func (l *LookUpAdvancedPosition) SetClientId(clientId string) {
	l.clientId = clientId
}
