package bkcommon

import (
	"encoding/json"
	"io"
)

type bitKubResponse[T any] struct {
	code   string
	result T
}

func NewBitKubResponse[T any]() *bitKubResponse[T] {
	return &bitKubResponse[T]{}
}

func (b *bitKubResponse[T]) DecodeBkResponse(r io.Reader) error {
	type Decoded[K any] struct {
		Code   string `json:"error"`
		Result K      `json:"result"`
	}
	decoded := new(Decoded[T])
	err := json.NewDecoder(r).Decode(decoded)
	if err != nil {
		return err
	}
	b.code = decoded.Code
	b.result = decoded.Result
	return nil
}

func (b *bitKubResponse[T]) GetBkInternalCode() string {
	return b.code
}

func (b *bitKubResponse[T]) GetBkResult() T {
	return b.result
}
