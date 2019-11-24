package regex

import (
	"regexp"

	"github.com/code-mv/logreporter-go-core/utils/errors"
)

// GetCaptureWithName returns the match with the expected name or an error
func GetCaptureWithName(regEx, input, expectedName string) (string, error) {

	// Get all named captures
	namedCaptures := GetAllNamedCaptures(regEx, input)

	// Check if the expected capture name exists
	val, ok := namedCaptures[expectedName]

	// If expected capture name does exist
	if !ok {
		// Return error
		return "", errors.Newf(errors.RegexCaptureNotFoundError, "Did not find capture with name = %s", expectedName)
	}

	// Return captured value
	return val, nil

}

// MustGetCaptureWithName either returns the match with the expectedName
// or returns an error
func MustGetCaptureWithName(regEx, input, expectedName string) string {

	// Get captured value or error
	capturedVal, err := GetCaptureWithName(regEx, input, expectedName)

	// Throw if error is not nil
	errors.ThrowOnErrorf(err, errors.RegexCaptureNotFoundError, "Could not find capture with name = %s", expectedName)

	// Return captured value
	return capturedVal

}

// GetAllNamedCaptures returns all named regex matches
func GetAllNamedCaptures(regEx, input string) map[string]string {

	// Compile regex
	var r, err = regexp.Compile(regEx)

	// Throw error if not null
	errors.ThrowOnErrorf(err, errors.RegexCompileError, "Could not compile regex from text = %s", regEx)

	// Find submatches
	match := r.FindStringSubmatch(input)

	// Create param map
	paramMap := make(map[string]string)

	// Interate regex group names
	for i, name := range r.SubexpNames() {
		// If match index is in rage
		if i > 0 && i <= len(match) {
			// Assign match value to param map
			paramMap[name] = match[i]
		}
	}

	// Return param map
	return paramMap
}
