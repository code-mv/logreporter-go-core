package filereader

import (
	"bufio"
	"os"

	"github.com/code-mv/logreporter-go-core/utils/errors"
)

type simpleFileReader struct {
}

// Read entries reads out log entries one by one and
// sends then into the entries channel
func (s *simpleFileReader) ReadEntries(filePath string, entriesChannel chan string) {

	// Check preconditions
	errors.CheckMandatoryFields(filePath)

	// Open file
	file, err := os.Open(filePath)

	// Handle error while opening file
	errors.ThrowOnErrorf(err, errors.OpenFileError, "Failed to open file with path = %s", filePath)

	// Create new bufio scanner
	scanner := bufio.NewScanner(file)

	// Custom split function
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanLines(data, atEOF)
		return
	}

	// Configure custom split function
	scanner.Split(split)

	// Defer the closing of the entries channel
	defer close(entriesChannel)

	// Scan file based on custom split function
	for scanner.Scan() {
		// Send log entry into the entries channel
		entriesChannel <- scanner.Text()

		// Handle scanner errors
		if err := scanner.Err(); err != nil {
			errors.ThrowOnErrorf(err, errors.FileReadError, "Invalid input = %s", err)
		}
	}

}

// NewFileReader is a service for reading in files
func NewFileReader() FileReader {
	return &simpleFileReader{}
}
