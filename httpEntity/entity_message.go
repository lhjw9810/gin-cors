package httpEntity

var (
	successCode = "200"
	errCode     = "-1"
)

type EntityMessage[T any] struct {
	Code   string `json:"code"`
	ErrMsg string `json:"error_msg,omitempty"`
	Body   T      `json:"body,omitempty"`
}

func Success[T any](body T) *EntityMessage[T] {
	return &EntityMessage[T]{
		Code: successCode,
		Body: body,
	}
}
func Empty[T any]() *EntityMessage[T] {
	return &EntityMessage[T]{
		Code: successCode,
	}
}
func Error[T any](code string, err error) *EntityMessage[T] {
	if len(code) > 0 {
		return &EntityMessage[T]{
			Code:   code,
			ErrMsg: err.Error(),
		}
	} else {
		return &EntityMessage[T]{
			Code:   errCode,
			ErrMsg: err.Error(),
		}
	}
}
