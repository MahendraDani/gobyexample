package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
  /*
  - map stores a set of key/value pairs and provides CONSTANT-TIME operations to store, retrieve or test for an item.
  - The key can be of any types whose value can be compared with ==, the value may be of any type.
  */
  counts := make(map[string]int)

  /*
  - bufio package is used for reading input in efficient manner.
  */
  input := bufio.NewScanner(os.Stdin)

  /*
  - each call to input.Scan() reads the next line and removes the newline character from end.
  - the result can be retrieved by calling input.Text().
  - Scan() function returns true if there is a line to read, else it returns false.
  */
	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
      /*
      - Printf() in Go, is similar to printf()
      */
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
