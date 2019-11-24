package fieldparser

import "regexp"

// SimpleFieldParserConfig defines some configurability
// that determines how the field parser identifies fields
type SimpleFieldParserConfig struct {
	fieldRegex string
}

// SimpleFieldParser is an implementation of FieldParser
// that can parse the example logs in the DigIO programming task
type SimpleFieldParser struct {
	config *SimpleFieldParserConfig
}

// ParseFields parses a raw log entry based on the provided config
func (p *SimpleFieldParser) ParseFields(rawLogEntry *string) map[string]interface{} {

	splitter, err := regexp.Compile(p.config.fieldRegex)

}
