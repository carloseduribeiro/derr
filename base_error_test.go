package derr

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWithError(t *testing.T) {
	t.Run("should add a single error to new BaseError", func(t *testing.T) {
		// given
		origErr := errors.New("original error")
		err := NewBaseError("code", "message", WithErrors(origErr))
		// when
		result := err.OrigErr()
		// then
		assert.Equal(t, origErr, result)
	})

	t.Run("should not add the error to the new BaseError when the error is nil", func(t *testing.T) {
		// given
		err := NewBaseError("code", "message", WithErrors(nil))
		// when
		result := err.OrigErr()
		// then
		assert.Nil(t, result)
	})
}

func TestWithErrorS(t *testing.T) {
	t.Run("should add multiple errors to the new BaseError", func(t *testing.T) {
		// given
		origErrs := []error{errors.New("first error"), errors.New("second error")}
		err := NewBaseError("code", "message", WithErrors(origErrs...))
		// when
		result := err.OrigErr()
		var batchErr BatchError
		ok := errors.As(result, &batchErr)
		// then
		assert.True(t, ok)
		assert.Equal(t, "BatchError", batchErr.Code())
		assert.Equal(t, "multiple errors occurred", batchErr.Message())
		assert.Equal(t, origErrs, batchErr.OrigErrs())
	})

	t.Run("shouldNotAddErrorsToBaseErrorWhenErrorsAreNil", func(t *testing.T) {
		// given
		err := NewBaseError("code", "message", WithErrors(nil))
		// when
		result := err.OrigErr()
		// then
		assert.Nil(t, result)
	})

	t.Run("shouldNotAddErrorsToBaseErrorWhenErrorsAreEmpty", func(t *testing.T) {
		// given
		err := NewBaseError("code", "message", WithErrors([]error{}...))
		// when
		result := err.OrigErr()
		// then
		assert.Nil(t, result)
	})
}

func TestBaseError_Error(t *testing.T) {
	t.Run("should return a formatted error message when no there are original error options", func(t *testing.T) {
		// given
		err := NewBaseError("code", "message")
		// when
		result := err.Error()
		// then
		assert.Equal(t, "code: message", result)
	})

	t.Run("should return a formatted error message with original error when there is one original error", func(t *testing.T) {
		// given
		origErr := errors.New("original error")
		err := NewBaseError("code", "message", WithErrors(origErr))
		// when
		result := err.Error()
		// then
		assert.Equal(t, "code: message\ncaused by: original error", result)
	})

	t.Run("should return a formatted error message with original errors when there are multiple original errors", func(t *testing.T) {
		// given
		origErrs := []error{errors.New("first error"), errors.New("second error")}
		err := NewBaseError("code", "message", WithErrors(origErrs...))
		// when
		result := err.Error()
		// then
		assert.Equal(t, "code: message\ncaused by: first error\nsecond error", result)
	})
}

func TestBaseError_String(t *testing.T) {
	t.Run("should return formatted error message when no original errors with the same result of Error method", func(t *testing.T) {
		// given
		err := NewBaseError("code", "message")
		// when
		errResult := err.Error()
		strResult := err.String()
		// then
		assert.Equal(t, "code: message", errResult)
		assert.Equal(t, "code: message", strResult)
	})

	t.Run("should return formatted error message with original errors when multiple original errors with the same result of Error method", func(t *testing.T) {
		// given
		origErrs := []error{errors.New("first error"), errors.New("second error")}
		err := NewBaseError("code", "message", WithErrors(origErrs...))
		// when
		errResult := err.Error()
		strResult := err.String()
		// then
		assert.Equal(t, "code: message\ncaused by: first error\nsecond error", errResult)
		assert.Equal(t, "code: message\ncaused by: first error\nsecond error", strResult)

	})
}

func TestBaseError_OrigErr(t *testing.T) {
	t.Run("should return nil when there are no original errors", func(t *testing.T) {
		// given
		err := NewBaseError("code", "message")
		// when
		result := err.OrigErr()
		// then
		assert.Nil(t, result)
	})

	t.Run("should return a single error when there is only one original error", func(t *testing.T) {
		// given
		origErr := errors.New("original error")
		err := NewBaseError("code", "message", WithErrors(origErr))
		// when
		result := err.OrigErr()
		// then
		assert.Equal(t, origErr, result)
	})

	t.Run("should return BatchError when there are multiple original errors", func(t *testing.T) {
		// given
		origErrs := []error{errors.New("first error"), errors.New("second error")}
		err := NewBaseError("code", "message", WithErrors(origErrs...))
		// when
		result := err.OrigErr()
		var batchErr BatchError
		ok := errors.As(result, &batchErr)
		// then
		assert.True(t, ok)
		assert.Equal(t, "BatchError", batchErr.Code())
		assert.Equal(t, "multiple errors occurred", batchErr.Message())
		assert.Equal(t, origErrs, batchErr.OrigErrs())
	})

	t.Run("should return BatchError with first error details when there are multiple original errors and the first error is of type error", func(t *testing.T) {
		// given
		origErrs := []error{NewBaseError("first_error_code", "first error message", nil), errors.New("second error")}
		err := NewBaseError("code", "message", WithErrors(origErrs...))
		// when
		result := err.OrigErr()
		var batchErr BatchError
		ok := errors.As(result, &batchErr)
		// then
		assert.True(t, ok)
		assert.Equal(t, "first_error_code", batchErr.Code())
		assert.Equal(t, "first error message", batchErr.Message())
		assert.Equal(t, origErrs[1:], batchErr.OrigErrs())
	})
}
