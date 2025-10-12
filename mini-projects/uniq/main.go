package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
parseLines() doesn't care whether the line was read from a file or stdin.
It just compares the current read line with previous one and returns the result.
*/
func parseLines(line *string, prevLine *string) string {
	switch *prevLine {
	case "":
		*prevLine = *line
		return ""
	case *line:
		return ""
	}
	result := *prevLine
	*prevLine = *line
	return result
}

// Read from stdin, read lines case-sensitive, write a copy of adjacent duplicates to stdout
func readAndParseLines() {
	scanner := bufio.NewScanner(os.Stdin)

	var prevLine string
	for scanner.Scan() {
		line := scanner.Text()
		curr := parseLines(&line, &prevLine)
		if curr != "" {
			fmt.Println(curr)
		}
	}
}

func main() {
	readAndParseLines()
}
