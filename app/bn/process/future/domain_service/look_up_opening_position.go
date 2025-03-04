package domainservice

type lookupOpeningPosition struct {
	isFound  bool
	amountB  string
	clientId string
}

func (l *lookupOpeningPosition) IsFound() bool {
	return l.isFound
}

func (l *lookupOpeningPosition) GetAmountB() string {
	return l.amountB
}

func (l *lookupOpeningPosition) GetClientId() string {
	return l.clientId
}

func (l *lookupOpeningPosition) SetIsFound(isFound bool) {
	l.isFound = isFound
}

func (l *lookupOpeningPosition) SetAmountB(amountB string) {
	l.amountB = amountB
}

func (l *lookupOpeningPosition) SetClientId(clientId string) {
	l.clientId = clientId
}

func NewLookupOpeningPosition() *lookupOpeningPosition {
	return &lookupOpeningPosition{}
}
