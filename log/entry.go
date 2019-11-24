package log

// Entry represents a single log entry
type Entry interface {
	getField(index int) string
}

// simpleEntry is a private implementation of Entry
type simpleEntry struct {
	raw string
}
