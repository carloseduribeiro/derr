package err

// An Error wraps lower level errors with code, message and an original error.
type Error interface {
	error
	// Code returns the short phrase depicting the classification of the error.
	Code() string
	// Message returns the error details message.
	Message() string
	// OrigErr returns the original error if one was set.  Nil is returned if not set.
	OrigErr() error
}

// BatchError is a batch of errors which also wraps lower level errors with code, message, and original errors.
// Calling Error() will include all errors that occurred in the batch.
type BatchError interface {
	Error
	// OrigErrs returns the original errors if one was set.  Nil is returned if not set.
	OrigErrs() []error
}

// NewError returns an Error object described by the code, message and options.
// The options can be used to include extra information about the error, like the original error.
func NewError(code, message string, options ...NewErrorOption) Error {
	return newBaseError(code, message, options...)
}

// NewBatchError returns an BatchError with a collection of errors as an array of errors.
// The options can be used to include extra information about the error, like the original error chain.
func NewBatchError(code, message string, options ...NewErrorOption) BatchError {
	return newBaseError(code, message, options...)
}
