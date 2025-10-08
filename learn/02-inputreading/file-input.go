package main

import (
  "fmt"
  "os"
  "bufio"
)

func main(){
  f, _ := os.Open("data.txt")
  scanner := bufio.NewScanner(f)

  for scanner.Scan() {
    fmt.Println(scanner.Text())
  }
  f.Close()
}
