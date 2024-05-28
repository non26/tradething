package kccommon

import (
	"encoding/json"
	"io"
)

type kucoinResponse[T any] struct {
	code string
	data T
}

func NewKucoinResponse[T any]() *kucoinResponse[T] {
	return &kucoinResponse[T]{}
}

func (b *kucoinResponse[T]) DecodeKcResponse(r io.Reader) error {
	type Decoded[K any] struct {
		Code string `json:"code"`
		Data K      `json:"data"`
	}
	decoded := new(Decoded[T])
	err := json.NewDecoder(r).Decode(decoded)
	if err != nil {
		return err
	}
	b.code = decoded.Code
	b.data = decoded.Data
	return nil
}

func (b *kucoinResponse[T]) GetKcInternalCode() string {
	return b.code
}

func (b *kucoinResponse[T]) GetKcResult() T {
	return b.data
}
