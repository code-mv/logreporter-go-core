package schema

// LogSchema defines a set of field definitions
// comprising a schema of fields in a log entry
type LogSchema interface {
	GetFieldDefinitions() []FieldDefinition
}
