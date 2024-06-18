package menuitem

type apiMenuStructure struct {
	api_path     string
	request_body []string
	api_id       int
	exchange_id  int
}

func NewApiMenu(path string, req_body []string, api_id int, exchange_id int) *apiMenuStructure {
	return &apiMenuStructure{
		api_path:     path,
		request_body: req_body,
	}
}

func (a *apiMenuStructure) GetApiPath() string {
	return a.api_path
}

func (a *apiMenuStructure) GetRequestBody() []string {
	return a.request_body
}

func (a *apiMenuStructure) GetApiId() int {
	return a.api_id
}
