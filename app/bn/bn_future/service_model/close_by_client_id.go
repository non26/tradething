package bnfuture

type CloseByClientIdServiceRequest struct {
	clientIds []string
}

func (m *CloseByClientIdServiceRequest) GetClientIds() []string {
	return m.clientIds
}

func (m *CloseByClientIdServiceRequest) SetClientIds(clientIds []string) {
	m.clientIds = clientIds
}
