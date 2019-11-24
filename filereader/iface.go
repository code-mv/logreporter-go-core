package filereader

// FileReader reads a file from the filesystem
// entry by entry
type FileReader interface {
	// Read entries reads out log entries one by one and
	// sends then into the entries channel
	ReadEntries(filePath string, entriesChannel chan string)
}
