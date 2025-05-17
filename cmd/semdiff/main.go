package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	v3 "github.com/xeger/semdiff/openapi/v3"
)

func main() {
	verbose := flag.Bool("v", false, "verbose output")
	flag.Parse()

	if flag.NArg() != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s [-v] <fileA> <fileB>\n", os.Args[0])
		os.Exit(1)
	}

	fileA := flag.Arg(0)
	fileB := flag.Arg(1)

	// Load and unmarshal first file
	fA, err := os.Open(fileA)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening %s: %v\n", fileA, err)
		os.Exit(1)
	}
	defer fA.Close()

	docA, err := v3.Unmarshal(fA)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading %s: %v\n", fileA, err)
		os.Exit(1)
	}

	// Load and unmarshal second file
	fB, err := os.Open(fileB)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening %s: %v\n", fileB, err)
		os.Exit(1)
	}
	defer fB.Close()

	docB, err := v3.Unmarshal(fB)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading %s: %v\n", fileB, err)
		os.Exit(1)
	}

	// Run diff
	diff := v3.Diff(docA, docB)

	// Decide what to output
	var output interface{} = diff.Change
	if *verbose {
		output = diff
	}

	jsonBytes, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling diff: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(jsonBytes))
}
