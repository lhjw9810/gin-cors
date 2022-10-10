package httpResponse

var (
	successCode = "200"
	errCode     = "-1"
)

type EntityMessage struct {
	Code   string      `json:"code"`
	ErrMsg string      `json:"error_msg,omitempty"`
	Body   interface{} `json:"body,omitempty"`
}

func Success(body interface{}) *EntityMessage {
	return &EntityMessage{
		Code: successCode,
		Body: body,
	}
}
func Empty() *EntityMessage {
	return &EntityMessage{
		Code: successCode,
	}
}
func Error(code string, err error) *EntityMessage {
	if len(code) > 0 {
		return &EntityMessage{
			Code:   code,
			ErrMsg: err.Error(),
		}
	} else {
		return &EntityMessage{
			Code:   errCode,
			ErrMsg: err.Error(),
		}
	}
}
