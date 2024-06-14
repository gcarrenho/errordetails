package errordetails

import (
	"fmt"
	"strings"
)

type ErrorDetails struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	err     error
	fields  []field
}

type field struct {
	key string
	val string
}

func NewErrorDetails(err error) *ErrorDetails {
	return &ErrorDetails{
		err: err,
	}
}

func (e *ErrorDetails) Error() string {
	var builder strings.Builder
	builder.WriteString(e.Message)

	for _, f := range e.fields {
		builder.WriteString(fmt.Sprintf(" | %s: %s", f.key, f.val))
	}

	if e.err != nil {
		builder.WriteString(fmt.Sprintf(" --> %s", e.err.Error()))
	}

	return builder.String()
}

func (e *ErrorDetails) Unwrap() error {
	return e.err
}

func (e *ErrorDetails) Str(key, val string) *ErrorDetails {
	e.fields = append(e.fields, field{key: key, val: val})
	return e
}

func (e *ErrorDetails) Int(key string, val int) *ErrorDetails {
	e.fields = append(e.fields, field{key: key, val: fmt.Sprintf("%d", val)})
	return e
}

func (e *ErrorDetails) Msg(message string) *ErrorDetails {
	e.Message = message
	return e
}
