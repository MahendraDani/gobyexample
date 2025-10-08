package main

import (
  "fmt"
  "os"
  "bufio"
)

func main(){
  fmt.Print("Input: ")
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    line := scanner.Text()
    fmt.Println("Read: ", line)
    fmt.Print("Input: ")
  }
}
