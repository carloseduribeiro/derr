package err

import "errors"

// NewErrorOption is a type for options that can be passed to NewError and NewBatchError functions.
type NewErrorOption func(*baseError)

// WithError includes an error in the new error object as original error.
func WithError(err error) NewErrorOption {
	return func(b *baseError) {
		if err != nil {
			if b.errs == nil {
				b.errs = make([]error, 0, 1)
			}
			b.errs = append(b.errs, err)
		}
	}
}

// WithErrorS includes errors in the new error object as original errors.
func WithErrorS(errs []error) NewErrorOption {
	return func(b *baseError) {
		if errs != nil && len(errs) > 0 {
			if b.errs == nil {
				b.errs = make([]error, 0, len(errs))
			}
			b.errs = append(b.errs, errs...)
		}
	}
}

// A baseError wraps the code and message which defines an error. It also can be used to wrap an original error object.
// Should be used as the root for errors satisfying the Error interface.
// Also for any error which does not fit into a specific error wrapper type.
type baseError struct {
	// code is a short no whitespace phrase depicting the classification of the error that is being created.
	code string
	// message is the free flow string containing detailed information about the error.
	message string
	// errs is the error objects which will be nested under the new errors to be returned. Allows building chained errors.
	errs []error
}

// newBaseError returns an error object for the code, message, and errors.
func newBaseError(code, message string, options ...NewErrorOption) *baseError {
	b := &baseError{
		code:    code,
		message: message,
	}
	for _, option := range options {
		if option != nil {
			option(b)
		}
	}
	return b
}

// Error returns the string representation of the error.
func (b baseError) Error() string {
	if len(b.errs) > 0 {
		return SprintError(b.code, b.message, WithOriginalErr(errorList(b.errs)))
	}
	return SprintError(b.code, b.message)
}

// String returns the string representation of the error. It's an alias for Error to satisfy the stringer interface.
func (b baseError) String() string {
	return b.Error()
}

// Code returns the short phrase that describes the error classification.
func (b baseError) Code() string {
	return b.code
}

// Message returns the error details message.
func (b baseError) Message() string {
	return b.message
}

// OrigErr returns the original error if one was set.
// Nil is returned if no error was set. If the full list is needed, use BatchError.
func (b baseError) OrigErr() error {
	switch len(b.errs) {
	case 0:
		return nil
	case 1:
		return b.errs[0]
	default:
		var err Error
		if errors.As(b.errs[0], &err) {
			return NewBatchError(err.Code(), err.Message(), WithErrorS(b.errs[1:]))
		}
		return NewBatchError("BatchError", "multiple errors occurred", WithErrorS(b.errs))
	}
}

// OrigErrs returns the original errors if at least one was set. An empty slice is returned if no error was set.
func (b baseError) OrigErrs() []error {
	return b.errs
}
