package utils

// ItemInSlice is a utility function that returns true
// if list contains the interface a
func ItemInSlice(a interface{}, list []interface{}) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
