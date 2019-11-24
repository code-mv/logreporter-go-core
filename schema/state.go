package schema

// simpleFieldState is a key/value pair
type simpleFieldState struct {
	key   string
	value string
}

// GetKey returns the key
func (s *simpleFieldState) GetKey() string {
	return s.key
}

// GetValue returns the value
func (s *simpleFieldState) GetValue() string {
	return s.value
}
