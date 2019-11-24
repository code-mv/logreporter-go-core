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

// AddAll adds all map entries from a source map to a target map
func AddAll(source map[string]string, target map[string]string) {
	for k, v := range source {
		target[k] = v
	}
}
