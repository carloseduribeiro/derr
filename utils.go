package derr

import (
	"fmt"
	"strings"
)

// sprintErrorOption is a type for options that can be passed to sprintError.
type sprintErrorOption func(*string)

// withExtra includes extra information about the error in formatted error code.
func withExtra(extra string) sprintErrorOption {
	return func(msg *string) {
		*msg = fmt.Sprintf("%s\n\t%s", *msg, extra)
	}
}

// withOriginalErr includes a line about origin error information.
func withOriginalErr(origErr error) sprintErrorOption {
	return func(msg *string) {
		*msg = fmt.Sprintf("%s\ncaused by: %s", *msg, origErr.Error())
	}
}

// sprintError returns a string of the formatted error code.
func sprintError(code, message string, options ...sprintErrorOption) string {
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
