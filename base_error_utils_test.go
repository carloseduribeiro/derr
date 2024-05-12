package derr

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSprintError(t *testing.T) {
	t.Run("should return the formated error code with extra information", func(t *testing.T) {
		// given
		code := "some_error"
		message := "error doing something"
		extra := "test of test"
		expected := "some_error: error doing something\n\ttest of test"
		// when
		result := sprintError(code, message, withExtra(extra))
		// then
		assert.Equal(t, expected, result)
	})

	t.Run("should return the formated error code with the original error", func(t *testing.T) {
		// given
		code := "some_error"
		message := "error doing something"
		original := errors.New("original error")
		expected := "some_error: error doing something\ncaused by: original error"
		// when
		result := sprintError(code, message, withOriginalErr(original))
		// then
		assert.Equal(t, expected, result)
	})

	t.Run("should return the formated error code with extra information and the original error", func(t *testing.T) {
		// given
		code := "some_error"
		message := "error doing something"
		extra := "test of test"
		originalErr := errors.New("original error")
		expected := "some_error: error doing something\n\ttest of test\ncaused by: original error"
		// when
		result := sprintError(code, message, withExtra(extra), withOriginalErr(originalErr))
		// then
		assert.Equal(t, expected, result)
	})

	t.Run("should return the formated error code with code and message only", func(t *testing.T) {
		// given
		code := "some_error"
		message := "error doing something"
		expected := "some_error: error doing something"
		// when
		result := sprintError(code, message)
		// then
		assert.Equal(t, expected, result)
	})
}

func TestErrorList_Error(t *testing.T) {
	t.Run("should return empty string when error list is empty", func(t *testing.T) {
		// given
		errList := errorList{}
		expected := ""
		// when
		result := errList.Error()
		// then
		assert.Equal(t, expected, result)
	})

	t.Run("should return single error message when error list contains one error", func(t *testing.T) {
		// given
		errList := errorList{errors.New("single error")}
		expected := "single error"
		// when
		result := errList.Error()
		// then
		assert.Equal(t, expected, result)
	})

	t.Run("should return concatenated error messages when error list contains multiple errors", func(t *testing.T) {
		// given
		errList := errorList{errors.New("first error"), errors.New("second error")}
		expected := "first error\nsecond error"
		// when
		result := errList.Error()
		// then
		assert.Equal(t, expected, result)
	})
}
