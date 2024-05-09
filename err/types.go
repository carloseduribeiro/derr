package err

type NotFoundError struct {
	*BaseError
}

func NewNotFoundError(code, message string, options ...NewErrorOption) *NotFoundError {
	return &NotFoundError{
		BaseError: NewBaseError(code, message, options...),
	}
}

type PreconditionsError struct {
	*BaseError
}

func NewPreconditionsError(code, message string, options ...NewErrorOption) *PreconditionsError {
	return &PreconditionsError{
		BaseError: NewBaseError(code, message, options...),
	}
}

type ValidationError struct {
	*BaseError
}

func NewValidationError(message string, options ...NewErrorOption) *ValidationError {
	return &ValidationError{
		BaseError: NewBaseError("ValidationError", message, options...),
	}
}

type InternalError struct {
	*BaseError
}

func NewInternalError(message string, options ...NewErrorOption) *InternalError {
	return &InternalError{
		BaseError: NewBaseError("InternalError", message, options...),
	}
}

type UnauthorizedError struct {
	*BaseError
}

func NewUnauthorizedError(code, message string, options ...NewErrorOption) *UnauthorizedError {
	return &UnauthorizedError{
		BaseError: NewBaseError(code, message, options...),
	}
}

type UnauthenticatedError struct {
	*BaseError
}

func NewUnauthenticatedError(code, message string, options ...NewErrorOption) *UnauthenticatedError {
	return &UnauthenticatedError{
		BaseError: NewBaseError(code, message, options...),
	}
}

type UnknownError struct {
	*BaseError
}

func NewUnknownError(message string, options ...NewErrorOption) *UnknownError {
	return &UnknownError{
		BaseError: NewBaseError("UnknownError", message, options...),
	}
}
