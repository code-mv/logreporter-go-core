package schema

import (
	"github.com/code-mv/logreporter-go-core/utils"
	"github.com/code-mv/logreporter-go-core/utils/regex"
)

// simpleLogSchema is a simple implementation of
// LogSchema
type simpleLogSchema struct {
	fieldMappers []FieldMapper
}

// GetFieldDefinitions returns a set of field definitions
// comprising the schema
func (s *simpleLogSchema) MapFields(positionalValues []string) map[string]string {

	result := make(map[string]string)

	// Iterate across all fields
	for i, mapper := range s.fieldMappers {
		value := positionalValues[i]
		newFields := mapper.Map(value)
		utils.AddAll(newFields, result)
	}

	return result

}

// NewLogSchema returns a simple implementation of LogSchema
func NewLogSchema() LogSchema {
	return &simpleLogSchema{
		fieldMappers: []FieldMapper{
			&simpleFieldMapper{
				FieldDefs: []*simpleFieldDefinition{
					&simpleFieldDefinition{
						Name:        IPAddress,
						FieldParser: func(s string) string { return s },
					},
				},
			},
			&simpleFieldMapper{
				FieldDefs: []*simpleFieldDefinition{
					&simpleFieldDefinition{
						Name:        UnknownField1,
						FieldParser: func(s string) string { return s },
					},
				},
			},
			&simpleFieldMapper{
				FieldDefs: []*simpleFieldDefinition{
					&simpleFieldDefinition{
						Name:        Username,
						FieldParser: func(s string) string { return s },
					},
				},
			},
			&simpleFieldMapper{
				FieldDefs: []*simpleFieldDefinition{
					&simpleFieldDefinition{
						Name: Timestamp,
						FieldParser: func(s string) string {
							return regex.MustGetCaptureWithName(`\[(?P<timestamp>[^\[\]]+)\]`, s, "timestamp")
						},
					},
				},
			},
			&simpleFieldMapper{
				FieldDefs: []*simpleFieldDefinition{
					&simpleFieldDefinition{
						Name: HTTPMethod,
						FieldParser: func(s string) string {
							return regex.MustGetCaptureWithName(`\"(?P<httpMethod>[^\"\s]+)[^\"]+\"`, s, "httpMethod")
						},
					},
					&simpleFieldDefinition{
						Name: URLPath,
						FieldParser: func(s string) string {
							return regex.MustGetCaptureWithName(`\"[^\"\s]+\s+(?P<urlPath>[^\"\s]+)[^\"]+\"`, s, "urlPath")
						},
					},
					&simpleFieldDefinition{
						Name: HTTPVersion,
						FieldParser: func(s string) string {
							return regex.MustGetCaptureWithName(`\"([^\"\s]+\s+){2}(?P<httpVersion>[^\"\s]+)[^\"]*\"`, s, "httpVersion")
						},
					},
				},
			},
			&simpleFieldMapper{
				FieldDefs: []*simpleFieldDefinition{
					&simpleFieldDefinition{
						Name:        HTTPStatusCode,
						FieldParser: func(s string) string { return s },
					},
				},
			},
			&simpleFieldMapper{
				FieldDefs: []*simpleFieldDefinition{
					&simpleFieldDefinition{
						Name:        UnknownNumber,
						FieldParser: func(s string) string { return s },
					},
				},
			},
			&simpleFieldMapper{
				FieldDefs: []*simpleFieldDefinition{
					&simpleFieldDefinition{
						Name: UnknownField2,
						FieldParser: func(s string) string {
							return regex.MustGetCaptureWithName(`\"(?P<unknownField2>[^\"]+)\"`, s, "unknownField2")
						},
					},
				},
			},
			&simpleFieldMapper{
				FieldDefs: []*simpleFieldDefinition{
					&simpleFieldDefinition{
						Name: UserAgentDetails,
						FieldParser: func(s string) string {
							return regex.MustGetCaptureWithName(`\"(?P<userAgentDetails>[^\"]+)\"`, s, "userAgentDetails")
						},
					},
				},
			},
		},
	}
}
