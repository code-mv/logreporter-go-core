package parser

import (
	"regexp"

	"github.com/code-mv/logreporter-go-core/fields"

	"github.com/code-mv/logreporter-go-core/errors"
	"github.com/code-mv/logreporter-go-core/regex"
)

// fieldParser is a function that sanitises a particular field
type fieldParser func(string) string

// fieldConfig is a grouping of fieldName and corresponding sanitiser
type fieldConfig struct {
	name        string
	fieldParser fieldParser
}

// fieldNames is a slice containing an ordered set of field names
var fieldNames = []fieldConfig{
	fieldConfig{
		name:        fields.IPAddress,
		fieldParser: func(s string) string { return s },
	},
	fieldConfig{
		name:        fields.UnknownField1,
		fieldParser: func(s string) string { return s },
	},
	fieldConfig{
		name:        fields.Username,
		fieldParser: func(s string) string { return s },
	},
	fieldConfig{
		name: fields.Timestamp,
		fieldParser: func(s string) string {
			return regex.MustGetCaptureWithName(`\[(?P<timestamp>[^\[\]]+)\]`, s, "timestamp")
		},
	},
	fieldConfig{
		name: fields.HTTPDetails,
		fieldParser: func(s string) string {

			return regex.MustGetCaptureWithName(`\"(?P<httpDetails>[^\"]+)\"`, s, "httpDetails")

		},
	},
	fieldConfig{
		name:        fields.HTTPStatusCode,
		fieldParser: func(s string) string { return s },
	},
	fieldConfig{
		name:        fields.UnknownNumber,
		fieldParser: func(s string) string { return s },
	},
	fieldConfig{
		name: fields.UnknownField2,
		fieldParser: func(s string) string {
			return regex.MustGetCaptureWithName(`\"(?P<unknownField2>[^\"]+)\"`, s, "unknownField2")
		},
	},
	fieldConfig{
		name: fields.UserAgentDetails,
		fieldParser: func(s string) string {
			return regex.MustGetCaptureWithName(`\"(?P<userAgentDetails>[^\"]+)\"`, s, "userAgentDetails")
		},
	},
}

// SimpleFieldParserConfig defines some configurability
// that determines how the field parser identifies fields
type SimpleFieldParserConfig struct {
	fieldMatcher string
}

// SimpleFieldParser is an implementation of FieldParser
// that can parse the example logs in the DigIO programming task
type SimpleFieldParser struct {
	config *SimpleFieldParserConfig
}

// Parse parses a raw log entry based on the provided config
func (p *SimpleFieldParser) Parse(rawLogEntry *string) map[string]interface{} {

	// Check entry conditions
	errors.CheckMandatoryFields(rawLogEntry)

	// Try to compile the configured regex for identifying fields
	matcher, err := regexp.Compile(p.config.fieldMatcher)

	// Panic if an error occurred
	errors.ThrowOnErrorf(err, errors.RegexCompileError, "Failed to compile regex string = %s", p.config.fieldMatcher)

	// Split into fields
	allFields := matcher.FindAllString(*rawLogEntry, -1)

	// Field values mapped to field names
	fieldMap := mapNamedFields(allFields)

	return fieldMap

}

// mapNamedFields maps the important, named fields into
// a map of field names and values
func mapNamedFields(allFields []string) map[string]interface{} {

	// Check entry conditions
	errors.CheckMandatoryFields(allFields)

	// Allocate map for result
	result := make(map[string]interface{})

	// Iterate across all fields
	for i, value := range allFields {

		if i < len(fieldNames) {
			// Get corresponding field config
			fieldConf := fieldNames[i]

			// Parse field and add value to map
			result[fieldConf.name] = fieldConf.fieldParser(value)
		}

	}

	// Return map result
	return result

}

// NewParser returns a simple implementation
// of the LogEntryParse interface
func NewParser() LogEntryParser {
	return &SimpleFieldParser{
		config: &SimpleFieldParserConfig{
			fieldMatcher: `[^\s"'\[\]]+|"([^"]*)"|'([^']*)'|\[([^"]*)\]`,
		},
	}
}
