package parser

import (
	"testing"

	"github.com/code-mv/logreporter-go-core/schema"
	"github.com/code-mv/logreporter-go-core/utils/errors"
)

func TestParseExampleLogFile(t *testing.T) {

	// Defer catch function
	defer errors.Catch(func(r interface{}) {
		t.Errorf("Caught panic with message = %s", r)
	}, errors.RegexCaptureNotFoundError)

	// Input string
	input := `72.44.32.10 - - [09/Jul/2018:15:48:07 +0200] "GET / HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (compatible; MSIE 10.6; Windows NT 6.1; Trident/5.0; InfoPath.2; SLCC1; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729; .NET CLR 2.0.50727) 3gpp-gba UNTRUSTED/1.0" junk extra`

	// Create new parser
	parser := NewLogEntryParser(schema.NewLogSchema())

	// Parse input
	result := parser.Parse(&input)

	// Define expected ip address
	expectedIPAddress := "72.44.32.10"

	// Fail if ip address not correct
	if result[schema.IPAddress] != expectedIPAddress {
		t.Errorf("Expected ip address = %s but was %s", expectedIPAddress, result[schema.IPAddress])
	}

	// Define expected unknown field 1
	expectedUnknownField1 := "-"

	// Fail if unknownField1 not correct
	if result[schema.UnknownField1] != expectedUnknownField1 {
		t.Errorf("Expected unknownField1 = %s but was %s", expectedUnknownField1, result[schema.UnknownField1])
	}

	// Define expected username
	expectedUsername := "-"

	// Fail if username not correct
	if result[schema.Username] != expectedUsername {
		t.Errorf("Expected username = %s but was %s", expectedUsername, result[schema.Username])
	}

	// Define expected timestamp
	expectedTimestamp := "09/Jul/2018:15:48:07 +0200"

	// Fail if timestamp not correct
	if result[schema.Timestamp] != expectedTimestamp {
		t.Errorf("Expected timestamp = %s but was %s", expectedTimestamp, result[schema.Timestamp])
	}

	// Define expected httpMethod
	expectedHTTPMethod := "GET"

	// Fail if httpDetails not correct
	if result[schema.HTTPMethod] != expectedHTTPMethod {
		t.Errorf("Expected httpDetails = %s but was %s", expectedHTTPMethod, result[schema.HTTPMethod])
	}

	// Define expected url path
	expectedURLPath := "/"

	// Fail if httpDetails not correct
	if result[schema.URLPath] != expectedURLPath {
		t.Errorf("Expected httpDetails = %s but was %s", expectedURLPath, result[schema.URLPath])
	}

	// Define expected http version
	expectedHTTPVersion := "HTTP/1.1"

	// Fail if httpDetails not correct
	if result[schema.HTTPVersion] != expectedHTTPVersion {
		t.Errorf("Expected httpDetails = %s but was %s", expectedHTTPVersion, result[schema.HTTPVersion])
	}

	// Define expected unknownField2
	expectedUnknownField2 := "-"

	// Fail if httpDetails not correct
	if result[schema.UnknownField2] != expectedUnknownField2 {
		t.Errorf("Expected unknownField2 = %s but was %s", expectedUnknownField2, result[schema.UnknownField2])
	}

	// Define expected userAgentDetails
	expectedUserAgentDetails := "Mozilla/5.0 (compatible; MSIE 10.6; Windows NT 6.1; Trident/5.0; InfoPath.2; SLCC1; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729; .NET CLR 2.0.50727) 3gpp-gba UNTRUSTED/1.0"

	// Fail if userAgentDetails not correct
	if result[schema.UserAgentDetails] != expectedUserAgentDetails {
		t.Errorf("Expected userAgentDetails = %s but was %s", expectedUserAgentDetails, result[schema.UserAgentDetails])
	}

}
