package bnfuture

type ClientIds struct {
	clientIds []string
}

func (m *ClientIds) GetCleintIds() []string {
	return m.clientIds
}

func (m *ClientIds) SetClientIds(clientIds []string) {
	m.clientIds = clientIds
}
