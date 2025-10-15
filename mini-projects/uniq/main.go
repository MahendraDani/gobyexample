package main

import (
	"bufio"
	"fmt"
	"os"
  "io"
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

func readAndParseLines(r io.Reader) {
	scanner := bufio.NewScanner(r)

	var prevLine string
	for scanner.Scan() {
		line := scanner.Text()
		curr := parseLines(&line, &prevLine)
		if curr != "" {
			fmt.Println(curr)
		}
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
      readAndParseLines(file)
      defer file.Close()
    }
  }else{
    readAndParseLines(os.Stdin)
  }
}
