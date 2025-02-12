package gjmt_models

// Response обертка для ответа
// @Description В этот формат завернуты все ответы от сервиса
type Response[T any] struct {
	Value     T      `json:"value,omitempty"`
	Error     string `json:"error,omitempty"`
	IsSuccess bool   `json:"isSuccess"`
}

// Deprecated: Use ResponseWithValue instead. This function will be removed in a future release.
// NewSuccessResponse обертка для ответа
func NewSuccessResponse[T any](value T) *Response[T] {
	return &Response[T]{
		Value:     value,
		IsSuccess: true,
	}
}

// Deprecated: Use ResponseWithError instead. This function will be removed in a future release.
func NewErrorResponse(err string) *Response[any] {
	return &Response[any]{
		Error:     err,
		IsSuccess: false,
	}
}

func ResponseWithError(err string) *Response[any] {
	return &Response[any]{
		Error:     err,
		IsSuccess: false,
	}
}

// NewSuccessResponse обертка для ответа
func ResponseWithValue[T any](value T) *Response[T] {
	return &Response[T]{
		Value:     value,
		IsSuccess: true,
	}
}

// ResponseErrorType обертка для ответа с ошибкой
//
//	@example {
//	  "value": nil,
//	  "error": "Text error message",
//	  "isSuccess": false
//	}
type ResponseErrorType Response[string]
