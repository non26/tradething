package requestbodyinfo

type IRequestBodyInfo interface {
	GetRequestBody() []string
}

type requestBodyInfo struct {
	field      string
	field_type string
}

func NewRequestBodyInfo(
	field string,
	field_type string,
) IRequestBodyInfo {
	return &requestBodyInfo{
		field,
		field_type,
	}
}

func (r *requestBodyInfo) GetRequestBody() []string {
	return []string{r.field, r.field_type}
}
