package err

import "errors"

// NewErrorOption is a type for options that can be passed to NewError and NewBatchError functions.
type NewErrorOption func(*BaseError)

// WithErrors includes errors in the new error object as original errors.
func WithErrors(errs ...error) NewErrorOption {
	return func(b *BaseError) {
		if errs != nil && len(errs) > 0 {
			if b.errs == nil {
				b.errs = make([]error, 0, len(errs))
			}
			b.errs = append(b.errs, errs...)
		}
	}
}

// A BaseError wraps the code and message which defines an error. It also can be used to wrap an original error object.
// Should be used as the root for errors satisfying the Error interface.
// Also for any error which does not fit into a specific error wrapper type.
type BaseError struct {
	// code is a short no whitespace phrase depicting the classification of the error that is being created.
	code string
	// message is the free flow string containing detailed information about the error.
	message string
	// errs is the error objects which will be nested under the new errors to be returned. Allows building chained errors.
	errs []error
}

// NewBaseError returns an error object for the code, message, and errors.
func NewBaseError(code, message string, options ...NewErrorOption) *BaseError {
	b := &BaseError{
		code:    code,
		message: message,
		errs:    make([]error, 0),
	}
	for _, option := range options {
		if option != nil {
			option(b)
		}
	}
	return b
}

// Error returns the string representation of the error.
func (b BaseError) Error() string {
	if len(b.errs) > 0 {
		return sprintError(b.code, b.message, withOriginalErr(errorList(b.errs)))
	}
	return sprintError(b.code, b.message)
}

// String returns the string representation of the error. It's an alias for Error to satisfy the stringer interface.
func (b BaseError) String() string {
	return b.Error()
}

// Code returns the short phrase that describes the error classification.
func (b BaseError) Code() string {
	return b.code
}

// Message returns the error details message.
func (b BaseError) Message() string {
	return b.message
}

// OrigErr returns the original error if one was set.
// Nil is returned if no error was set. If the full list is needed, use BatchError.
func (b BaseError) OrigErr() error {
	switch len(b.errs) {
	case 0:
		return nil
	case 1:
		return b.errs[0]
	default:
		var err Error
		if errors.As(b.errs[0], &err) {
			return NewBatchError(err.Code(), err.Message(), WithErrors(b.errs[1:]...))
		}
		return NewBatchError("BatchError", "multiple errors occurred", WithErrors(b.errs...))
	}
}

// Unwrap allows errors wrapped by BaseError to be compatible with standard error unwrapping mechanism.
func (b BaseError) Unwrap() error {
	return b.OrigErr()
}

// OrigErrs returns the original errors if at least one was set. An empty slice is returned if no error was set.
func (b BaseError) OrigErrs() []error {
	return b.errs
}
