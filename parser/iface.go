package parser

// LogEntryParser parses the individual fields within a log entry
type LogEntryParser interface {
	// Takes in a raw log entry and parses it into
	// a map of field names and values
	Parse(rawLogEntry *string) map[string]interface{}
}
