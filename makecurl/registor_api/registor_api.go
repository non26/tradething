package registorapi

import apiinfo "tradething/makecurl/registor_api/api_info"

type IRegistorApi interface {
	RegistorNewApi(
		api_info apiinfo.IApiInfo,
	)
}

type registorApi struct {
	max_api int
	apis    map[int]apiinfo.IApiInfo
}

func NewRegistorApi() IRegistorApi {
	return &registorApi{
		max_api: 0,
		apis:    make(map[int]apiinfo.IApiInfo),
	}
}

func (r *registorApi) RegistorNewApi(
	api_info apiinfo.IApiInfo,
) {
	r.apis[r.max_api] = api_info
	r.max_api = r.max_api + 1
}
