package schema

// FieldParser is a function that sanitises a particular field
type FieldParser func(string) string

// simpleFieldDefinition is a grouping of fieldName and corresponding sanitiser
type simpleFieldDefinition struct {
	Name        string
	FieldParser FieldParser
}

// GetState returns the key and value of the state of a field
func (s *simpleFieldDefinition) GetState(value string) *simpleFieldState {

	return &simpleFieldState{
		key:   s.Name,
		value: s.FieldParser(value),
	}

}
