package models

type Response[T any] struct {
	Value     T      `json:"value,omitempty"`
	Error     string `json:"error,omitempty"`
	IsSuccess bool   `json:"isSuccess"`
}

func NewSuccessResponse[T any](value T) *Response[T] {
	return &Response[T]{
		Value:     value,
		IsSuccess: true,
	}
}

func NewErrorResponse(err string) *Response[any] {
	return &Response[any]{
		Error:     err,
		IsSuccess: false,
	}
}
