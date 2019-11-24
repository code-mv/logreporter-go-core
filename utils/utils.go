package utils

// StringInSlice is a utility function that returns true
// if list contains the string a
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
