package api

type IApi interface {
	GetUserInputField() []string
	SetUserInputField()
	SetUserInputValue(user_input []string) error
	ExecuteCurl() error
}
