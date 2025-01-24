package bnfuture

type AdvanceConfig struct {
}

type InvalidatePosition struct {
	ClientIds []string `json:"client_ids"`
}

type ActivePosition struct {
	ClientIds []string `json:"client_ids"`
}
