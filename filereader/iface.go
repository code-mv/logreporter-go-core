package filereader

// FileReader reads a file from the filesystem
// entry by entry
type FileReader interface {
	ReadEntries(filePath string, entriesChannel chan string)
}
