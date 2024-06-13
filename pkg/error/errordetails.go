package errordetails

import (
	"fmt"
	"strings"
)

type ErrorDetails struct {
	err    error
	title  string
	fields []field
}

type field struct {
	key string
	val string
}

func NewErrorDetails(err error, title string) *ErrorDetails {
	return &ErrorDetails{
		err:   err,
		title: title,
	}
}

func (e *ErrorDetails) Error() string {
	var builder strings.Builder
	builder.WriteString(e.title)

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
