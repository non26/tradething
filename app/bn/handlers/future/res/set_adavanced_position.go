package res

type SetAdvancedPositionResponses struct {
	Responses []SetAdvancedPositionResponse `json:"responses"`
}

func NewSetAdvancedPositionResponses() *SetAdvancedPositionResponses {
	return &SetAdvancedPositionResponses{
		Responses: []SetAdvancedPositionResponse{},
	}
}

func (s *SetAdvancedPositionResponses) Add(response *SetAdvancedPositionResponse) {
	s.Responses = append(s.Responses, *response)
}

type SetAdvancedPositionResponse struct {
	ClientId string `json:"client_id"`
	Status   string `json:"status"`
}

func NewSetAdvancedPositionResponse() *SetAdvancedPositionResponse {
	return &SetAdvancedPositionResponse{
		ClientId: "",
		Status:   "",
	}
}

func (s *SetAdvancedPositionResponse) Fail(clientId string) {
	s.ClientId = clientId
	s.Status = "fail"
}

func (s *SetAdvancedPositionResponse) Success(clientId string) {
	s.ClientId = clientId
	s.Status = "success"
}
