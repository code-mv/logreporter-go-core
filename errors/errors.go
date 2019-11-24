package errors

import (
	"fmt"
	"log"

	"github.com/code-mv/logreporter-go-core/utils"

	"github.com/pkg/errors"
)

// ErrorCode is an identifier for a specific type of error
type ErrorCode string

const (
	// errFormat is the format of our standard error message
	errFormat = "Error code %s :: %s"
	// RuntimeError errorcode
	RuntimeError ErrorCode = "001"
	// RecoveryError errorcode
	RecoveryError ErrorCode = "002"
	// OpenFileError errorcode
	OpenFileError ErrorCode = "003"
	// RegexCompileError errorcode
	RegexCompileError ErrorCode = "004"
	// RegexCaptureNotFoundError errorcode
	RegexCaptureNotFoundError ErrorCode = "005"
	// JSONMarshalError errorcode
	JSONMarshalError ErrorCode = "006"
)

// RecoveryFunc is a callback function that gets invoked
// after a successful recovery
type RecoveryFunc func(interface{})

// New creates a new error quoting the provided error code
func New(errCode ErrorCode, msg string) error {
	// Return error with standard message format
	return fmt.Errorf(errFormat, errCode, msg)
}

// Newf creates a new error quoting the provided error code
func Newf(errCode ErrorCode, msg string, args ...interface{}) error {
	// Return error with standard message format
	return fmt.Errorf(errFormat, errCode, fmt.Sprintf(msg, args...))
}

// WrapOnError wraps and rethrows if the given error is not nil
func WrapOnError(err error, errCode ErrorCode, msg string) error {
	// Return nil if error is nil
	if err == nil {
		return nil
	}
	// Log message
	log.Printf(errFormat, errCode, msg)
	// Wrap and return error
	return errors.Wrapf(err, errFormat, errCode, msg)
}

// ThrowOnError logs and panics if the given error is not nil
func ThrowOnError(err error, errCode ErrorCode, msg string) {
	if err != nil {
		// Log message
		log.Printf(errFormat, errCode, msg)
		// panic
		panic(errCode)
	}
}

// ThrowOnErrorf logs and panics if the given error is not nil
func ThrowOnErrorf(err error, errCode ErrorCode, msg string, args ...interface{}) {
	if err != nil {
		// Log message
		log.Printf(errFormat, errCode, fmt.Sprintf(msg, args...))
		// panic
		panic(errCode)
	}
}

// Catch recovers for a particular error code
func Catch(recovery RecoveryFunc, expectedErrCode ...ErrorCode) {

	// Attempt recovery
	if r := recover(); r != nil {

		// Check if actual error code is not one of the expected ones
		if !utils.ItemInSlice(r, errorCodesToInterfaces(expectedErrCode)) {
			// Re-panic if error code is not one of the expected ones
			panic(RecoveryError)
		}

		recovery(r)
		return

	}

}

// CheckMandatoryFields panics with an error if any of the
// mandatory fields are null
func CheckMandatoryFields(fields ...interface{}) {

	// Iterate values of mandatory fields
	for _, value := range fields {

		if value == nil {
			ThrowOnError(errors.New(""), RuntimeError, "Missing one or more mandatory fields")
		}

	}
}

// errorCodesToInterfaces converts an array of strings to
// an array of interfaces
func errorCodesToInterfaces(errCodes []ErrorCode) []interface{} {

	// Create a slice of interfaces
	new := make([]interface{}, len(errCodes))

	// Iterate over the slice of errCodes
	for i, v := range errCodes {
		// Assign the error code values to the interface slice
		new[i] = v
	}

	// Return slice of interfaces
	return new

}
