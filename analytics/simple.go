package analytics

import "sort"

// valueMapContainer stores counts mapped to the
// values they are a count of
type valueMapContainer struct {
	valueMap map[string]int
}

// simpleAnalyticsContainer stores value counts
// mapped to the fields they correspond to
type simpleAnalyticsContainer struct {
	fieldMap map[string]*valueMapContainer
}

// AddStats adds a map of fields and values, parsed from
// a log file, to the stat container
func (s *simpleAnalyticsContainer) AddStats(fields map[string]string) {

	// Add a state for each field/value pair in the map
	for k, v := range fields {
		s.AddStat(k, v)
	}

}

// AddStat adds a single field and value to the analytics container
func (s *simpleAnalyticsContainer) AddStat(fieldName string, value string) {

	// Get fieldMap
	fieldMap := s.fieldMap

	// If fieldMap is nill
	if fieldMap == nil {
		fieldMap = make(map[string]*valueMapContainer)
		s.fieldMap = fieldMap
	}

	// Get a value map by fieldName
	vmc, ok := fieldMap[fieldName]

	// If the fieldName entry doesn't exist in the container
	if !ok {
		// Create a new map of ints
		valueMap := make(map[string]int)
		// Create a new value map container
		vmc = &valueMapContainer{
			valueMap: valueMap,
		}
		// Add the new value map container to the analytics container
		fieldMap[fieldName] = vmc
	}

	// Get the current count of the given value
	count, ok := vmc.valueMap[value]

	// If there is no current value
	if !ok {
		// Set count to 0
		count = 0
	}

	// Increment the count by one and store it back
	vmc.valueMap[value] = count + 1

}

// CountUniqueValueForField counts the number of unique values
// for a particular field name
func (s *simpleAnalyticsContainer) CountUniqueValuesForField(fieldName string) int {

	// Get a value map by fieldName
	vmc, ok := s.fieldMap[fieldName]

	// If the fieldName entry doesn't exist in the container
	if !ok {
		// Return 0
		return 0
	}

	// Otherwise the number of unique values is the length
	// of the value map
	return len(vmc.valueMap)

}

// GetTopNResults returns the top n results for the given fieldName
func (s *simpleAnalyticsContainer) GetTopNResults(fieldName string, n int) ([]Stat, bool) {

	// Get a value map by fieldName
	vmc, ok := s.fieldMap[fieldName]

	// If the fieldName entry doesn't exist in the container
	if !ok {
		// Return false indicating that there are no results
		return nil, false
	}

	// Convert the value map to a sortable stat list
	statList := valueMapToStatList(vmc.valueMap)

	// Create a new string slice
	results := make([]Stat, 0)

	// Iterate over the stat list
	for i, v := range statList {
		// Iterate while index is less than n
		if i < n {
			// Append fields in this range to results
			results = append(results, v)
		}
	}

	// Return results
	return results, true

}

// valueMapToStatList converts a map[string]int to a sortable
// statList
func valueMapToStatList(valueMap map[string]int) statList {
	// Create a new statList
	statList := make(statList, len(valueMap))
	// Create an index and initialise it to 0
	i := 0
	// Iterate over the valueMap
	for k, v := range valueMap {
		// Add a new stat per key/value pair in the valueMap
		statList[i] = stat{k, v}
		// Increment index
		i++
	}
	// Do a descending sort of the statList
	sort.Sort(sort.Reverse(statList))
	// return the statList
	return statList
}

// stat is a simple key/value pair
type stat struct {
	value string
	count int
}

// GetValue gets the value of a stat
func (s stat) GetValue() string {
	return s.value
}

// GetCount gets the count of a stat
func (s stat) GetCount() int {
	return s.count
}

// statList is a sortable list of stats
type statList []stat

// Len returns the length of a statList
func (p statList) Len() int { return len(p) }

// Less determines which of two stats are greater in value
func (p statList) Less(i, j int) bool { return p[i].count < p[j].count }

// Swap swaps two items in the statList
func (p statList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// NewContainer returns a new analytics container
func NewContainer() Container {
	return &simpleAnalyticsContainer{}
}
