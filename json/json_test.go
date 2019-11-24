package json

import (
	"testing"

	"github.com/code-mv/logreporter-go-core/errors"
)

func TestToJSON(t *testing.T) {

	// Defer catch function
	defer errors.Catch(func(r interface{}) {
		t.Errorf("Caught panic with message = %s", r)
	}, errors.RegexCaptureNotFoundError)

	// Create map
	object := make(map[string]interface{})

	// Add some values
	object["field1"] = "value1"
	object["field2"] = 10.0

	json := ToJSON(object)

	t.Log(json)

}
