package domainservice

type LookUpHistory struct {
	isFound  bool
	clientId string
}

func NewLookUpHistory() *LookUpHistory {
	return &LookUpHistory{}
}

func (l *LookUpHistory) IsFound() bool {
	return l.isFound
}

func (l *LookUpHistory) GetClientId() string {
	return l.clientId
}

func (l *LookUpHistory) SetIsFound(isFound bool) {
	l.isFound = isFound
}

func (l *LookUpHistory) SetClientId(clientId string) {
	l.clientId = clientId
}
