package err

import (
	"fmt"
	"strings"
)

// SprintErrorOption is a type for options that can be passed to SprintError.
type SprintErrorOption func(*string)

// WithExtra includes extra information about the error in formatted error code.
func WithExtra(extra string) SprintErrorOption {
	return func(msg *string) {
		*msg = fmt.Sprintf("%s\n\t%s", *msg, extra)
	}
}

// WithOriginalErr includes a line about origin error information.
func WithOriginalErr(origErr error) SprintErrorOption {
	return func(msg *string) {
		*msg = fmt.Sprintf("%s\ncaused by: %s", *msg, origErr.Error())
	}
}

// SprintError returns a string of the formatted error code.
func SprintError(code, message string, options ...SprintErrorOption) string {
	msg := fmt.Sprintf("%s: %s", code, message)
	for _, option := range options {
		option(&msg)
	}
	return msg
}

// An error list that satisfies builtin error interface
type errorList []error

// Error returns the string representation of the error.
func (e errorList) Error() string {
	var msg strings.Builder
	for i, err := range e {
		if i > 0 {
			msg.WriteString("\n")
		}
		msg.WriteString(err.Error())
	}
	return msg.String()
}
