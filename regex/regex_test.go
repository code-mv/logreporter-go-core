package regex

import (
	"testing"

	"github.com/code-mv/logreporter-go-core/errors"
)

func TestGetAllNamedCapturesPositive(t *testing.T) {

	// Define regex
	regEx := `^(?P<field1>\d{4})-(?P<field2>[^-\d]+)`

	// Test input
	input := "4562-test"

	// Get result
	result := GetAllNamedCaptures(regEx, input)

	// Test for expected results
	if result["field1"] != "4562" || result["field2"] != "test" {
		t.Errorf("Named captures did not have the expected values")
	}

}

func TestGetCaptureWithName(t *testing.T) {

	// Defer catch function
	defer errors.Catch(func(r interface{}) {
		t.Errorf("Caught panic with message = %s", r)
	}, errors.RuntimeError)

	// Define regex
	regEx := `^(?P<field1>\d{4})-(?P<field2>[^-\d]+)`

	// Test input
	input := "4562-test"

	// Get capture 1
	capture1, err := GetCaptureWithName(regEx, input, "field1")

	// Throw error
	errors.ThrowOnError(err, errors.RuntimeError, "Error getting capture 1")

	// Get capture 2
	capture2, err := GetCaptureWithName(regEx, input, "field2")

	errors.ThrowOnError(err, errors.RuntimeError, "Error getting capture 2")

	// Test for expected results
	if capture1 != "4562" || capture2 != "test" {
		t.Errorf("Named captures did not have the expected values")
	}

}

func TestMustGetCaptureWithName(t *testing.T) {

	// Defer catch function
	defer errors.Catch(func(r interface{}) {
		t.Errorf("Caught panic with message = %s", r)
	}, errors.RegexCaptureNotFoundError)

	// Define regex
	regEx := `^(?P<field1>\d{4})-(?P<field2>[^-\d]+)`

	// Test input
	input := "4562-test"

	// Get capture 1
	capture1 := MustGetCaptureWithName(regEx, input, "field1")

	// Get capture 2
	capture2 := MustGetCaptureWithName(regEx, input, "field2")

	// Test for expected results
	if capture1 != "4562" || capture2 != "test" {
		t.Errorf("Named captures did not have the expected values")
	}

}
