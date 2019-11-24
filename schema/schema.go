package schema

import (
	"github.com/code-mv/logreporter-go-core/utils/regex"
)

const (
	// IPAddress field
	IPAddress string = "ipAddress"
	// UnknownField1 field
	UnknownField1 string = "unknownField1"
	// Username field
	Username string = "username"
	// Timestamp field
	Timestamp string = "timestamp"
	// HTTPDetails field
	HTTPDetails string = "httpDetails"
	// HTTPStatusCode field
	HTTPStatusCode string = "httpStatusCode"
	// UnknownNumber field
	UnknownNumber string = "unknownNumber"
	// UnknownField2 field
	UnknownField2 string = "unknownField2"
	// UserAgentDetails field
	UserAgentDetails string = "userAgentDetails"
)

// FieldParser is a function that sanitises a particular field
type FieldParser func(string) string

// FieldDefinition is a grouping of fieldName and corresponding sanitiser
type FieldDefinition struct {
	Name        string
	FieldParser FieldParser
}

// simpleLogSchema is a simple implementation of
// LogSchema
type simpleLogSchema struct {
	fieldDefinitions []FieldDefinition
}

// GetFieldDefinitions returns a set of field definitions
// comprising the schema
func (s *simpleLogSchema) GetFieldDefinitions() []FieldDefinition {
	return s.fieldDefinitions
}

// NewLogSchema returns a simple implementation of LogSchema
func NewLogSchema() LogSchema {
	return &simpleLogSchema{
		fieldDefinitions: []FieldDefinition{
			FieldDefinition{
				Name:        IPAddress,
				FieldParser: func(s string) string { return s },
			},
			FieldDefinition{
				Name:        UnknownField1,
				FieldParser: func(s string) string { return s },
			},
			FieldDefinition{
				Name:        Username,
				FieldParser: func(s string) string { return s },
			},
			FieldDefinition{
				Name: Timestamp,
				FieldParser: func(s string) string {
					return regex.MustGetCaptureWithName(`\[(?P<timestamp>[^\[\]]+)\]`, s, "timestamp")
				},
			},
			FieldDefinition{
				Name: HTTPDetails,
				FieldParser: func(s string) string {

					return regex.MustGetCaptureWithName(`\"(?P<httpDetails>[^\"]+)\"`, s, "httpDetails")

				},
			},
			FieldDefinition{
				Name:        HTTPStatusCode,
				FieldParser: func(s string) string { return s },
			},
			FieldDefinition{
				Name:        UnknownNumber,
				FieldParser: func(s string) string { return s },
			},
			FieldDefinition{
				Name: UnknownField2,
				FieldParser: func(s string) string {
					return regex.MustGetCaptureWithName(`\"(?P<unknownField2>[^\"]+)\"`, s, "unknownField2")
				},
			},
			FieldDefinition{
				Name: UserAgentDetails,
				FieldParser: func(s string) string {
					return regex.MustGetCaptureWithName(`\"(?P<userAgentDetails>[^\"]+)\"`, s, "userAgentDetails")
				},
			},
		},
	}
}
