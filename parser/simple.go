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
func (p *SimpleLogEntryParser) Parse(rawLogEntry *string) map[string]interface{} {

	// Check entry conditions
	errors.CheckMandatoryFields(rawLogEntry)

	// Try to compile the configured regex for identifying fields
	matcher, err := regexp.Compile(p.config.fieldMatcher)

	// Panic if an error occurred
	errors.ThrowOnErrorf(err, errors.RegexCompileError, "Failed to compile regex string = %s", p.config.fieldMatcher)

	// Split into fields
	allFields := matcher.FindAllString(*rawLogEntry, -1)

	// Field values mapped to field names
	fieldMap := mapNamedFields(allFields, p.schema.GetFieldDefinitions())

	return fieldMap

}

// mapNamedFields maps the important, named fields into
// a map of field names and values
func mapNamedFields(allFields []string, fieldDefs []schema.FieldDefinition) map[string]interface{} {

	// Check entry conditions
	errors.CheckMandatoryFields(allFields)

	// Allocate map for result
	result := make(map[string]interface{})

	// Iterate across all fields
	for i, value := range allFields {

		if i < len(fieldDefs) {
			// Get corresponding field def
			fieldDef := fieldDefs[i]

			// Parse field and add value to map
			result[fieldDef.Name] = fieldDef.FieldParser(value)
		}

	}

	// Return map result
	return result

}

// NewLogEntryParser returns a simple implementation
// of the LogEntryParse interface
func NewLogEntryParser(schema schema.LogSchema) LogEntryParser {
	return &SimpleLogEntryParser{
		config: &SimpleLogEntryParserConfig{
			fieldMatcher: `[^\s"'\[\]]+|"([^"]*)"|'([^']*)'|\[([^"]*)\]`,
		},
		schema: schema,
	}
}
