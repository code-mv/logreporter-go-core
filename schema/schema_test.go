package schema

import "testing"

func TestGetFieldDefinitions(t *testing.T) {

	schema := NewLogSchema()
	schema.GetFieldDefinitions()

}
