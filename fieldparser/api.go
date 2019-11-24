package fieldparser

// FieldParser parses the individual fields within a log entry
type FieldParser interface {
	// Takes in a raw log entry and parses it into
	// a map of field names and values
	ParseFields(rawLogEntry *string) map[string]interface{}
}
