package errordetails

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/rs/zerolog"
)

type ErrorDetails struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	err     error
	fields  []field
	file    string
	line    int
}

type field struct {
	key string
	val string
}

func NewErrorDetails(err error) *ErrorDetails {
	// Captura la información del archivo y línea
	_, file, line, _ := runtime.Caller(1)
	return &ErrorDetails{
		err:  err,
		file: file,
		line: line,
	}
}

func (e *ErrorDetails) Error() string {
	var builder strings.Builder
	//builder.WriteString(e.Message)
	builder.WriteString(fmt.Sprintf("%s:%d: %s", e.file, e.line, e.Message))

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

// Implement the MarshalZerologObject method to support zerolog
func (e *ErrorDetails) MarshalZerologObject(event *zerolog.Event) {
	event.Str("message", e.Message)
	event.Str("file", e.file)
	event.Int("line", e.line)

	for _, f := range e.fields {
		event.Str(f.key, f.val)
	}

	if e.err != nil {
		event.Str("error", e.err.Error())
	}
}
