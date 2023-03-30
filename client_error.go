package ernie_api

import "fmt"

type RequestError struct {
	StatusCode int
	Err        error
}

type ResponseError struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
	Err  error
}

func (e *ResponseError) Error() string {
	return fmt.Sprintf("status code %d , message %s", e.Code, e.Msg)
}

func (e *ResponseError) Unwrap() error {
	return e.Err
}

func (e *RequestError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return fmt.Sprintf("status code %d", e.StatusCode)
}

func (e *RequestError) Unwrap() error {
	return e.Err
}
