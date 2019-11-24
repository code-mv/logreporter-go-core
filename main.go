package main

import (
	"github.com/code-mv/logreporter-go-core/filereader"
	"github.com/code-mv/logreporter-go-core/parser"
	"github.com/code-mv/logreporter-go-core/schema"
)

func main() {

	// Create a channel for streaming entries from the log file
	entriesChannel := make(chan string)

	// Get new file reader
	fileReader := filereader.NewFileReader()

	// Read the lines of the log file
	go fileReader.ReadEntries(`P:\workspaces\digio\tests\programming-task-example-data.log`, entriesChannel)

	// Get log schema
	logSchema := schema.NewLogSchema()

	// Get log entry parser
	parser := parser.NewLogEntryParser(logSchema)

	for logEntry := range entriesChannel {
		parsedEntry := parser.Parse(&logEntry)
	}

}
