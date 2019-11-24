package schema

const (
	// IPAddress field
	IPAddress string = "ipAddress"
	// UnknownField1 field
	UnknownField1 string = "unknownField1"
	// Username field
	Username string = "username"
	// Timestamp field
	Timestamp string = "timestamp"
	// HTTPMethod field
	HTTPMethod string = "httpMethod"
	// URLPath field
	URLPath string = "urlPath"
	// HTTPVersion field
	HTTPVersion string = "httpVersion"
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

// LogSchema defines a set of field definitions
// comprising a schema of fields in a log entry
type LogSchema interface {
	MapFields(positionalValues []string) map[string]string
}

// FieldMapper collects positional values and returns
// one or more derived fields
type FieldMapper interface {
	Map(value string) map[string]string
}
