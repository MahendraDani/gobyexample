package main

import (
  "bufio"
  "fmt"
  "os"
)

// Read from stdin, read lines case-sensitive, write a copy of adjacent duplicates to stdout
func readAndParseLines(){
  scanner := bufio.NewScanner(os.Stdin)
  
  var prevLine string
  for scanner.Scan() {
    line := scanner.Text()
    if(prevLine == ""){
      prevLine = line
    }else if(prevLine==line){
      continue
    }else{
      fmt.Println(prevLine)
      prevLine = line
    }
  }
}

func main(){
  readAndParseLines()
}
