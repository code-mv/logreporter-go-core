package errors

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
)

func TestNew(t *testing.T) {

	// Define error message
	errMsg := "Test error"

	// Create new error
	err := New(RuntimeError, errMsg)

	// Build the expected error msg
	expectedErrorMsg := fmt.Sprintf(errFormat, RuntimeError, errMsg)

	if err.Error() != expectedErrorMsg {
		t.Errorf("Expected error message to be in standard format but was = %s", err.Error())
	}

}

func TestNewf(t *testing.T) {

	// Define error message
	errMsg := "Test error %s"

	// Define an error message arg
	errMsgArg := "message"

	// Create new error
	err := Newf(RuntimeError, errMsg, errMsgArg)

	// Build the expected error msg
	expectedErrorMsg := fmt.Sprintf(errFormat, RuntimeError, fmt.Sprintf(errMsg, errMsgArg))

	if err.Error() != expectedErrorMsg {
		t.Errorf("Expected error message to be in standard format but was = %s", err.Error())
	}

}

func TestWrapOnErrorPositive(t *testing.T) {

	// Create new error
	err := errors.New("Test error")

	// Wrap error
	wrapped := WrapOnError(err, OpenFileError, "New error")

	// Fail if rethrow cause is not the original error
	if err != errors.Cause(wrapped) {
		t.Errorf("Expected cause to be the original error but it wasn't")
	}

}

func TestWrapOnErrorNegativeNotWrapped(t *testing.T) {

	// Create new error
	err := errors.New("Test error")

	// Rethrow error
	err2 := errors.New("Test error 2")

	// Fail if rethrow cause is not the original error
	if err == errors.Cause(err2) {
		t.Errorf("Expected cause to NOT be the original error but it was")
	}

}

func TestWrapOnErrorNegativeNilError(t *testing.T) {

	// Wrap nil error
	wrapped := WrapOnError(nil, OpenFileError, "New error")

	// Fail if rethrow cause is not the original error
	if wrapped != nil {
		t.Errorf("Expected wrapped exception to be nil")
	}

}

func TestThrowOnErrorCatchPositive(t *testing.T) {

	// Set is failed to true
	isFailed := true

	// Defer function that determines pass or fail
	defer func() {
		if isFailed {
			t.Errorf("Is failed flag is set to true")
		}
	}()

	// Defer catch function
	defer Catch(func(r interface{}) {
		isFailed = false
	}, OpenFileError)

	// Create new error
	err := errors.New("Test error")

	// Throw error
	ThrowOnError(err, OpenFileError, "New error")

}

func TestThrowOnErrorCatchNegativeNilError(t *testing.T) {

	// Defer catch function
	defer Catch(func(r interface{}) {
		t.Fail()
	}, OpenFileError)

	// Panic nil error
	ThrowOnError(nil, OpenFileError, "New error")

}

func TestThrowOnErrorfCatchPositive(t *testing.T) {

	// Set is failed to true
	isFailed := true

	// Defer function that determines pass or fail
	defer func() {
		if isFailed {
			t.Errorf("Is failed flag is set to true")
		}
	}()

	// Defer catch function
	defer Catch(func(r interface{}) {
		isFailed = false
	}, OpenFileError)

	// Create new error
	err := errors.New("Test error")

	// Throw error
	ThrowOnErrorf(err, OpenFileError, "New error %s", "message")

}

func TestThrowOnErrorfCatchNegativeNilError(t *testing.T) {

	// Defer catch function
	defer Catch(func(r interface{}) {
		t.Fail()
	}, OpenFileError)

	// Panic nil error
	ThrowOnErrorf(nil, OpenFileError, "New error %s", "message")

}
