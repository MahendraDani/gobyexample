package main

import (
  "bufio"
  "fmt"
  "os"
)

func main(){
  counts := make(map[string]int)
  files := os.Args[1:]

  if len(files)==0 {
    countLines(os.Stdin, counts)
  } else {
    for _, arg := range files {
      f, err := os.Open(arg)
      if err != nil {
        fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
        continue
      }
      countLines(f, counts)
      f.Close()
    }
  }

  for line, n := range counts {
    if n > 1 {
      fmt.Printf("%d\t%s\n",n,line)
    }
  }
}

/*
- A map is a reference to the data structure created by make
- When a map is passed to a function, the function receives a copy of the reference, so any changes the called function makes to the underlying data structure will be visible through the caller’s map reference too.
*/
func countLines(f *os.File, counts map[string]int){
  input := bufio.NewScanner(f)
  for input.Scan(){
    counts[input.Text()]++
  }
}
