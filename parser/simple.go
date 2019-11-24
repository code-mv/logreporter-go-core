package parser

import (
	"regexp"

	"github.com/code-mv/logreporter-go-core/schema"
	"github.com/code-mv/logreporter-go-core/utils/errors"
)

// SimpleLogEntryParserConfig defines some configurability
// that determines how the field parser identifies fields
type SimpleLogEntryParserConfig struct {
	fieldMatcher string
}

// SimpleLogEntryParser is an implementation of FieldParser
// that can parse the example logs in the DigIO programming task
type SimpleLogEntryParser struct {
	schema schema.LogSchema
	config *SimpleLogEntryParserConfig
}

// Parse parses a raw log entry based on the provided config
func (p *SimpleLogEntryParser) Parse(rawLogEntry *string) map[string]string {

	// Check entry conditions
	errors.CheckMandatoryFields(rawLogEntry)

	// Try to compile the configured regex for identifying fields
	matcher, err := regexp.Compile(p.config.fieldMatcher)

	// Panic if an error occurred
	errors.ThrowOnErrorf(err, errors.RegexCompileError, "Failed to compile regex string = %s", p.config.fieldMatcher)

	// Split into fields
	allFields := matcher.FindAllString(*rawLogEntry, -1)

	// Field values mapped to field names
	fieldMap := p.schema.MapFields(allFields)

	// Return field map
	return fieldMap

}

// NewLogEntryParser returns a simple implementation
// of the LogEntryParse interface
func NewLogEntryParser(schema schema.LogSchema) LogEntryParser {
	// Return log entry parser
	return &SimpleLogEntryParser{
		config: &SimpleLogEntryParserConfig{
			fieldMatcher: `[^\s"'\[\]]+|"([^"]*)"|'([^']*)'|\[([^"]*)\]`,
		},
		schema: schema,
	}
}
