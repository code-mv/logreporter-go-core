package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/code-mv/logreporter-go-core/analytics"

	"github.com/code-mv/logreporter-go-core/filereader"
	"github.com/code-mv/logreporter-go-core/parser"
	"github.com/code-mv/logreporter-go-core/schema"
)

// displayCountStat displays a "count unique values" stat
// on stdout
func displayCountStat(fieldName string, analyticsContainer analytics.Container) {

	// Get the count of unique values from the analytics container
	countStat := analyticsContainer.CountUniqueValuesForField(fieldName)
	// Print the results to stdout
	fmt.Printf("Unique count of field %s = %d\n", fieldName, countStat)

}

// displayTopNStat displays a "what are top N of X" stat to stdout
func displayTopNStat(fieldName string, n int, analyticsContainer analytics.Container) {

	// Get the top n stat from the analytics container
	topNStat, ok := analyticsContainer.GetTopNResults(fieldName, n)

	// If there is no stat for this field name
	if !ok {
		// Print to stdout that there are no stats
		fmt.Printf("There are currently no stats for %s\n", fieldName)
		// Exit
		return
	}

	// Create a new strings builder
	sb := new(strings.Builder)

	// Write the headline
	sb.WriteString(fmt.Sprintf("Top %d stats for %s ...\n", n, fieldName))

	// Iterate over the stats
	for i, v := range topNStat {
		// Write each top stat
		sb.WriteString(fmt.Sprintf("%d. %s (%d)\n", i+1, v.GetValue(), v.GetCount()))
	}

	// Print results to stdout
	fmt.Print(sb.String())

}

// printAllResults prints the intended results to stdout
func printAllResults(analyticsContainer analytics.Container) {

	// Print blank line
	fmt.Println()
	// Display the number of unique IP addresses
	displayCountStat(schema.IPAddress, analyticsContainer)
	// Print blank line
	fmt.Println()
	// Display the top 3 most common URLs
	displayTopNStat(schema.URLPath, 3, analyticsContainer)
	// Print blank line
	fmt.Println()
	// Display the top 3 most common IP addresses
	displayTopNStat(schema.IPAddress, 3, analyticsContainer)
	// Print blank line
	fmt.Println()

}

func main() {

	// Define a filePath flag
	filePath := flag.String("filePath", `C:\programming-task-example-data.log`, "The path to the log file being reported on")

	// Parses all flags
	flag.Parse()

	// Create a channel for streaming entries from the log file
	entriesChannel := make(chan string)

	// Get new file reader
	fileReader := filereader.NewFileReader()

	// Read the lines of the log file
	go fileReader.ReadEntries(*filePath, entriesChannel)

	// Get log schema
	logSchema := schema.NewLogSchema()

	// Get log entry parser
	parser := parser.NewLogEntryParser(logSchema)

	// Create new analytics container
	analyticsContainer := analytics.NewContainer()

	// Iterate over entries channel
	for logEntry := range entriesChannel {
		// Parse the log entry
		parsedEntry := parser.Parse(&logEntry)
		// Add stats to analytics container
		analyticsContainer.AddStats(parsedEntry)
	}

	// Prints the results to stdout
	printAllResults(analyticsContainer)

}
