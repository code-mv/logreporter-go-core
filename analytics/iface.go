package analytics

// Container stores values and performs
// aggregate functions against those values
type Container interface {
	// AddStats adds a map of fields and values, parsed from
	// a log file, to the stat container
	AddStats(fields map[string]string)
	// AddStat adds a single field and value to the analytics container
	AddStat(fieldName string, value string)
	// CountUniqueValueForField counts the number of unique values
	// for a particular field name
	CountUniqueValuesForField(fieldName string) int
	// GetTopNResults returns the top n results for the given fieldName
	GetTopNResults(fieldName string, n int) (SameCountGroupList, bool)
}

// Stat is a value and a corresponding count of number
// of data points with that value
type Stat interface {
	// GetValue gets the value
	GetValue() string
	// GetCount gets the count corresponding to the field
	GetCount() int
}

// SameCountGroup is a grouping of stats with the same score
type SameCountGroup interface {
	GetStats() []Stat
	AddStat(stat Stat)
}

// SameCountGroupList is a list of same count group
type SameCountGroupList []SameCountGroup
