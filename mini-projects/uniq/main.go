package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
parseLines compares the current line with the previous one.
It returns the previous line if it's different from the current one.
*/
func parseLines(line *string, prevLine *string, isFirst *bool) string {
	if *isFirst {
		// First line - don't print yet
		*prevLine = *line
		*isFirst = false
		return ""
	}

	if *prevLine != *line {
		// Lines are different - print the previous line
		result := *prevLine
		*prevLine = *line
		return result
	}

	// Lines are the same - skip (but update prevLine to current)
	*prevLine = *line
	return ""
}

/*
readAndParseLines reads lines from the given reader,
uses parseLines() to filter duplicates, and prints the result.
*/
func readAndParseLines(r io.Reader) {
	scanner := bufio.NewScanner(r)
	var prevLine string
	isFirst := true

	for scanner.Scan() {
		line := scanner.Text()
		curr := parseLines(&line, &prevLine, &isFirst)

		if curr != "" {
			fmt.Println(curr)
		}
	}

	// Print the last line (it's always been buffered and not printed yet)
	if !isFirst {
		fmt.Println(prevLine)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading:", err)
		os.Exit(1)
	}
}

func main() {

	if len(os.Args) > 1 {
		if os.Args[1] == "-" {
			readAndParseLines(os.Stdin)
		} else {
			file, err := os.Open(os.Args[1])
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error opening file:", err)
				os.Exit(1)
			}
			defer file.Close()

			readAndParseLines(file)
		}
	} else {
		readAndParseLines(os.Stdin)
	}
}
