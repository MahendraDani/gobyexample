package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()

	file, _ := os.Open("data.txt")
	defer file.Close()

	buf := make([]byte, 1)
	// var line []byte
	lineCount := 0

	for {
		n, err := file.Read(buf)
		if n == 0 || err != nil {
			// Print the last line if there's any remaining content
			// if len(line) > 0 {
			// 	fmt.Println(string(line))
			// 	lineCount++
			// }
			break
		}

		if buf[0] == '\n' {
			// Found newline, print the line
			// fmt.Println(string(line))
			lineCount++
			// line = nil
		} else {
			// Append byte to current line
			// line = append(line, buf[0])
		}
	}

	fmt.Printf("\nUnbuffered read: %d lines in %v\n", lineCount, time.Since(start))
}
