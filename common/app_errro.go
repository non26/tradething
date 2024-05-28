package common

type AppError struct {
	ExchangeErrorCode    string
	ExchangeErrorMessage string
	ServiceName          string
	ServiceErrorMessage  string
}

// requet error pattern
// response error pattern
// other error pattern
