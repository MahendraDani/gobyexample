package main

import (
  "fmt"
  "os"
)

func main(){
  buf := make([]byte, 100)
  n, _ := os.Stdin.Read(buf)
  fmt.Println("Read bytes: ", n)
  fmt.Println("Data: ", string(buf[:n]))
}
