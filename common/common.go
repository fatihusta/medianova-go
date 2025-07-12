package common

import (
	"encoding/json"
	"net/http"
)

type JSONResult []byte

func (b JSONResult) String() string {
	return string(b)
}

type Result[T any] struct {
	Status  int
	Error   error
	Headers http.Header
	Body    T
}

func NewResult[T any]() *Result[T] {
	var body T
	return &Result[T]{
		Status: -1,
		Error:  nil,
		Body:   body,
	}
}

func (r *Result[T]) JSON() JSONResult {
	b, err := json.Marshal(r)
	if err != nil {
		return JSONResult([]byte("{}"))
	}
	return JSONResult(b)
}

type CtxKey string

func (k CtxKey) String() string {
	return string(k)
}
