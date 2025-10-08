package main

import (
  "fmt"
  "bufio"
  "os"
)

func main(){
  reader := bufio.NewReader(os.Stdin)
  line, _ := reader.ReadString('\n')
  fmt.Println("Read: ", line)
}
