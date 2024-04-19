package err

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewError(t *testing.T) {
	t.Run("should return an error with the given code message and the original error", func(t *testing.T) {
		// given
		origErr := errors.New("original error")
		// when
		err := NewError("code", "message", WithError(origErr))
		// then
		assert.Equal(t, "code", err.Code())
		assert.Equal(t, "message", err.Message())
		assert.Equal(t, origErr, err.OrigErr())
	})

	t.Run("should return an error with nil original error when none is provided", func(t *testing.T) {
		// when
		err := NewError("code", "message")
		// then
		assert.Equal(t, "code", err.Code())
		assert.Equal(t, "message", err.Message())
		assert.Nil(t, err.OrigErr())
	})
}

func TestNewBatchError(t *testing.T) {
	t.Run("should return BatchError with given code and message and original errors", func(t *testing.T) {
		// given
		origErrs := []error{errors.New("first error"), errors.New("second error")}
		// when
		err := NewBatchError("code", "message", WithErrorS(origErrs))
		var batchErr BatchError
		ok := errors.As(err, &batchErr)
		// then
		assert.True(t, ok)
		assert.Equal(t, "code", batchErr.Code())
		assert.Equal(t, "message", batchErr.Message())
		assert.Equal(t, origErrs, batchErr.OrigErrs())
	})

	t.Run("should return batch error with nil original errors when none are provided", func(t *testing.T) {
		// when
		err := NewBatchError("code", "message")
		var batchErr BatchError
		ok := errors.As(err, &batchErr)
		// then
		assert.True(t, ok)
		assert.Equal(t, "code", batchErr.Code())
		assert.Equal(t, "message", batchErr.Message())
		assert.Nil(t, batchErr.OrigErrs())
	})
}
