package domainservice

type LookUpOpeningPosition struct {
	isFound  bool
	clientId string
	symbol   string
	quantity string
}

func NewLookUpOpeningPosition() *LookUpOpeningPosition {
	return &LookUpOpeningPosition{}
}

func (l *LookUpOpeningPosition) SetIsFound(isFound bool) {
	l.isFound = isFound
}

func (l *LookUpOpeningPosition) SetClientId(clientId string) {
	l.clientId = clientId
}

func (l *LookUpOpeningPosition) SetSymbol(symbol string) {
	l.symbol = symbol
}

func (l *LookUpOpeningPosition) SetQuantity(quantity string) {
	l.quantity = quantity
}

func (l *LookUpOpeningPosition) GetClientId() string {
	return l.clientId
}

func (l *LookUpOpeningPosition) GetSymbol() string {
	return l.symbol
}

func (l *LookUpOpeningPosition) GetQuantity() string {
	return l.quantity
}

func (l *LookUpOpeningPosition) IsFound() bool {
	return l.isFound
}
