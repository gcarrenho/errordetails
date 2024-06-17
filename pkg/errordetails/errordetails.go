package errordetails

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

/*
{
  "level": "error",
  "error": {
    "file": "C:/Users/CarrenoG/go/src/errordetails/cmd/main.go",
    "line": 11,
    "message": "error test"
  },
  "time": "2024-06-15T17:26:14+02:00"
}
*/ /*
type ErrorDetails struct {
	ErrorForClient ErrorForClient `json:"error"`
	err            error
	fields         []field
	file           string `json:"file"`
	line           int
}

type ErrorForClient struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
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
}*/

// ErrorDetails struct to capture all relevant error information
type ErrorDetails struct {
	ErrorMessage string    `json:"error_message"`
	ErrorType    string    `json:"error_type"`
	File         string    `json:"file"`
	Line         int       `json:"line"`
	Function     string    `json:"function"`
	Timestamp    time.Time `json:"timestamp"`
	Context      []field
	StackTrace   string `json:"stack_trace"`
	err          error
}

// field struct for contextual information
type field struct {
	key string
	val string
}

// NewErrorDetails creates a new ErrorDetails instance and captures current context
func NewErrorDetails(err error) *ErrorDetails {
	pc, file, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	stackBuf := make([]byte, 1024)
	stackLen := runtime.Stack(stackBuf, false)
	stackTrace := strings.TrimSpace(string(stackBuf[:stackLen]))

	return &ErrorDetails{
		ErrorMessage: err.Error(),
		ErrorType:    fmt.Sprintf("%T", err),
		File:         file,
		Line:         line,
		Function:     fn.Name(),
		Timestamp:    time.Now(),
		StackTrace:   stackTrace,
		err:          err,
	}
}

// Error implements the error interface
func (e *ErrorDetails) Error() string {
	return e.err.Error()
}

// Unwrap returns the underlying error
func (e *ErrorDetails) Unwrap() error {
	return e.err
}

// Str adds a string key-value pair to the context
func (e *ErrorDetails) Str(key, val string) *ErrorDetails {
	e.Context = append(e.Context, field{key: key, val: val})
	return e
}

// Int adds an integer key-value pair to the context
func (e *ErrorDetails) Int(key string, val int) *ErrorDetails {
	e.Context = append(e.Context, field{key: key, val: fmt.Sprintf("%d", val)})
	return e
}

// Msg sets the error message
func (e *ErrorDetails) Msg(message string) *ErrorDetails {
	e.ErrorMessage = message
	return e
}

// ToClientError transforms ErrorDetails to a ClientError
func (e *ErrorDetails) ToClientError() *ClientError {
	return &ClientError{
		Message: e.ErrorMessage,
	}
}

// ClientError struct to return error message to the client
type ClientError struct {
	Message string `json:"message"`
}

// MarshalZerologObject implements the zerolog.LogObjectMarshaler interface
func (e *ErrorDetails) MarshalZerologObject(event *zerolog.Event) {
	event.Str("message", e.ErrorMessage)
	event.Str("file", e.File)
	event.Str("method", e.Function)
	event.Int("line", e.Line)
	event.Str("trace", e.StackTrace)
	event.Str("error_type", e.ErrorType)

	for _, f := range e.Context {
		event.Str(f.key, f.val)
	}

	if e.err != nil {
		event.Str("error", e.Error())
	}
}
