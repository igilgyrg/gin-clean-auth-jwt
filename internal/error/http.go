package error

import "encoding/json"

type Error interface {
	Status() int
	Error() string
	Causes() interface{}
	Marshal() []byte
}

type HttpError struct {
	HttpStatus  int         `json:"status"`
	HttpError   string      `json:"error"`
	HttpMessage string      `json:"message"`
	HttpCauses  interface{} `json:"-"`
}

func (e *HttpError) Error() string {
	return e.HttpError
}

func (e *HttpError) Status() int {
	return e.HttpStatus
}

func (e *HttpError) Causes() interface{} {
	return e.HttpCauses
}

func (e *HttpError) Message() string {
	return e.HttpMessage
}

func (h *HttpError) Marshal() []byte {
	marshal, err := json.Marshal(h)
	if err != nil {
		return nil
	}

	return marshal
}
