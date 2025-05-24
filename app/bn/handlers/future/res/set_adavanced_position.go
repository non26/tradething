package res

type SetAdvancedPositionResponses struct {
	Responses []SetAdvancedPositionResponse
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
	Message  string `json:"message,omitempty"`
}

func NewSetAdvancedPositionResponse() *SetAdvancedPositionResponse {
	return &SetAdvancedPositionResponse{
		ClientId: "",
		Status:   "",
	}
}

func (s *SetAdvancedPositionResponse) Fail(clientId string, message string) {
	s.ClientId = clientId
	s.Status = "fail"
	s.Message = message
}

func (s *SetAdvancedPositionResponse) Success(clientId string) {
	s.ClientId = clientId
	s.Status = "success"
}
