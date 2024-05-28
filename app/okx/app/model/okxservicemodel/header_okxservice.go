package model

type okxExternalServiceHeaders struct {
	AccessKeyHeader  string
	AccessSignHeader string
	AccessTimestamp  string
	AccessPassPhase  string
	SimulateTrading  string
}

func NewOKXExternalServiceHeader() (ok *okxExternalServiceHeaders) {
	ok.AccessKeyHeader = "OK-ACCESS-KEY"
	ok.AccessSignHeader = "OK-ACCESS-SIGN"
	ok.AccessTimestamp = "OK-ACCESS-TIMESTAMP"
	ok.AccessPassPhase = "OK-ACCESS-PASSPHRASE"
	ok.SimulateTrading = "x-simulated-trading"
	return ok
}
