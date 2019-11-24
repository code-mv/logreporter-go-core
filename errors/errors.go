package errors

import (
	"log"

	"github.com/pkg/errors"
)

// ErrorCode is an identifier for a specific type of error
type ErrorCode string

const (
	// errFormat is the format of our standard error message
	errFormat = "Encountered error with code = %s"
	// RecoveryError errorcode
	RecoveryError ErrorCode = "001"
	// OpenFileError errorcode
	OpenFileError ErrorCode = "002"
)

// RecoverFunc is a callback function that gets invoked
// after a successful recovery
type RecoverFunc func()

// Rethrow wraps and rethrows an error
func Rethrow(err error, errCode ErrorCode) error {
	// Log message
	log.Printf(errFormat, errCode)
	// Wrap and return error
	return errors.Wrapf(err, errFormat, errCode)
}

// Panic logs and panics based on the given error
func Panic(err error, errCode string) {
	// Log message
	log.Printf(errFormat, errCode)
	// panic
	panic(errCode)
}

// Recover recovers for a particular error code
func Recover(expectedErrCode ...ErrorCode) {

	if r := recover(); r != nil {

		if utils.expectedErrCode != r {
			panic(RecoveryError)
		}

	}

}
