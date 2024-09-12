package repository

type IDynamodb interface{}

type dynamodb struct{}

func NewDynamodb() IDynamodb {
	return &dynamodb{}
}
