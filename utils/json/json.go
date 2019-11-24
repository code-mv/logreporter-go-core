package json

import (
	"encoding/json"

	"github.com/code-mv/logreporter-go-core/utils/errors"
)

// ToJSON marshalls any object into JSON
func ToJSON(object interface{}) string {

	// Check preconditions
	errors.CheckMandatoryFields(object)

	// Try to marhshal object to bytes
	bytes, err := json.Marshal(object)

	// Throw if there's an error
	errors.ThrowOnError(err, errors.JSONMarshalError, "Failed to marshal object to json")

	return string(bytes)

}

// ToPrettyJSON marshalls any object into JSON
func ToPrettyJSON(object interface{}) string {

	// Check preconditions
	errors.CheckMandatoryFields(object)

	// Try to marhshal object to bytes
	bytes, err := json.MarshalIndent(object, "", "  ")

	// Throw if there's an error
	errors.ThrowOnError(err, errors.JSONMarshalError, "Failed to marshal object to json")

	return string(bytes)

}
