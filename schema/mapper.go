package schema

// simpleFieldMapper is a bucket for taking single values
// and translating them into one or more fields
type simpleFieldMapper struct {
	FieldDefs []*simpleFieldDefinition
}

// Maps a value to one or more derived fields
func (s *simpleFieldMapper) Map(value string) map[string]string {

	// Create a new map
	result := make(map[string]string)

	// Iterate over the fieldMapper field defs
	for _, v := range s.FieldDefs {
		// Get the state of the field def
		state := v.GetState(value)
		// Add the key and value of the field def to the result
		result[state.key] = state.value
	}

	// Return the result
	return result

}
